package verse

type Verse struct {
	ID      int    `gorm:"column:id;type:int;primaryKey" json:"id"`
	SurahID int    `gorm:"column:surahId;type:int" json:"surah_id"`
	VerseNo int    `gorm:"column:verseNo;type:int" json:"verse_no"`
	Arabic  string `gorm:"column:arabicText;type:string" json:"arabic"`
	Indo    string `gorm:"column:indoText;type:string" json:"indo"`
	Latin   string `gorm:"column:readText;type:string" json:"latin"`
}

func (Verse) TableName() string {
	return "verse"
}
