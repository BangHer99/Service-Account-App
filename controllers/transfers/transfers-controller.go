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

		if rowUser.Balance >= transfer.Amount {
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

func TransferHistory(db *sql.DB, FromId int, transferH entities.TransferHistory) ([]entities.TransferHistory, error) {

	statementTH, err := db.Query("SELECT t.id, u.name_user, t.amount, t.created_at, t.to_account_telp FROM users u INNER JOIN transfers t ON t.from_account_telp = u.no_telp WHERE t.from_account_telp=? OR t.to_account_telp=? ORDER BY t.id DESC", FromId, FromId)
	if err != nil {
		return nil, err
	}

	var transferHis []entities.TransferHistory
	for statementTH.Next() {
		var rowTransfer entities.TransferHistory
		var temp int
		errTh := statementTH.Scan(&rowTransfer.Id, &rowTransfer.From_account_name, &rowTransfer.Amount, &rowTransfer.Created_at, &temp)
		if errTh != nil {
			return nil, err
		}
		statementconv, errqry := db.Query("SELECT name_user FROM users WHERE no_telp=?", temp)
		if errqry != nil {
			return nil, errqry
		}
		for statementconv.Next() {
			errConv := statementconv.Scan(&rowTransfer.To_account_name)
			if errConv != nil {
				return nil, errConv
			} else {
				transferHis = append(transferHis, rowTransfer)
			}
		}
	}
	return transferHis, nil
}
