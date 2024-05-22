package data

import (
	"database/sql"
)

type Dal struct {
	DB *sql.DB
}

func (d *Dal) GetUserById(id int) *UserModel {
	sql := `SELECT id, username, email, bio, altitude, longitude, created_at, updated_at FROM users WHERE id = ?;`
	row := d.DB.QueryRow(sql, id)
	u := UserModel{}
	err := row.Scan(&u.Id, &u.Username, &u.Email, &u.Bio, &u.Altitude, &u.Longitude, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return &u
}

func (d *Dal) UserExists(username, email string) bool {
	sql := `SELECT COUNT(1) FROM users WHERE username = ? OR email = ?;`
	row := d.DB.QueryRow(sql, username, email)
	count := 0
	err := row.Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func (d *Dal) NewUser(ru *RegistrationModel) int64 {
	sql := `
		INSERT INTO users
		(username, email, password_hash, bio, created_at, updated_at)
		VALUES (?,?,?,?,?,?)`
	row, err := d.DB.Exec(sql, ru.Username, ru.Email, ru.PasswordHash, ru.Bio, ru.CreatedAt, ru.UpdatedAt)
	if err != nil {
		panic(err)
	}
	LastInsertId, err := row.LastInsertId()
	if err != nil {
		panic(err)
	}
	return LastInsertId
}
