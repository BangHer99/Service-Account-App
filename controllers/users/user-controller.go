package users

import (
	"be11/account-service-app/entities"
	"database/sql"
	"fmt"
)

func SignUp(db *sql.DB, register entities.Users) (int, error) {
	statementRegis, err := db.Prepare("INSERT INTO users(no_telp, password_user, name_user, gender, balance, currency) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return -1, err
	}

	resInsert, err := statementRegis.Exec(register.NoTelp, register.Password, register.Name, register.Gender, register.Balance, "Rupiah")
	if err != nil {
		return -1, err
	} else {
		row, err := resInsert.RowsAffected()
		if err != nil {
			return -1, err
		}
		return int(row), nil
	}

}

func SignIn(db *sql.DB, inputTelp int, inputPass string) ([]entities.Users, error, string) {
	statementLogin, err := db.Query("SELECT * FROM users WHERE no_telp=?", inputTelp)
	if err != nil {
		return nil, err, "phone number not registeredr"
	}

	dataUser := []entities.Users{}
	for statementLogin.Next() {
		var rowUser entities.Users
		err := statementLogin.Scan(&rowUser.NoTelp, &rowUser.Password, &rowUser.Name, &rowUser.Gender, &rowUser.Balance, &rowUser.Currency, &rowUser.CreatedAt, &rowUser.UpdateAt)
		if err != nil {
			return nil, err, ""
		}

		dataUser = append(dataUser, rowUser)

		if rowUser.Password == inputPass {
			return dataUser, nil, "succes"
		} else {
			return nil, err, "password wrong "
		}

	}

	return nil, nil, ""

}

func UpdateDataUser(db *sql.DB, updateUser entities.Users, InputName, InputGender, InputPassword string) (int, string, error) {
	if InputName == "y" {
		fmt.Print("Update username : ")
		fmt.Scan(&updateUser.Name)

		statement, err := db.Prepare("UPDATE users SET name_user = ? WHERE no_telp = ?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resUpdate, err := statement.Exec(updateUser.Name, updateUser.NoTelp)
		if err != nil {
			return -1, "error update : ", err
		} else {
			row, _ := resUpdate.RowsAffected()
			return int(row), "succes update", nil
		}
	}

	if InputGender == "y" {
		fmt.Print("Update Gender : ")
		fmt.Scan(&updateUser.Gender)

		statement, err := db.Prepare("UPDATE users SET gender = ? WHERE no_telp=?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resUpdate, err := statement.Exec(updateUser.Gender, updateUser.NoTelp)
		if err != nil {
			return -1, "error update : ", err
		} else {
			row, _ := resUpdate.RowsAffected()
			return int(row), "Succes Update", nil
		}
	}

	if InputPassword == "y" {
		fmt.Print("Update Password : ")
		fmt.Scan(&updateUser.Password)

		statement, err := db.Prepare("UPDATE users SET password_user = ? WHERE no_telp=?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resUpdate, err := statement.Exec(&updateUser.Password, updateUser.NoTelp)
		if err != nil {
			return -1, "error update : ", err
		} else {
			row, _ := resUpdate.RowsAffected()
			return int(row), "Success Update", nil
		}
	}
	return 0, "success update", nil

}

func DeleteDataUser(db *sql.DB, deleteUser entities.Users, deleteByName string, deleteByNoTelp int) (int, string, error) {
	if deleteByName == "y" {
		fmt.Print("delete name")
		fmt.Scan(&deleteUser.Name)

		statement, err := db.Prepare("DELETE FROM users WHERE no_telp=?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resDelete, err := statement.Exec(deleteUser.Name)
		if err != nil {
			return -1, "Error Delete : ", err
		} else {
			row, _ := resDelete.RowsAffected()
			if row > 0 {
				return -1, "error row afected : ", err
			}
			return int(row), "succes delete", err
		}
	}
	if deleteByNoTelp == 1 {
		fmt.Print("delete no. telp : ")
		fmt.Scan(&deleteUser.NoTelp)

		statement, err := db.Prepare("DELETE FROM  users WHERE no_telp=?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resDelete, err := statement.Exec(deleteUser.Name)
		if err != nil {
			return -1, "error delete : ", err
		} else {
			row, _ := resDelete.RowsAffected()
			if row > 0 {
				return -1, "error row Afected : ", err
			}
		}
	}
	return 0, "Success", nil

}

func ReadOtherUser(db *sql.DB, inputReadTelp string) ([]entities.OtherUser, error, string) {
	statementReadUserByNoTelp, err := db.Query("SELECT name_user, gender, created_at, updated_at FROM users WHERE no_telp=?", inputReadTelp)
	if err != nil {
		return nil, err, "phone number not found"
	}

	otherUser := []entities.OtherUser{}
	for statementReadUserByNoTelp.Next() {
		var rowUser entities.OtherUser
		err := statementReadUserByNoTelp.Scan(&rowUser.Name, &rowUser.Gender, &rowUser.Created_at, &rowUser.Updated_at)
		if err != nil {
			return nil, err, ""
		}

		otherUser = append(otherUser, rowUser)

	}

	return otherUser, nil, ""

}
