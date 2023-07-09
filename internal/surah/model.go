package surah

type Surah struct {
	ID         uint   `gorm:"column:id;type:int;primaryKey" json:"id"`
	Arabic     string `gorm:"column:arabic;type:string" json:"arabic"`
	Latin      string `gorm:"column:latin;type:string" json:"latin"`
	English    string `gorm:"column:english;type:string" json:"english"`
	Indo       string `gorm:"column:indonesia;type:string" json:"indo"`
	Localtion  int    `gorm:"column:localtion;type:int" json:"localtion"`
	Sajda      int    `gorm:"column:sajda;type:int" json:"sadja"`
	VerseTotal int    `gorm:"column:verseTotal;type:int" json:"verse_total"`
}

func (Surah) TableName() string {
	return "surah"
}
