package endpointcount

type StatisticsService interface {
	IncrementCount(endpoint string) error
	GetStatistics() ([]Statistics, error)
	GetTotalCountForEndpoint(endpoint string) (int, error)
}

type statisticsService struct {
	statisticsRepository StatisticsRepository
}

func NewStatisticsService(statisticsRepository StatisticsRepository) StatisticsService {
	return &statisticsService{
		statisticsRepository: statisticsRepository,
	}
}

func (s *statisticsService) IncrementCount(endpoint string) error {
	err := s.statisticsRepository.IncrementCount(endpoint)
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

func (s *statisticsService) GetTotalCountForEndpoint(endpoint string) (int, error) {
	totalCount, err := s.statisticsRepository.GetTotalCountForEndpoint(endpoint)
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}
