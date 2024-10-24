package db

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type DbProvider struct {
	pool *pgxpool.Pool
	ctx  context.Context
	lg   *zap.SugaredLogger
}

func NewDbProvider(connString string, logger *zap.SugaredLogger) (*DbProvider, error) {

	logger.Debug("Postgres connection string: ", connString)

	dbCtx := context.Background()

	pool, err := pgxpool.Connect(dbCtx, connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = pool.Ping(dbCtx)
	if err != nil {
		return nil, fmt.Errorf("db ping failed: %v", err)
	}

	return &DbProvider{
		pool: pool,
		ctx:  dbCtx,
		lg:   logger}, nil
}

func (d *DbProvider) Close() {
	d.pool.Close()
}

// SetupDatabase initializes the messages table.
func (d *DbProvider) SetupDb() error {

	_, err := d.pool.Exec(context.Background(),
		`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`+
			`CREATE TABLE IF NOT EXISTS messages (
            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
            username VARCHAR(255) NOT NULL,
            message TEXT NOT NULL,
            timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`)
	if err != nil {
		return fmt.Errorf("failed to create messages table: %v", err)
	}
	return nil
}

// Retrieves all messages ordered by timestamp.
func (d *DbProvider) SelectAllMessages() (domain.MessageHistory, error) {

	rows, err := d.pool.Query(d.ctx,
		"SELECT id, username, message, timestamp FROM messages ORDER BY timestamp")
	if err != nil {
		return domain.MessageHistory{}, fmt.Errorf("failed to fetch messages: %v", err)
	}
	defer rows.Close()

	var messages domain.MessageHistory
	for rows.Next() {
		var msg domain.Message
		if err := rows.Scan(&msg.ID, &msg.From, &msg.Text, &msg.Timestamp); err != nil {
			return domain.MessageHistory{}, fmt.Errorf("failed to scan message: %v", err)
		}
		messages.MessageList = append(messages.MessageList, msg)
	}

	return messages, nil
}

// Inserts a new message into the messages table.
func (d *DbProvider) InsertMessage(msg domain.InsertMessage) (string, error) {

	var id string
	err := d.pool.QueryRow(d.ctx,
		`INSERT INTO messages (id, username, message, timestamp) 
		VALUES (gen_random_uuid(), $1, $2, CURRENT_TIMESTAMP) RETURNING id`,
		msg.From, msg.Text).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert message: %v", err)
	}
	return id, nil
}
