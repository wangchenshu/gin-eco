package model

// Host
type Host struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

// EcoWords
type EcoWords struct {
	ID     int    `json:"id" gorm:"column:id"`
	Author string `json:"author" gorm:"type:varchar(32)"`
	Words  string `json:"words" gorm:"type:varchar(200)"`
}

// EcoWisdomAdages
type EcoWisdomAdages struct {
	ID     int    `json:"id" gorm:"column:id"`
	Author string `json:"author" gorm:"type:varchar(32)"`
	Words  string `json:"words" gorm:"type:varchar(200)"`
}

// EcoInspirationals
type EcoInspirationals struct {
	ID     int    `json:"id" gorm:"column:id"`
	Author string `json:"author" gorm:"type:varchar(32)"`
	Words  string `json:"words" gorm:"type:varchar(200)"`
}

// EcoPhorisms
type EcoPhorisms struct {
	ID     int    `json:"id" gorm:"column:id"`
	Author string `json:"author" gorm:"type:varchar(32)"`
	Words  string `json:"words" gorm:"type:varchar(200)"`
}
