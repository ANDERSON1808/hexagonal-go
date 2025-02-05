package db

import (
	"ANDERSON1808/hexagonal-go/internal/domain"
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

// UserRepositoryImpl maneja la persistencia de UserEntity
type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *PostgresDB) domain.UserRepository {
	if err := db.DB.AutoMigrate(&UserEntity{}, &AuditLog{}); err != nil {
		log.Fatalf("❌ Error en la migración: %v", err)
	}
	log.Println("✅ Migración completada con éxito")
	return &UserRepositoryImpl{db: db.DB}
}

func (r *UserRepositoryImpl) Save(user *domain.User) error {
	entity := FromDomain(user)
	if err := r.db.Create(entity).Error; err != nil {
		return err
	}
	logAudit(r.db, "User", "CREATE", entity)
	return nil
}

func (r *UserRepositoryImpl) FindByID(id uint) (*domain.User, error) {
	var entity UserEntity
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return entity.ToDomain(), nil
}

func (r *UserRepositoryImpl) FindAll() ([]domain.User, error) {
	var entities []UserEntity
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}

	var users []domain.User
	for _, entity := range entities {
		users = append(users, *entity.ToDomain())
	}
	return users, nil
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	var entity UserEntity
	if err := r.db.First(&entity, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&entity).Error; err != nil {
		return err
	}
	logAudit(r.db, "User", "DELETE", entity)
	return nil
}

func logAudit(db *gorm.DB, entity, action string, data any) {
	jsonData, _ := json.Marshal(data)
	audit := AuditLog{
		EntityName: entity,
		Action:     action,
		Details:    string(jsonData),
	}
	db.Create(&audit)
}
