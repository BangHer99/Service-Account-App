package main

import (
	"be11/account-service-app/config"
	"be11/account-service-app/controllers/topups"
	"be11/account-service-app/controllers/transfers"
	"be11/account-service-app/controllers/users"
	"be11/account-service-app/entities"
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := config.ConnectToDatabase()
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	var back = "n"

	for back != "y" {

		var option = 0

		// menu sign in dan sign up
		fmt.Print("Account Service\n 1. Sign In\n 2. Sign Up\n 3. Exit\n")
		fmt.Print("Enter your choice : ")
		fmt.Scan(&option)

		// fitur sign in
		if option == 1 {
			var inputTelp int
			var inputPass string

			fmt.Print("Input no telp : ")
			fmt.Scan(&inputTelp)
			fmt.Print("Input password: ")
			fmt.Scan(&inputPass)

			var backMenu = ""
			var tempName = ""
			for backMenu != "n" {
				dataUser, err, strErr := users.SignIn(db, inputTelp, inputPass)
				if err != nil {
					fmt.Print(strErr, err)
					backMenu = "n"
				} else if len(dataUser) == 0 {
					fmt.Println("Account not found")
					backMenu = "n"
				} else {
					fmt.Print("\n----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------\n")
					for _, v := range dataUser {
						if v.Gender == "men" {
							v.Gender += "\t"
						}
						fmt.Println("| Username :", v.Name, "\t", "| Gender :", v.Gender, "\t", "| no telp :", v.NoTelp, "\t",
							"| Currency :", v.Currency, "\t", "| Balance :", v.Balance, "\t", "| Created_at :", v.CreatedAt, "\t", "| Updated_at :", v.UpdateAt, "|")

						tempName = v.Name
					}
					fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

					// Menu setelah log in
					var optionMenuLog int

					fmt.Print("Menu Account Service\n 1. Search Other User\n 2. Update Data\n 3. Delete Account\n 4. Top Up \n 5. Transfer\n 6. Top Up History\n 7. Transfer History\n 8. Exit Porgram\n")
					fmt.Print("Enter Your Choice : ")
					fmt.Scan(&optionMenuLog)

					switch optionMenuLog {

					// fitur read other user
					case 1:

						var inputRead string
						fmt.Print("input phone number other user : ")
						fmt.Scan(&inputRead)

						resRead, errRead, strErrRead := users.ReadOtherUser(db, inputRead)
						if errRead != nil {
							fmt.Println(strErrRead, errRead)
						} else if len(resRead) == 0 {
							fmt.Println("phone number not found")
						} else {

							fmt.Print("\n----------------------------------------------------------------------------------------------------------------------------\n")
							for _, v := range resRead {
								if v.Gender == "men" {
									v.Gender += "\t"
								}
								fmt.Println("| Username :", v.Name, "\t", "| Gender :", v.Gender, "\t", "| Created_at :", v.Created_at, "\t", "| Updated_at :", v.Updated_at, "|")
							}
							fmt.Println("----------------------------------------------------------------------------------------------------------------------------")

						}

					// fitur update
					case 2:

						var updateUser entities.Users

						fmt.Print("Update username : ")
						updateUser.Name, _ = reader.ReadString('\n')
						updateUser.Name = strings.TrimSuffix(updateUser.Name, "\n")

						fmt.Print("Update gender : ")
						updateUser.Gender, _ = reader.ReadString('\n')
						updateUser.Gender = strings.TrimSuffix(updateUser.Gender, "\n")

						fmt.Print("Update password : ")
						updateUser.Password, _ = reader.ReadString('\n')
						updateUser.Password = strings.TrimSuffix(updateUser.Password, "\n")

						if updateUser.Password != "" {
							inputPass = updateUser.Password
						}

						updatePassByte := []byte(updateUser.Password)
						hashPass, _ := bcrypt.GenerateFromPassword(updatePassByte, bcrypt.DefaultCost)
						updateUser.Password = string(hashPass)

						succes, strUpt, errUptusers := users.UpdateDataUser(db, &updateUser, inputTelp)
						if err != nil {
							fmt.Print(strUpt, errUptusers)
						} else if succes >= 0 {
							fmt.Println("\n update success ")
						}

					// fitur delete
					case 3:

						var selectOption string

						for selectOption != "y" && selectOption != "n" {

							fmt.Print("apakah anda yakin ingin menghapus akun? (y/n) : ")
							fmt.Scan(&selectOption)

							if selectOption == "y" {
								ress, strErrr, errDel := users.DeleteDataUser(db, inputTelp)
								if errDel != nil {
									fmt.Print(strErrr, errDel)
								} else if ress >= 0 {
									fmt.Println("Succes Delete Account")
								}

								backMenu = "n"

							} else if selectOption == "n" {

								backMenu = "y"

							} else {

								fmt.Println("input wrong, Please input y or n")

							}

						}

					// fitur top up
					case 4:

						var topup entities.TopUps
						var toTelp int
						fmt.Print("input to account for top-up (no_telp) : ")
						fmt.Scan(&toTelp)

						if toTelp != -1 {
							fmt.Print("input amout : ")
							fmt.Scan(&topup.Amount)

						}

						_, errTop, strTop := topups.TopUp(db, inputTelp, toTelp, topup)
						if errTop != nil {
							fmt.Println(errTop, strTop)
						} else {
							fmt.Println(strTop)
							for _, value := range dataUser {
								value.Balance += topup.Amount
							}
						}

					// fitur transfer
					case 5:

						var transfer entities.Transfers

						fmt.Println("From account transfer(no_telp) : ", inputTelp)
						transfer.From_account_telp = inputTelp

						fmt.Print("input to account transfer(no_telp) : ")
						fmt.Scan(&transfer.To_account_telp)

						if transfer.To_account_telp != 0 {
							fmt.Print("input amount : ")
							fmt.Scan(&transfer.Amount)

						}

						toId := transfer.To_account_telp

						_, errTf, strTf := transfers.Transfer(db, inputTelp, toId, transfer)
						if errTf != nil {
							fmt.Println(errTf, strTf)
						} else {
							fmt.Println(strTf)
							for _, v := range dataUser {
								v.Balance -= transfer.Amount
							}
						}

					// fitur top up history
					case 6:

						var historytopup entities.TopUpHistory

						ressHT, errHt := topups.ToupUpHistory(db, inputTelp, historytopup)
						if errHt != nil {
							fmt.Println(errHt.Error())
						} else if len(ressHT) == 0 {
							fmt.Println("phone number not found")
						} else {

							fmt.Print("\n-----------------------------------------------------------------------------------------------------------------------------------------------------------\n")
							for _, v := range ressHT {
								if v.NameUser == tempName {
									fmt.Println("| date topup :", v.CreatedAt, "\t", "| from : ", v.NameUser, "\t", "| to :", v.To_account_name, "\t", "| Amount : +", v.Amount, "\t", "| Receive Money By Topup  |")
								} else {
									fmt.Println("| date topup :", v.CreatedAt, "\t", "| to :", v.To_account_name, "\t", "| from : ", v.NameUser, "\t", "| Amount : +", v.Amount, "\t", "| Receive Money By Topup  |")
								}
							}
							fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------")
						}

					// fitur transfer history
					case 7:

						var transferHistory entities.TransferHistory
						ressTh, errTh := transfers.TransferHistory(db, inputTelp, transferHistory)
						if errTh != nil {
							fmt.Println(errTh.Error())
						} else if len(ressTh) == 0 {
							fmt.Println("phone number not found")
						} else {

							fmt.Print("\n------------------------------------------------------------------------------------------------------------------------------------------------------\n")
							for _, v := range ressTh {
								if v.From_account_name == tempName {
									fmt.Println("| date transfer :", v.Created_at, "\t", "| id :", v.Id, "\t", "| from :", v.From_account_name, "\t", "| to :", v.To_account_name, "\t", "| Amount : -", v.Amount, "\t", "| Spent Money By transasfer  |")
								} else if v.To_account_name == tempName {
									fmt.Println("| date transfer :", v.Created_at, "\t", "| id :", v.Id, "\t", "| to :", v.To_account_name, "\t", "| from :", v.From_account_name, "\t", "| Amount : +", v.Amount, "\t", "| Receive money By transfers |")
								}
							}
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------")

						}
					// fitur exit
					case 8:

						back = "y"

					default:

						fmt.Println("input wrong, please input 1, 2, 3, 4, 5, 6 ,7 or 8 ")

					}

					if optionMenuLog == 1 || optionMenuLog == 2 || optionMenuLog == 4 || optionMenuLog == 5 || optionMenuLog == 6 || optionMenuLog == 7 {
						if back != "y" && backMenu != "n" {
							backMenu = "kosong"

							for backMenu == "kosong" {
								fmt.Print("Back to Menu (y/n) : ")
								fmt.Scan(&backMenu)

								if backMenu != "y" && backMenu != "n" {
									fmt.Println("input wrong, please input y or n")
									backMenu = "kosong"
								} else if backMenu == "n" {
									back = "y"
								}
							}
						}

					}
				}
			}

			// fitur sign up
		} else if option == 2 {

			var succes string

			for succes != "y" {
				var newUser entities.Users
				fmt.Print("input no telp : ")
				fmt.Scan(&newUser.NoTelp)

				fmt.Print("input username : ")
				newUser.Name, _ = reader.ReadString('\n')
				newUser.Name = strings.TrimSuffix(newUser.Name, "\n")

				fmt.Print("input gender : ")
				newUser.Gender, _ = reader.ReadString('\n')
				newUser.Gender = strings.TrimSuffix(newUser.Gender, "\n")

				fmt.Print("input password : ")
				newUser.Password, _ = reader.ReadString('\n')
				newUser.Password = strings.TrimSuffix(newUser.Password, "\n")
				inputPassByte := []byte(newUser.Password)
				hardPass, _ := bcrypt.GenerateFromPassword(inputPassByte, bcrypt.DefaultCost)
				newUser.Password = string(hardPass)

				if newUser.NoTelp != -1 && newUser.Name != "" && newUser.Gender != "" && newUser.Password != "" {
					resInsert, err := users.SignUp(db, &newUser)
					if err != nil {
						fmt.Print("sign up error :", err.Error(), "\n")
					} else if resInsert > 0 {
						fmt.Print("Sign Up success\n")
						succes = "y"
					}
				} else {
					fmt.Println("Mohon lengkapi data terlebih dahulu")
				}
			}
		} else if option == 3 {
			back = "y"

		} else {
			fmt.Println("input wrong, please input 1,2 or 3")
		}
	}
	if back == "y" {
		fmt.Println("Terima Kasih Telah Bertransaksi")
	}

}
