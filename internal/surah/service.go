package surah

import (
	"quranapi/pkg/sqlite"
)

func GetSurah(req SurahRequest) (SurahResponse, error) {

	if req.Limit < 1 {
		req.Limit = 10
	}

	currentPage := req.Page
	if req.Page < 1 {
		req.Page = 0
		currentPage++
	} else {
		req.Page = (req.Page - 1) * req.Limit
	}

	var surah []Surah
	db := sqlite.Connect()
	db.Limit(req.Limit).Offset(req.Page).Find(&surah)

	var totalSurah int64
	db.Model(&surah).Count(&totalSurah)
	return SurahResponse{
		Data:        surah,
		TotalPage:   int(totalSurah) / req.Limit,
		CurrentPage: currentPage,
		TotalData:   int(totalSurah),
		StatusCode:  200,
	}, nil
}
