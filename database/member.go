package database

import "time"

// Member 멤버 관리 테이블
type Member struct {
	Idx       uint      `gorm:"primary_key; auto_increment:true" json:"idx"`
	ID        string    `gorm:"type:varchar(255);not null; unique;" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null; unique;" json:"name"`
	Pw        string    `gorm:"type:varchar(255);not null" json:"pw"`
	Image     string    `gorm:"type:varchar(255);" json:"image"`
	Megaphone uint      `gorm:"not null" sql:"DEFAULT:0" json:"megaphone"`
	JoinedAt  time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp" json:"joined_at"`
}
