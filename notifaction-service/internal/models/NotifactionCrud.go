package models

type Notification struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text;not null"`
	ImageURL    string `gorm:"type:varchar(500)"`
	CreatedAt   int64  `gorm:"autoCreateTime"`
	UpdatedAt   int64  `gorm:"autoUpdateTime"`
}

//sdfgdpjdfgjdijdsjgjdfiojgsdifjgge
//rigbnaepiunbeobnenbunbnhgeogoenobjenbjanejbanjenjnjbanjdrbn
