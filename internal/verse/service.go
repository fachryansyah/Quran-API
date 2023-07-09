package verse

import (
	"errors"
	"quranapi/pkg/sqlite"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetVerseBySurah(req VerseRequest) (VerseResponse, error) {
	var verse []Verse

	var statusCode int
	var err error

	db := sqlite.Connect()
	query := db.Find(&verse, "surahId = ?", req.SurahID)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			statusCode = fiber.StatusNotFound
			err = gorm.ErrRecordNotFound
		}
	}

	return VerseResponse{
		Data:       verse,
		StatusCode: statusCode,
	}, err
}
