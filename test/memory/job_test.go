package unit

import (
	"go-video-encoder-course/domain/entity"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := entity.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	job, err := entity.NewJob("path", "Converted", video)
	require.Nil(t, err)
	require.NotNil(t, job)
}