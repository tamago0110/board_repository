package model

type User struct {
	ID       string    `gorm:"primary_key" json:"id"`
	Email    string    `gorm:"not null;unique" json:"email" validate:"email"`
	Password string    `gorm:"not null;->:false;<-:create" json:"password"`
	Profile  Profile   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Consumer []Lead    `gorm:"foreignKey:Consumer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Producer []Lead    `gorm:"foreignKey:Producer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Boards   []Board   `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Displays []Display `gorm:"foreignKey:RejectedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Profile struct {
	ID     uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID string `gorm:"not null" json:"user_id"`
	Name   string `gorm:"not null" json:"name"`
	Image  string `json:"image"`
}

type Lead struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Consumer string `gorm:"not null" json:"consumer"`
	Producer string `gorm:"not null" json:"producer"`
}
