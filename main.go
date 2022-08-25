package main

import (
	"be11/account-service-app/config"
	"be11/account-service-app/controllers/topups"
	"be11/account-service-app/controllers/transfers"
	"be11/account-service-app/controllers/users"
	"be11/account-service-app/entities"
	"fmt"
)

func main() {
	db := config.ConnectToDatabase()
	defer db.Close()

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

					fmt.Print("Account Service\n 1. Search Other User\n 2. Update Data\n 3. Delete Data\n 4. Top Up \n 5. Transfer\n 6. Top Up History\n 7. Transfer History\n 8. Exit Porgram\n")
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
						var (
							updateUser    entities.Users
							InputName     string
							InputGender   string
							InputPassword string
						)

						fmt.Print("input no. telp want to change : ")
						fmt.Scan(&updateUser.NoTelp)

						fmt.Print("update username (y/n) : ")
						fmt.Scan(&InputName)
						resName, strName, errName := users.UpdateDataUser(db, updateUser, InputName, "", "")
						if errName != nil {
							fmt.Println(strName, errName)
						} else if resName > 0 {
							fmt.Println("row affected : ", resName)
						}

						fmt.Print("Update gender (y/n) : ")
						fmt.Scan(&InputGender)
						resGen, strGen, errGen := users.UpdateDataUser(db, updateUser, " ", InputGender, " ")
						if errGen != nil {
							fmt.Println(strGen, errGen)
						} else if resGen > 0 {
							fmt.Println("row Afected : ", resGen)
						}

						fmt.Print("update password (y/n)")
						fmt.Scan(&InputPassword)
						resPass, strPass, errPass := users.UpdateDataUser(db, updateUser, " ", "", InputPassword)
						if errPass != nil {
							fmt.Println(strPass, errPass)
						} else if resPass > 0 {
							fmt.Println("row affected : ", resPass)
						}

					// fitur delete
					case 3:

						var (
							deleteUser   entities.Users
							deleteName   string
							deleteNoTelp int
						)

						fmt.Print("delete by name (y/n) : ")
						fmt.Scan(&deleteName)
						resName, strName, errName := users.DeleteDataUser(db, deleteUser, deleteName, deleteNoTelp) //<<<<malah int
						if errName != nil {
							fmt.Println(strName, errName)
						} else if resName > 0 {
							fmt.Println("row affected : ", resName)
						}

						fmt.Print("delete By No Telepon : (y/n)")
						fmt.Scan(&deleteNoTelp)
						resNoTelp, intNotelp, errTelp := users.DeleteDataUser(db, deleteUser, deleteName, deleteNoTelp)
						if errTelp != nil {
							fmt.Println(intNotelp, resNoTelp)
						} else if resNoTelp > 0 {
							fmt.Println("row affected", resNoTelp)
						}

					// fitur top up
					case 4:
						var topup entities.TopUps
						var toTelp int
						fmt.Println("input to account for top-up (no_telp) : ")
						fmt.Scan(&toTelp)

						fmt.Print("input amout : ")
						fmt.Scan(&topup.Amount)

						_, errTop, strTop := topups.TopUp(db, inputTelp, toTelp, topup)
						if errTop != nil {
							fmt.Println(errTop, strTop)
						} else {
							fmt.Println(strTop)
							for _, value := range dataUser {
								value.Balance += topup.Amount
							}
						}

						backMenu = "kosong"

						for backMenu == "kosong" {
							fmt.Print("Back to Menu (y/n) : ")
							fmt.Scan(&backMenu)

							if backMenu != "y" && backMenu != "n" {
								fmt.Println("input wrong, please input y or n")
								backMenu = "kosong"
							} else if backMenu == "n" {
								back = "close"
							}
						}

					// fitur transfer
					case 5:
						var transfer entities.Transfers

						fmt.Println("From account transfer(no_telp) : ", inputTelp)
						transfer.From_account_telp = inputTelp

						fmt.Print("input to account transfer(no_telp) : ")
						fmt.Scan(&transfer.To_account_telp)

						fmt.Print("input amount : ")
						fmt.Scan(&transfer.Amount)

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

							fmt.Print("\n------------------------------------------------------------------------------------------------------------------------------------------------------\n")
							for _, v := range ressHT {
								fmt.Println("| date topup :", v.CreatedAt, "\t", "| to :", v.To_account_name, "\t", "| Amount : +", v.Amount, "\t", "|from : ", v.NameUser, "\t", "| Receive Money By Topup  |")
							}
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------")
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
						backMenu = "n"

					}

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

			// fitur sign up
		} else if option == 2 {

			var succes string

			for succes != "y" {
				var newUser entities.Users
				fmt.Print("input no telp : ")
				fmt.Scan(&newUser.NoTelp)

				fmt.Print("input username : ")
				fmt.Scan(&newUser.Name)

				fmt.Print("input gender : ")
				fmt.Scan(&newUser.Gender)

				fmt.Print("input password : ")
				fmt.Scan(&newUser.Password)

				if newUser.NoTelp != 0 && newUser.Name != "" && newUser.Gender != "" && newUser.Password != "" {
					resInsert, err := users.SignUp(db, newUser)
					if err != nil {
						fmt.Print("sign up error :", err.Error())
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
