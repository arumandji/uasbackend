package services

import (
	"context"

	"uas_backend/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type ReportService interface {
	GetAchievementStats() ([]bson.M, error)
}

type reportService struct {
	achMongoRepo *repositories.AchievementMongoRepository
}

func NewReportService(achMongoRepo *repositories.AchievementMongoRepository) ReportService {
	return &reportService{
		achMongoRepo: achMongoRepo,
	}
}

func (s *reportService) GetAchievementStats() ([]bson.M, error) {
	stats, err := s.achMongoRepo.AggregateStats(context.Background())
	if err != nil {
		return nil, err
	}
	return stats, nil
}
