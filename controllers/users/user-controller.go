package users

import (
	"be11/account-service-app/entities"
	"database/sql"
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
