package endpointcount

import (
	"gorm.io/gorm"
)

type StatisticsRepository interface {
	IncrementCount(endpoint string) error
	GetStatistics() ([]Statistics, error)
	GetUniqueUserAgentsCount() (int, error)
    GetTotalCountForEndpoint(endpoint string) (int, error)    
}

type statisticsRepository struct {
	db *gorm.DB
}

func NewStatisticsRepository(db *gorm.DB) StatisticsRepository {
	return &statisticsRepository{
		db: db,
	}
}

func (r *statisticsRepository) IncrementCount(endpoint string ) error {
    statistics := Statistics{}
    
    // Cari data statistik berdasarkan endpoint yang diberikan
    err := r.db.Where("endpoint = ?", endpoint).First(&statistics).Error
    if err != nil {
        // Jika data tidak ditemukan, buat data baru
        statistics.Endpoint = endpoint
        statistics.Count = 1

        var uniqueUserAgentCount int64
        err = r.db.Model(&Statistics{}).Where("endpoint = ?", endpoint).Count(&uniqueUserAgentCount).Error
        if err != nil {
            return err
        }
        statistics.UniqueUserAgent = int(uniqueUserAgentCount)

        // Simpan data statistik baru ke dalam tabel
        err = r.db.Create(&statistics).Error
        if err != nil {
            return err
        }
    } else {
        // Jika data sudah ada, tambahkan count-nya
        statistics.Count++
        err = r.db.Save(&statistics).Error
        if err != nil {
            return err
        }
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

func (r *statisticsRepository) GetTotalCountForEndpoint(endpoint string) (int, error) {
    var totalCount int
    result := r.db.Table("statistics").Select("SUM(count) as total_count").Where("endpoint = ?", endpoint).Scan(&totalCount)
    if result.Error != nil {
        return 0, result.Error
    }

    return totalCount, nil
}