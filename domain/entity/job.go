package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Job struct {
	ID string `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OutputBucketPath string `json:"outputBucketPath" valid:"notnull" gorm:"type:varchar(255)"`
	Status string `json:"status" valid:"notnull" gorm:"type:varchar(255)"`
	Video *Video `json:"video" valid:"-"`
	VideoID string `json:"videoId" valid:"-" gorm:"column:video_id;type:uuid;notnull"`
	Error string `json:"error" valid:"-"`
	CreatedAt time.Time `json:"createdAt" valid:"-"`
	UpdatedAt time.Time `json:"updatedAt" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{
		OutputBucketPath: output,
		Status: status,
		Video: video,
	}
	job.prepare()
	err := job.Validate()
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (job *Job) prepare() {
	job.ID = uuid.NewV4().String()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
}

func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job)
	if err != nil {
		return err
	}
	return nil
}