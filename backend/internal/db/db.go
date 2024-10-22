package db

import (
	"backend/internal/domain"
	"context"
	"database/sql"
	"fmt"
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

func NewdbProvider(db_config string, context context.Context, logger *zap.SugaredLogger) *DbProvider {
	return &DbProvider{
		lg:     *logger,
		config: db_config,
		ctx:    context,
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

	_, err = p.db.ExecContext(p.ctx,
		`
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

		CREATE TABLE messages (
			id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
			user VARCHAR(255) NOT NULL,
			message TEXT NOT NULL,
			timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);`,
	)

	if err != nil {
		return err
	}
	return nil

}

func (p *DbProvider) SelectAllMessages() (domain.MessageHistory, error) {
	rows, err := p.db.QueryContext(p.ctx,
		`SELECT * FROM messages ORDER BY timestamp ASC;`,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var messages domain.MessageHistory

	for rows.Next() {
		var msg domain.Message
		err := rows.Scan(&msg.ID, &msg.From, &msg.Text, &msg.Timestamp)
		if err != nil {
			return domain.MessageHistory{}, fmt.Errorf("failed to scan row: %w", err)
		}
		messages.MessageList = append(messages.MessageList, msg)
	}

	if err := rows.Err(); err != nil {
		return domain.MessageHistory{}, fmt.Errorf("row iteration error: %w", err)
	}

	return messages, nil
}

func (p *DbProvider) InsertMessage(msg domain.Message) error {

	query := `
		INSERT INTO messages (id, user, message, timestamp)
		VALUES (uuid_generate_v4(), $1, $2, CURRENT_TIMESTAMP);
	`
	_, err := p.db.ExecContext(p.ctx, query, msg.From, msg.Text)
	if err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}

	return nil
}
