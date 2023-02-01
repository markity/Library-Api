package user

import (
	"database/sql"
	"library-api/dao"
)

func TryChangePassword(userID int64, oldPasswordCrypto []byte, newPasswordCrypto []byte) (bool, error) {
	tx, err := dao.DB.Begin()
	if err != nil {
		return false, err
	}

	row := tx.QueryRow("SELECT id FROM user WHERE id = ? AND password_crypto = ? FOR UPDATE", userID, oldPasswordCrypto)
	var discard int64
	err = row.Scan(&discard)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	// 已经锁住, 现在可UPDATE
	_, err = tx.Exec("UPDATE user SET password_crypto = ? WHERE id = ?", newPasswordCrypto, userID)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil

}
