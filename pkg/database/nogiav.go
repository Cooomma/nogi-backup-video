package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type NogiAv struct {
	ID             int `gorm:"primaryKey"`
	AvID           int
	BvID           string
	Author         string
	MemberID       int
	Title          string
	Subtitle       string
	Description    string
	CoverURL       string
	VideoCreatedAt int
	Duration       int
	Idol           string
	Program        string
	SubtitleGroup  string
	StorageKey     string
	CreatedAt      int
	UpdatedAt      int
	DeletedAt      int
}

func ConnectDB(username, password, url, dbName string) (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, url, dbName)
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func IsBilibiliIDExisted(db *gorm.DB, avID int, bvID string) (bool, error) {
	var nogiAv NogiAv
	if avID != 0 {
		result := db.Where("av_id = ?", avID).First(&nogiAv)
		return result.RowsAffected > 0, result.Error
	}
	if len(bvID) != 0 {
		result := db.Where("bv_id = ?", bvID).First(&nogiAv)
		return result.RowsAffected > 0, result.Error
	}
	return false, fmt.Errorf("neither avID nor bvID should not be empty")
}

func UpsertVideo(db *gorm.DB, avID int, bvID string, video Video) error {
	var nogiav NogiAv
	result := db.FirstOrCreate(&nogiav, NogiAv{
		AvID:           avID,
		BvID:           bvID,
		Author:         video.Author,
		MemberID:       video.MemberID,
		Title:          video.Title,
		Subtitle:       video.Subtitle,
		Description:    video.Description,
		CoverURL:       video.CoverURL,
		VideoCreatedAt: video.VideoCreatedAt,
		Duration:       video.Duration,
		Idol:           video.Idol,
		Program:        video.Program,
		SubtitleGroup:  video.SubtitleGroup,
	})
	return result.Error
}

func UpdateVideoStorageKey(db *gorm.DB, avID int, bvID string, storageKey string) (bool, error) {
	if avID != 0 {
		result := db.Model(&NogiAv{}).Where("av_id = ?", avID).Update("storage_key", storageKey)
		return result.RowsAffected > 0, result.Error
	} else if len(bvID) != 0 {
		result := db.Model(&NogiAv{}).Where("bv_id = ?", bvID).Update("storage_key", storageKey)
		return result.RowsAffected > 0, result.Error
	}
	return false, fmt.Errorf("Neither avID nor bvID are empty")
}
