package main

import (
	"be11/account-service-app/config"
	"be11/account-service-app/controllers/users"
	"be11/account-service-app/entities"
	"fmt"
)

func main() {
	db := config.ConnectToDatabase()
	defer db.Close()

	var option int
	var back = "n"

	for back != "y" {

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

			dataUser, err, strErr := users.SignIn(db, inputTelp, inputPass)
			if err != nil {
				fmt.Print(strErr, err)
			} else {

				fmt.Print("\n--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------\n")
				for _, v := range dataUser {
					if v.Gender == "men" {
						v.Gender += "\t"
					}
					fmt.Println("| Username :", v.Name, "\t", "| Gender :", v.Gender, "\t", "| no telp :", v.NoTelp, "\t",
						"| Currency :", v.Currency, "\t", "| Balance :", v.Balance, "\t", "| Created_at :", v.CreatedAt, "\t", "| Updated_at :", v.UpdateAt, "|")
				}
				fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

				// Menu setelah log in
				var optionMenuLog int
				var backMenu = "y"

				for backMenu != "n" {
					fmt.Print("Account Service\n 1. Update Data\n 2. Delete Data\n 3. Top Up \n 4. Transfer\n 5. Top Up History\n 6. Transfer History\n")
					fmt.Scan(&optionMenuLog)

					switch optionMenuLog {

					// fitur update
					case 1:
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
					case 2:
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
					case 3:

					// fitur transfer
					case 4:

					// fitur top up history
					case 5:

					// fitur transfer history
					case 6:
					}

				}

			}

			// fitur sign up
		} else if option == 2 {

			var newUser entities.Users
			fmt.Print("input no telp : ")
			fmt.Scan(&newUser.NoTelp)
			if newUser.NoTelp != 0 {
				fmt.Print("input username : ")
				fmt.Scan(&newUser.Name)
				if newUser.Name != "" {
					fmt.Print("input gender : ")
					fmt.Scan(&newUser.Gender)
					if newUser.Gender != "" {
						fmt.Print("input Password : ")
						fmt.Scan(&newUser.Password)
					}
				}

			}

			resInsert, err := users.SignUp(db, newUser)

			if err != nil {
				fmt.Print("sign up error :", err.Error())
			} else if resInsert > 0 {
				fmt.Print("Sign Up success\n")

			}

		}
	}

}
