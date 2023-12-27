package infra

import (
	"blog/configs"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	conn *pgxpool.Pool
}

func NewPostgresDB(c *configs.DBConfig, logger log.Logger) (*PostgresDB, func()) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?charset=utf8&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Name)
	conn, err := pgxpool.New(
		context.Background(),
		dsn,
	)
	if err != nil {
		panic("cannot connect to db")
	}
	cleanup := func() {
		logger.Log(log.LevelInfo, "closing db connection")
		conn.Close()
	}
	return &PostgresDB{
		conn: conn,
	}, cleanup
}

func (db *PostgresDB) Conn() *pgxpool.Pool {
	return db.conn
}