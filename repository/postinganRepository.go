package repository

import (
	"github.com/aldisaputra17/post_opinia/models"
	"gorm.io/gorm"
)

type PostinganRepository interface {
	InsertPostingan(p models.Postingan) models.Postingan
	UpdatePostingan(p models.Postingan) models.Postingan
	DeletePostingan(p models.Postingan)
	AllPostingan() []models.Postingan
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

func (db *postinganConnection) InsertPostingan(p models.Postingan) models.Postingan {
	db.connection.Create(&p)
	db.connection.Preload("User").Find(&p)
	return p
}

func (db *postinganConnection) UpdatePostingan(p models.Postingan) models.Postingan {
	db.connection.Create(&p)
	return p
}

func (db *postinganConnection) DeletePostingan(p models.Postingan) {
	db.connection.Delete(&p)
}

func (db *postinganConnection) FindPostinganByID(postID uint64) models.Postingan {
	var post models.Postingan
	db.connection.Preload("User").Find(&post, postID)
	return post
}

func (db *postinganConnection) AllPostingan() []models.Postingan {
	var postingans []models.Postingan
	db.connection.Preload("User").Find(&postingans)
	return postingans
}

//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
