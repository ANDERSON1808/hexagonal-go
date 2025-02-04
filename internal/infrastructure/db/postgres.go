package db

import (
	"fmt"
	"log"
	"os"
	"time"

	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewPostgresDB() *PostgresDB {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := connectDB(gormpostgres.Open(connect))
	if err != nil {
		log.Fatalf("❌ Error conectando a PostgreSQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Error conectando a PostgreSQL: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("❌ Error conectando a PostgreSQL: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	log.Println("✅ Conectado a PostgreSQL con GORM")

	return &PostgresDB{DB: db}
}

func connectDB(dialect gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
