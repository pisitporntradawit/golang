package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type ConfigDB struct {
	DatabaseUri string
}

func Config() *pgxpool.Pool {
	cfg, _ := loadConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := pgxpool.New(ctx, cfg.DatabaseUri)
	if err != nil {
		log.Fatalf("❌ เชื่อมต่อไม่สำเร็จ: %v", err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("❌ Ping ไม่สำเร็จ: %v", err)
	}
	fmt.Println("✅ เชื่อมต่อ PostgreSQL สำเร็จ!")
	return conn
}

// ตรวจสอบ ENV ว่าอยู่ที่ Dev หรือ Production
func loadConfig() (*ConfigDB, error) {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Fail Load .env")
		}
	}
	return &ConfigDB{
		DatabaseUri: os.Getenv("DATABASE_URI_SUPA"),
	}, nil
}
