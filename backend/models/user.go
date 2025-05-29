package models

import (
	"database/sql"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, username, role FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func DeleteUser(db *sql.DB, id int64) error {
	_, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}

func UpdateUser(db *sql.DB, id int64, username, role string) error {
	_, err := db.Exec(`UPDATE users SET username=$1, role=$2 WHERE id=$3`, username, role, id)
	return err
}
func UpdatePassword(db *sql.DB, userID int, newHash string) error {
	_, err := db.Exec(`UPDATE users SET password_hash = $1 WHERE id = $2`, newHash, userID)
	return err
}
