package postgress

import (
	"database/sql"
	"github.com/batmanboxer/mockchatapp/models"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostGres() (*Postgres, error) {
	connStr := "user=postgres dbname=mockchatapp password=mysecretpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	execute := `CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );`
	_, err = db.Exec(execute)
	if err != nil {
		return nil, err
	}
	return &Postgres{
		db: db,
	}, nil
}

func (postgres *Postgres) AddAccount(signUpData models.SignUpData) error {
	addAccountQuery := `INSERT INTO users(name,email,password)VAlUES($1,$2,$3)`
	_, err := postgres.db.Exec(addAccountQuery, signUpData.Name, signUpData.Email, signUpData.Password)

	if err != nil {
		return err
	}

	return nil
}
