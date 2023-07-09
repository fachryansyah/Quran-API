package surah

type SurahRequest struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

type SurahResponse struct {
	TotalPage   int     `json:"total_page"`
	CurrentPage int     `json:"current_page"`
	TotalData   int     `json:"total_data"`
	StatusCode  int     `json:"status_code"`
	Data        []Surah `json:"data"`
}
