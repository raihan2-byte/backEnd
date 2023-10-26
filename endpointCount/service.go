package endpointcount

type StatisticsService interface {
	IncrementCount(endpoint string, useragent string) error
	GetStatistics() ([]Statistics, error)
	GetUniqueUserAgentsCount() (int, error)
	GetTotalUniqueUserAgents() (int, error)
}

type statisticsService struct {
	statisticsRepository StatisticsRepository
}

func NewStatisticsService(statisticsRepository StatisticsRepository) StatisticsService {
	return &statisticsService{
		statisticsRepository: statisticsRepository,
	}
}

func (s *statisticsService) GetTotalUniqueUserAgents() (int, error) {
	totalUniqueUserAgents, err := s.statisticsRepository.GetTotalUniqueUserAgents()
	if err != nil {
		return 0, err
	}
	return totalUniqueUserAgents, nil
}

func (s *statisticsService) IncrementCount(endpoint string, useragent string) error {
	err := s.statisticsRepository.IncrementCount(endpoint, useragent)
	if err != nil {
		return err
	}

	return nil
}

func (s *statisticsService) GetStatistics() ([]Statistics, error) {
	statistics, err := s.statisticsRepository.GetStatistics()
	if err != nil {
		return nil, err
	}

	return statistics, nil
}

func (s *statisticsService) GetUniqueUserAgentsCount() (int, error) {
	count, err := s.statisticsRepository.GetUniqueUserAgentsCount()
	if err != nil {
		return 0, err
	}

	return count, nil
}
