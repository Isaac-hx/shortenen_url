package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"
	"url_shortener_api/internal/config"

	_ "github.com/lib/pq" // To register the driver.
)

func NewDbConn(ctx context.Context, cfg *config.Config) (*sql.DB, error) {
	//Formatting dns postgres
	dsn := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.DbConfig.DbUser, cfg.DbConfig.DbPassword),
		Host:     net.JoinHostPort(cfg.DbConfig.DbHost, strconv.Itoa(cfg.DbConfig.DbPort)),
		Path:     cfg.DbConfig.DbName,
		RawQuery: "sslmode=disable",
	}

	//open connection to dsn
	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to connect database! :%v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	//Try ping to database
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("[ERROR] Failed to ping database!: %w", err)
	}
	return db, nil
}
