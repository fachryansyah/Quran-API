package verse

type VerseRequest struct {
	SurahID int `param:"surahId"`
}

type VerseResponse struct {
	TotalData  int     `json:"total_data"`
	StatusCode int     `json:"status_code"`
	Data       []Verse `json:"data"`
}
