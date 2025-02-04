package db

import (
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&UserEntity{}, &AuditLog{})
	if err != nil {
		log.Fatalf("❌ Error ejecutando migraciones: %v", err)
	}
	log.Println("✅ Migraciones ejecutadas correctamente")
}
