package topups

import (
	"be11/account-service-app/entities"
	"database/sql"
)

func TopUp(db *sql.DB, fromId int, toId int, topup entities.TopUps) (int, error, string) {

	var result string
	statementInsertTopup, errTop := db.Prepare("INSERT INTO top_up (from_account_telp ,to_account_telp,amount)VALUES (?,?,?)")
	if errTop != nil {
		return -1, errTop, "error statement"
	}

	resTopUp, err := statementInsertTopup.Exec(fromId, toId, topup.Amount)
	if err != nil {
		return -1, err, "error Exec"
	} else {
		row, err := resTopUp.RowsAffected()
		if err != nil {
			return int(row), err, "err"
		}

	}

	statementAmout, errAm := db.Query("SELECT balance FROM users WHERE no_telp = ?", toId)
	if errAm != nil {
		return -1, errAm, "error minus topup"
	}

	for statementAmout.Next() {
		var rowUser entities.Users
		errUpdate := statementAmout.Scan(&rowUser.Balance)
		if errUpdate != nil {
			return -1, err, "error scan"
		}

		updateBalance, errTop := db.Prepare("UPDATE users SET balance =? WHERE no_telp=?")
		if errTop != nil {
			return -1, errTop, "error exec"
		}
		succesUpdate, errUpd := updateBalance.Exec((rowUser.Balance + topup.Amount), toId)
		if errUpd != nil {
			return -1, errTop, "Error Update"
		} else {
			row, _ := succesUpdate.RowsAffected()
			if row > 0 {
				result = "Top-Up Success"
			}
		}

	}
	return 0, nil, result
}
func ToupUpHistory(db *sql.DB, FromId int, HisTop entities.TopUpHistory) ([]entities.TopUpHistory, error) {

	statementHT, err := db.Query("SELECT t.id , u.name_user, t.amount, t.created_at , t.to_account_telp FROM users u INNER JOIN top_up t ON t.from_account_telp = u.no_telp WHERE t.to_account_telp = ? OR t.from_account_telp =? ORDER BY t.id DESC", FromId, FromId)
	if err != nil {
		return nil, err
	}
	var TopupHis []entities.TopUpHistory
	for statementHT.Next() {
		var rowTopupH entities.TopUpHistory
		var temp int
		errHt := statementHT.Scan(&rowTopupH.Id, &rowTopupH.NameUser, &rowTopupH.Amount, &rowTopupH.CreatedAt, &temp)
		if errHt != nil {
			return nil, err
		}
		statementconv, errqry := db.Query("SELECT name_user FROM users WHERE no_telp=?", temp)
		if errqry != nil {
			return nil, errqry
		}
		for statementconv.Next() {
			errConv := statementconv.Scan(&rowTopupH.To_account_name)
			if errConv != nil {
				return nil, errConv
			} else {
				TopupHis = append(TopupHis, rowTopupH)
			}
		}
	}
	return TopupHis, nil

}
