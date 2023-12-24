package endpointcount

import (
	"gorm.io/gorm"
)

type StatisticsRepository interface {
	IncrementCount(endpoint string, useragent string) error
	GetStatistics() ([]Statistics, error)
	GetUniqueUserAgentsCount() (int, error)
	GetTotalUniqueUserAgents() (int, error)
}

type statisticsRepository struct {
	db *gorm.DB
}

func NewStatisticsRepository(db *gorm.DB) StatisticsRepository {
	return &statisticsRepository{
		db: db,
	}
}


func (r *statisticsRepository) GetTotalUniqueUserAgents() (int, error) {
    var totalUniqueUserAgents int64
    err := r.db.Model(&Statistics{}).Distinct("user_agent").Count(&totalUniqueUserAgents).Error
    if err != nil {
        return 0, err
    }
    return int(totalUniqueUserAgents), nil
}

func (r *statisticsRepository) IncrementCount(endpoint string, useragent string) error {

	statistics := Statistics{}
	err := r.db.FirstOrCreate(&statistics, Statistics{Endpoint: endpoint, UserAgent: useragent}).Error
	if err != nil {
		return err
	}

	var uniqueUserAgentCount int64
	err = r.db.Model(&Statistics{}).Distinct("user_agent").Where("endpoint = ?", endpoint).Count(&uniqueUserAgentCount).Error
	if err != nil {
		return err
	}
	statistics.UniqueUserAgent = int(uniqueUserAgentCount)

	statistics.Count++
	statistics.UserAgent = useragent
	// statistics.UniqueUserAgent = uniqueUserAgentCount

	err = r.db.Save(&statistics).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *statisticsRepository) GetStatistics() ([]Statistics, error) {
	var statistics []Statistics
	err := r.db.Find(&statistics).Error
	if err != nil {
		return nil, err
	}

	return statistics, nil
}



func (r *statisticsRepository) GetUniqueUserAgentsCount() (int, error) {
	var count int64
	err := r.db.Model(&Statistics{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *statisticsRepository) GetTotalCountByEndpoint(endpoint string) (int64, error) {
	var totalCount int64

	result := r.db.Model(&Statistics{}).Where("endpoint = ?", endpoint).Select("SUM(count)").Row()
	err := result.Scan(&totalCount)
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}