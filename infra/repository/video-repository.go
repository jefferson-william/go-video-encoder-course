package repository

import (
	"fmt"
	"go-video-encoder-course/domain/entity"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type VideoRepository interface {
	Insert(video *entity.Video) (*entity.Video, error)
	Find(id string) (*entity.Video, error)
}

type VideoRepositoryDatabase struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDatabase {
	return &VideoRepositoryDatabase{Db: db}
}

func (repo VideoRepositoryDatabase) Insert(video *entity.Video) (*entity.Video, error) {
	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}
	err := repo.Db.Create(video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (repo VideoRepositoryDatabase) Find(id string) (*entity.Video, error) {
	var video entity.Video
	repo.Db.First(&video, "id = ?", id)
	if video.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}
	return &video, nil
}