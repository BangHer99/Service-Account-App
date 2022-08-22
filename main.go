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

					// fitur delete
					case 2:

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
