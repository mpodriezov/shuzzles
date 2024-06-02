package data

import "time"

type SessionUser struct {
	Id        int
	Username  string
	Email     string
	SessionId string
	ExpiresOn time.Time
	Role      byte
}

type SessionModel struct {
	SessionId string
	UserId    uint64
	ExpiresOn time.Time
}

func (d *Dal) FindUserBySession(sessionId string) *SessionUser {
	sql := `SELECT u.id, u.username, u.email, u.role, s.session_id, s.expires_on FROM users u JOIN sessions s ON u.id = s.user_id WHERE s.session_id = ?;`
	row := d.DB.QueryRow(sql, sessionId)
	u := SessionUser{}
	err := row.Scan(&u.Id, &u.Username, &u.Email, &u.Role, &u.SessionId, &u.ExpiresOn)
	if err != nil {
		return nil
	}
	if u.ExpiresOn.Before(time.Now()) {
		d.DeleteSession(sessionId)
		return nil
	}
	return &u
}

func (d *Dal) DeleteSession(sessionId string) {
	sql := `DELETE FROM sessions WHERE session_id = ?;`
	_, err := d.DB.Exec(sql, sessionId)
	if err != nil {
		panic(err)
	}
}

func (d *Dal) CreateSession(sessionId string, userId uint64, expiresAt time.Time) *SessionModel {
	sql := `INSERT INTO sessions (session_id, user_id, expires_on) VALUES (?, ?, ?);`
	_, err := d.DB.Exec(sql, sessionId, userId, expiresAt)
	if err != nil {
		panic(err)
	}
	return &SessionModel{UserId: userId, SessionId: sessionId, ExpiresOn: expiresAt}
}
