package models
import "time"
type LogEntry struct {
	ID        uint `gorm:"primaryKey"`
	IP        string
	Email     string
	Method    string
	Path      string
	Status    int
	Message   string
	Duration  string
	Timestamp time.Time
}
