package transfers

import (
	"be11/account-service-app/entities"
	"database/sql"
)

func Transfer(db *sql.DB, fromId, toId int, transfer entities.Transfers) (int, error, string) {
	var ress string
	statementInsertTransfer, errIns := db.Prepare("INSERT INTO transfers(from_account_telp, to_account_telp, amount) VALUES (?,?,?)")
	if errIns != nil {
		return -1, errIns, "error statement"
	}

	resInsert, err := statementInsertTransfer.Exec(transfer.From_account_telp, transfer.To_account_telp, transfer.Amount)
	if err != nil {
		return -1, err, "error exec"
	} else {
		_, err := resInsert.RowsAffected()
		if err != nil {
			return -1, err, "err "
		}
	}

	statementFromTransfer, errFrom := db.Query("SELECT balance FROM users WHERE no_telp=?", fromId)
	if errFrom != nil {
		return -1, errFrom, "error minus"
	}

	for statementFromTransfer.Next() {
		var rowUser entities.Users
		errUpdate := statementFromTransfer.Scan(&rowUser.Balance)
		if errUpdate != nil {
			return -1, err, "error scan"
		}

		if rowUser.Balance > transfer.Amount {
			updateBalance, errIns := db.Prepare("UPDATE users SET balance=? WHERE no_telp=?")
			if errIns != nil {
				return -1, errIns, "error exec"
			}
			succesUpdate, errUpd := updateBalance.Exec((rowUser.Balance - transfer.Amount), fromId)
			if errUpd != nil {
				return -1, errIns, "error update"
			} else {
				row, _ := succesUpdate.RowsAffected()
				if row > 0 {
					ress = "transfer success 50%"
				}
			}

		} else {
			return 0, nil, "Saldo Tidak Mencukupi"
		}
	}

	statementToTransfer, errFrom := db.Query("SELECT balance FROM users WHERE no_telp=?", toId)
	if errFrom != nil {
		return -1, errFrom, "error minus"
	}

	for statementToTransfer.Next() {
		var rowUser entities.Users

		errUpdate := statementToTransfer.Scan(&rowUser.Balance)
		if errUpdate != nil {
			return -1, err, "error scan"
		}
		updateBalance, errIns := db.Prepare("UPDATE users SET balance=? WHERE no_telp=?")
		if errIns != nil {
			return -1, errIns, "error exec"
		}
		succesUpdate, errUpd := updateBalance.Exec((rowUser.Balance + transfer.Amount), toId)
		if errUpd != nil {
			return -1, errIns, "error update"
		} else {
			row, _ := succesUpdate.RowsAffected()
			if row > 0 {
				return 0, nil, "transfer success done"
			}
		}

	}

	return 0, nil, ress
}
