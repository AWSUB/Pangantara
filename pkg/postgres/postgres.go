package postgres

import (
	"fmt"
	"log"
	"sppg-backend/config"
	"sppg-backend/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	cfg := config.AppConfig

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	DB = db
	log.Println("Koneksi PostgreSQL (GORM) berhasil!")
}

func Migrate() {

	DB.Exec("SET session_replication_role = replica;")

	err := DB.AutoMigrate(
		&entity.User{},
		&entity.SPPG{},
		&entity.Supplier{},
		&entity.Product{},
		&entity.Stock{},
		&entity.Order{},
		&entity.OrderDetail{},
		&entity.Transaction{},
	)


	DB.Exec("SET session_replication_role = DEFAULT;")

	if err != nil {
		log.Fatalf("AutoMigrate gagal: %v", err)
	}
	log.Println("AutoMigrate berhasil!")
}
