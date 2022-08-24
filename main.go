package main

import (
	"be11/account-service-app/config"
	"be11/account-service-app/controllers/transfers"
	"be11/account-service-app/controllers/users"
	"be11/account-service-app/entities"
	"fmt"
	"os"
)

func main() {
	db := config.ConnectToDatabase()
	defer db.Close()

	var back = "n"

	for back != "y" {
		if back == "close" {
			fmt.Println("Terimakasih telah bertransaks")
			os.Exit(1)
		}

		var option = 0

		// menu sign in dan sign up
		fmt.Print("Account Service\n 1. Sign In\n 2. Sign Up\n")
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
			for backMenu != "n" {
				dataUser, err, strErr := users.SignIn(db, inputTelp, inputPass)
				if err != nil {
					fmt.Print(strErr, err)
				} else {
					fmt.Print("\n---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------\n")
					for _, v := range dataUser {
						if v.Gender == "men" {
							v.Gender += "\t"
						}
						fmt.Println("| Username :", v.Name, "\t", "| Gender :", v.Gender, "\t", "| no telp :", v.NoTelp, "\t",
							"| Currency :", v.Currency, "\t", "| Balance :", v.Balance, "\t", "| Created_at :", v.CreatedAt, "\t", "| Updated_at :", v.UpdateAt, "|")
					}
					fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

					// Menu setelah log in
					var optionMenuLog int

					fmt.Print("Account Service\n 1. Search Other User\n 2. Update Data\n 3. Delete Data\n 4. Top Up \n 5. Transfer\n 6. Top Up History\n 7. Transfer History\n 8. Exit Porgram")
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
							fmt.Print(strErrRead, errRead)
						} else {

							fmt.Print("\n------------------------------------------------------------------------------------------------------------------------------\n")
							for _, v := range resRead {
								if v.Gender == "men" {
									v.Gender += "\t"
								}
								fmt.Println("| Username :", v.Name, "\t", "| Gender :", v.Gender, "\t", "| Created_at :", v.Created_at, "\t", "| Updated_at :", v.Updated_at, "|")
							}
							fmt.Println("--------------------------------------------------------------------------------------------------------------------------------")

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

					// fitur top up
					case 4:

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

					// fitur top up history
					case 6:

					// fitur transfer history
					case 7:

					// fitur exit
					case 8:
						fmt.Println("Terimakasih telah bertransaks")
						os.Exit(1)
					}

				}

			}

			// fitur sign up
		} else if option == 2 {

			var newUser entities.Users
			fmt.Print("input no telp : ")
			fmt.Scan(&newUser.NoTelp)

			fmt.Print("input username : ")
			fmt.Scan(&newUser.Name)

			fmt.Print("input gender : ")
			fmt.Scan(&newUser.Gender)

			fmt.Print("input password : ")
			fmt.Scan(&newUser.Password)

			resInsert, err := users.SignUp(db, newUser)

			if err != nil {
				fmt.Print("sign up error :", err.Error())
			} else if resInsert > 0 {
				fmt.Print("Sign Up success\n")

			}

		}
	}

}
