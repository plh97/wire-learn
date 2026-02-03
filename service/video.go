package service

import (
	"github.com/google/wire"
	"github.com/plh97/wire-learn/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VideoService struct {
	db     *gorm.DB
	logger *zap.Logger
}

var NewVideoServiceProvider = wire.NewSet(NewVideoService)

func NewVideoService(db *gorm.DB, logger *zap.Logger) *VideoService {
	return &VideoService{
		db:     db,
		logger: logger,
	}
}

func (s *VideoService) GetVideo(id int) (*model.Video, error) {
	var video model.Video
	err := s.db.First(&video, id).Error
	if err != nil {
		s.logger.Error("Error fetching video", zap.Error(err))
		return nil, err
	}
	s.logger.Info("Video data fetched", zap.Any("data", video))
	return &video, nil
}
