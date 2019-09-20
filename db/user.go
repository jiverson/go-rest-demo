package db

import (
	"database/sql"

	"github.com/jiverson/go-rest-demo/model"
)

func (db *Database) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	sqlStatement := `
		SELECT id, email, hashed_password
		FROM users
		WHERE (email = $1)
		ORDER BY id
		ASC LIMIT 1`
	if err := db.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Email, &user.HashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (db *Database) CreateUser(user *model.User) error {
	// temp so I can test the values
	// s := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(s)
	// s1 := strconv.Itoa(r.Intn(1000))

	sqlStatement := `
		INSERT INTO users (email, hashed_password, created_at, updated_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id`
	err := db.QueryRow(sqlStatement, user.Email, user.HashedPassword).Scan(&user.ID)

	return err
}
