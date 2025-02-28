package unit

import (
	"go-video-encoder-course/domain/entity"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoError(t *testing.T) {
	video := entity.NewVideo()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIdIsNotUuid(t *testing.T) {
	video := entity.NewVideo()
	video.ID = "abc"
	video.ResourceId = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := entity.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceId = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	err := video.Validate()
	require.Nil(t, err)
}