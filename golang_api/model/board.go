package model

type Board struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedBy string    `gorm:"not null" json:"created_by"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	Displays  []Display `gorm:"foreignKey:BoardID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Display struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	BoardID    uint   `gorm:"not null" json:"board_id"`
	RejectedBy string `gorm:"not null" json:"rejected_by"`
}

type NotDisplayUser struct {
	Producer string
}

type NotDisplayBoard struct {
	BoardID uint
}
