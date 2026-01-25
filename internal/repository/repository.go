package repository

import (
	"fmt"

	"github.com/frukas/scraperiago/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	gormDB *gorm.DB
}

func NewRepository() (*Repository, error) {

	db, err := gorm.Open(sqlite.Open("Articles.db?_busy_timeout=5000"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// sqlDB, _ := db.DB()
	// sqlDB.SetMaxOpenConns(1)
	// db.Exec("PRAGMA journal_mode=WAL;")

	return &Repository{gormDB: db}, err
}

func (repo *Repository) Save(Article models.Article) error {

	return repo.gormDB.Save(&Article).Error
}

func (repo *Repository) GetAll() ([]models.Article, error) {
	var Articles []models.Article

	err := repo.gormDB.Find(&Articles).Error

	return Articles, err
}

func (repo *Repository) SaveAll(Articles []models.Article) error {

	return repo.gormDB.Clauses(clause.OnConflict{DoNothing: true}).Create(&Articles).Error
}

func (repo *Repository) Migration(objectDB interface{}) {
	repo.gormDB.AutoMigrate(&objectDB)
}

func (repo *Repository) Exists(address string) bool {
	var Article models.Article
	err := repo.gormDB.Select("id").Where("address = ?", address).First(&Article).Error

	if err != nil {
		fmt.Println("New File")
		return false
	}

	fmt.Println("File already exist")
	return true
}
