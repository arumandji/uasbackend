package services

import (
	"uas_backend/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

type ReportService interface {
	GetStatistics() (bson.M, error)
}

type reportService struct {
	achMongoRepo repositories.AchievementMongoRepository
}

func NewReportService(achRepo repositories.AchievementMongoRepository) ReportService {
	return &reportService{achMongoRepo: achRepo}
}

// GetStatistics aggregates achievement data from MongoDB
func (s *reportService) GetStatistics() (bson.M, error) {
	stats, err := s.achMongoRepo.AggregateStats()
	if err != nil {
		return nil, err
	}
	return stats, nil
}
