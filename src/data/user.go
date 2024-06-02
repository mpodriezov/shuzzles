package data

import (
	"database/sql"
	"time"
)

const Role_Admin byte = 2
const Role_User byte = 4

type RegistrationModel struct {
	Username     string
	Email        string
	PasswordHash string
	Bio          string
	Role         uint16
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserModel struct {
	Id           uint64
	Username     string
	Email        string
	Bio          string
	PasswordHash string
	Role         byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (d *Dal) GetUserById(id uint64) *UserModel {
	sql := `SELECT id, username, email, bio, created_at, updated_at, password_hash, role FROM users WHERE id = ?;`
	row := d.DB.QueryRow(sql, id)
	u := UserModel{}
	err := row.Scan(&u.Id, &u.Username, &u.Email, &u.Bio, &u.CreatedAt, &u.UpdatedAt, &u.PasswordHash, &u.Role)
	if err != nil {
		panic(err)
	}
	return &u
}

func (d *Dal) FindUserByUsername(username string) (*UserModel, error) {
	sqlStr := `SELECT id, username, email, bio, created_at, updated_at, password_hash, role FROM users WHERE username = ?;`
	row := d.DB.QueryRow(sqlStr, username)
	u := UserModel{}
	err := row.Scan(&u.Id, &u.Username, &u.Email, &u.Bio, &u.CreatedAt, &u.UpdatedAt, &u.PasswordHash, &u.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No rows returned
		}
		return nil, err // Other error occurred
	}
	return &u, nil
}

func (d *Dal) ListUsers() []*UserModel {
	sql := `SELECT id, username, email, bio, created_at, updated_at, role FROM users;`
	rows, err := d.DB.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := make([]*UserModel, 0)
	for rows.Next() {
		u := UserModel{}
		err := rows.Scan(&u.Id, &u.Username, &u.Email, &u.Bio, &u.CreatedAt, &u.UpdatedAt, &u.Role)
		if err != nil {
			panic(err)
		}
		users = append(users, &u)
	}
	return users
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
		(username, email, password_hash, bio, created_at, updated_at, role)
		VALUES (?,?,?,?,?,?,?)`
	row, err := d.DB.Exec(sql, ru.Username, ru.Email, ru.PasswordHash, ru.Bio, ru.CreatedAt, ru.UpdatedAt, ru.Role)
	if err != nil {
		panic(err)
	}
	LastInsertId, err := row.LastInsertId()
	if err != nil {
		panic(err)
	}
	return LastInsertId
}

func (d *Dal) UpdateUser(userId uint64, username, email, bio string, role byte) {
	sqlStr := `UPDATE users SET username = ?, email = ?, bio = ?, role = ? WHERE id = ?;`
	_, err := d.DB.Exec(sqlStr, username, email, bio, role, userId)
	if err != nil {
		panic(err)
	}
}

func (d *Dal) DoesUserExistWithEmail(userEmail string) bool {
	sqlStr := `SELECT COUNT(1) FROM users WHERE email = ?;`
	row := d.DB.QueryRow(sqlStr, userEmail)
	count := 0
	err := row.Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func (d *Dal) DoesUserExistWithUsername(username string) bool {
	sqlStr := `SELECT COUNT(1) FROM users WHERE username = ?;`
	row := d.DB.QueryRow(sqlStr, username)
	count := 0
	err := row.Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
