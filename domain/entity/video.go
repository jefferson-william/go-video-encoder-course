package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID string `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	ResourceId string `json:"resourceId" valid:"notnull" gorm:"type:varchar(255)"`
	FilePath string `json:"filePath" valid:"notnull" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"createdAt" valid:"-"`
	Jobs []*Job `json:"jobs" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}
	return nil
}