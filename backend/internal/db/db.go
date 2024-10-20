package db

import (
	"context"
	"database/sql"
	"log"

	"go.uber.org/zap"
	_ "modernc.org/sqlite"
)

type DbProvider struct {
	db     *sql.DB
	config string
	ctx    context.Context
	lg     zap.SugaredLogger
}

func NewdbProvider(db_config string, ctxt context.Context, logger *zap.SugaredLogger) *DbProvider {
	return &DbProvider{
		lg:     *logger,
		config: db_config,
		ctx:    ctxt,
	}
}

func (p *DbProvider) InitDatabase() error {
	var err error

	p.db, err = sql.Open("postgres", p.config)
	if err != nil {
		panic(err)
	}
	p.lg.Info("Connection to db successful")
	defer p.db.Close()

	_, err = p.db.ExecContext(
		context.Background(),
		`CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
	)

	if err != nil {
		return err
	}
	return nil

}

type User struct {
	Email          string
	PasswordHash   string
	LastConnection string
}

func (p *DbProvider) CreateUser(u User) {

	insertUsers := `INSERT INTO users(email, pwhash, ltconnect) VALUES (?, ?, ?)`
	statement, err := p.db.Prepare(insertUsers)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(u.Email, u.PasswordHash, u.LastConnection)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (p *DbProvider) SelectUserByEmail(userEmail string) (User, error) {
	rows, err := p.db.QueryContext(
		context.Background(),
		`SELECT * FROM users WHERE email=?;`, userEmail,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var user User
	for rows.Next() {
		if err := rows.Scan(
			&user.Email, &user.PasswordHash, &user.LastConnection,
		); err != nil {
			return user, err
		}
	}
	p.lg.Info("User retrieved by email: ")
	return user, nil
}
