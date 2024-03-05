package db

import (
	"dsched/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database interface {
	GetTextByID(id int) (string, error)
	StoreText(text string) (uint, error)
}

type textModel struct {
	gorm.Model
	Text string
}

type GormDatabase struct {
	db *gorm.DB
}

func (g *GormDatabase) GetTextByID(id int) (string, error) {
	var text textModel
	err := g.db.First(&text, id).Error
	if err != nil {
		return "", err
	}
	return text.Text, nil
}

func (g *GormDatabase) StoreText(text string) (uint, error) {
	model := textModel{Text: text}
	err := g.db.Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

func NewDatabase(config *config.Config) (*GormDatabase, error) {
	db, err := gorm.Open(sqlite.Open(config.DB.URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&textModel{})
	if err != nil {
		return nil, err
	}
	return &GormDatabase{db: db}, nil
}
