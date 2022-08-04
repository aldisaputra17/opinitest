package repository

import (
	"github.com/aldisaputra17/post_opinia/models"
	"gorm.io/gorm"
)

type PostinganRepository interface {
	InsertPostingan(p models.Postingan) (models.Postingan, error)
	UpdatePostingan(p models.Postingan) (models.Postingan, error)
	DeletePostingan(p models.Postingan)
	AllPostingan() ([]models.Postingan, error)
	FindPostinganByID(postinganID uint64) models.Postingan
}

type postinganConnection struct {
	connection *gorm.DB
}

func NewPostinganRepository(dbConn *gorm.DB) PostinganRepository {
	return &postinganConnection{
		connection: dbConn,
	}
}

func (db *postinganConnection) InsertPostingan(p models.Postingan) (models.Postingan, error) {
	err := db.connection.Create(&p).Error
	if err != nil {
		return p, err
	}
	return p, nil
}

func (db *postinganConnection) UpdatePostingan(p models.Postingan) (models.Postingan, error) {
	err := db.connection.Create(&p).Error
	if err != nil {
		return p, err
	}
	return p, nil
}

func (db *postinganConnection) DeletePostingan(p models.Postingan) {
	db.connection.Delete(&p)
}

func (db *postinganConnection) FindPostinganByID(postinganID uint64) models.Postingan {
	var post models.Postingan
	db.connection.Where("id = ?", postinganID).Joins("PostType").Joins("User").Find(&post)
	return post
}

func (db *postinganConnection) AllPostingan() ([]models.Postingan, error) {
	var postingans []models.Postingan
	err := db.connection.Find(&postingans).Error
	if err != nil {
		return postingans, err
	}
	return postingans, nil
}
