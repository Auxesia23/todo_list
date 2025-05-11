package repository

import (
	"context"

	"github.com/Auxesia23/todo_list/internal/models"
	"gorm.io/gorm"
	"time"
)

type LogsRepository interface{
	Create(ctx context.Context, log models.LogEntry) error
	GetAll(ctx context.Context)([]models.LogEntry, error)
}

type LogsRepo struct {
	DB *gorm.DB
}

func NewLogsRepository(db *gorm.DB) LogsRepository{
	return &LogsRepo{
		DB: db,
	}
}

func (repo *LogsRepo) Create(ctx context.Context, log models.LogEntry) error{
	err := repo.DB.WithContext(ctx).Create(&log).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *LogsRepo) GetAll(ctx context.Context) ([]models.LogEntry, error) {
	var logs []models.LogEntry

	// Ambil waktu saat ini
	now := time.Now()

	// Awal hari (misalnya: 2025-05-11 00:00:00)
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Akhir hari (misalnya: 2025-05-11 23:59:59)
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := repo.DB.WithContext(ctx).
		Where("timestamp >= ? AND timestamp < ?", startOfDay, endOfDay).
		Order("timestamp DESC").
		Find(&logs).Error

	return logs, err
}
