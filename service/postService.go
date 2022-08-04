package service

import (
	"fmt"
	"log"

	"github.com/aldisaputra17/post_opinia/dataobject"
	"github.com/aldisaputra17/post_opinia/models"
	"github.com/aldisaputra17/post_opinia/repository"
	"github.com/mashingan/smapping"
)

type PostinganService interface {
	Create(p dataobject.PostCreateObj) (models.Postingan, error)
	GetAll() ([]models.Postingan, error)
	Update(p dataobject.PostUpdatedObj) (models.Postingan, error)
	Delete(p models.Postingan)
	FindById(postinganID uint64) models.Postingan
	IsAllowedToEdit(userID string, postID uint64) bool
}

type postService struct {
	postRepository repository.PostinganRepository
}

func NewPostService(postRepo repository.PostinganRepository) PostinganService {
	return &postService{
		postRepository: postRepo,
	}
}

func (service *postService) Create(p dataobject.PostCreateObj) (models.Postingan, error) {
	postCreate := models.Postingan{}
	err := smapping.FillStruct(postCreate, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res, err := service.postRepository.InsertPostingan(postCreate)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *postService) GetAll() ([]models.Postingan, error) {
	return service.postRepository.AllPostingan()
}

func (service *postService) Update(p dataobject.PostUpdatedObj) (models.Postingan, error) {
	postUpdate := models.Postingan{}
	err := smapping.FillStruct(postUpdate, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v :", err)
	}
	res, err := service.postRepository.UpdatePostingan(postUpdate)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *postService) Delete(p models.Postingan) {
	service.postRepository.DeletePostingan(p)
}

func (service *postService) FindById(postinganID uint64) models.Postingan {
	return service.postRepository.FindPostinganByID(postinganID)
}

func (service *postService) IsAllowedToEdit(userID string, postID uint64) bool {
	p := service.postRepository.FindPostinganByID(postID)
	id := fmt.Sprintf("%v", p.User_ID)
	return userID == id
}
