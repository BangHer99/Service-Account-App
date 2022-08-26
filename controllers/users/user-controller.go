package users

import (
	"be11/account-service-app/entities"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(db *sql.DB, register *entities.Users) (int, error) {

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
		return nil, err, "phone number not registered"
	}

	dataUser := []entities.Users{}
	for statementLogin.Next() {
		var rowUser entities.Users
		err := statementLogin.Scan(&rowUser.NoTelp, &rowUser.Password, &rowUser.Name, &rowUser.Gender, &rowUser.Balance, &rowUser.Currency, &rowUser.CreatedAt, &rowUser.UpdateAt)
		if err != nil {
			return nil, err, ""
		}

		dataUser = append(dataUser, rowUser)

		errMatch := bcrypt.CompareHashAndPassword([]byte(rowUser.Password), []byte(inputPass))
		if errMatch == nil {
			return dataUser, nil, "succes"
		} else {
			return nil, err, "password wrong "
		}

	}

	return nil, nil, ""

}

func UpdateDataUser(db *sql.DB, updateUser *entities.Users, inputTelp int) (int, string, error) {

	if updateUser.Name != "" && updateUser.Password != "" && updateUser.Gender != "" {
		statement, err := db.Prepare("UPDATE users SET name_user = ?, gender = ?, password_user = ? WHERE no_telp = ?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resUpdate, err := statement.Exec(updateUser.Name, updateUser.Gender, updateUser.Password, inputTelp)
		if err != nil {
			return -1, "error update : ", err
		} else {
			row, _ := resUpdate.RowsAffected()
			return int(row), "succes update", nil
		}

	} else if updateUser.Name != "" && updateUser.Password == "" && updateUser.Gender != "" {
		statement, err := db.Prepare("UPDATE users SET name_user = ?, gender = ? WHERE no_telp = ?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resUpdate, err := statement.Exec(updateUser.Name, updateUser.Gender, inputTelp)
		if err != nil {
			return -1, "error update : ", err
		} else {
			row, _ := resUpdate.RowsAffected()
			return int(row), "succes update", nil
		}
	} else if updateUser.Name != "" && updateUser.Password != "" && updateUser.Gender == "" {
		statement, err := db.Prepare("UPDATE users SET name_user = ? password_user = ? WHERE no_telp = ?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resUpdate, err := statement.Exec(updateUser.Name, updateUser.Password, inputTelp)
		if err != nil {
			return -1, "error update : ", err
		} else {
			row, _ := resUpdate.RowsAffected()
			return int(row), "succes update", nil
		}
	} else if updateUser.Name != "" && updateUser.Password == "" && updateUser.Gender == "" {
		statement, err := db.Prepare("UPDATE users SET name_user = ? WHERE no_telp = ?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resUpdate, err := statement.Exec(updateUser.Name, inputTelp)
		if err != nil {
			return -1, "error update : ", err
		} else {
			row, _ := resUpdate.RowsAffected()
			return int(row), "succes update", nil
		}
	} else if updateUser.Name == "" && updateUser.Password != "" && updateUser.Gender != "" {
		statement, err := db.Prepare("UPDATE users SET gender = ?, password_user = ? WHERE no_telp = ?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resUpdate, err := statement.Exec(updateUser.Gender, updateUser.Password, inputTelp)
		if err != nil {
			return -1, "error update : ", err
		} else {
			row, _ := resUpdate.RowsAffected()
			return int(row), "succes update", nil
		}
	} else if updateUser.Name == "" && updateUser.Password == "" && updateUser.Gender != "" {
		statement, err := db.Prepare("UPDATE users SET gender = ? WHERE no_telp = ?")
		if err != nil {
			return -1, "error statement : ", err
		}
		resUpdate, err := statement.Exec(updateUser.Gender, inputTelp)
		if err != nil {
			return -1, "error update : ", err
		} else {
			row, _ := resUpdate.RowsAffected()
			return int(row), "succes update", nil
		}
	} else {
		return 0, "", nil
	}

}

func DeleteDataUser(db *sql.DB, inputTelp int) (int, string, error) {

	statement, err := db.Prepare("DELETE FROM users WHERE no_telp=?")
	if err != nil {
		return -1, "error statement : ", err
	}
	resDelete, err := statement.Exec(inputTelp)
	if err != nil {
		return -1, "Error Delete : ", err
	} else {
		row, errDel := resDelete.RowsAffected()
		if errDel != nil {
			return -1, "Error Delete", errDel
		}
		return int(row), "succes delete", err
	}

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
