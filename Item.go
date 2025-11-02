package youtubestorage

type Item struct {
	ID    int `gorm:"autoIncrement;primaryKey"`
	Query string
	Type  int
}
