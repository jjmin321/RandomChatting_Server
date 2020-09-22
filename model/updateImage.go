package model

import (
	"RandomChatting_Server/database"
)

// UpdateImage : 멤버 찾기
func UpdateImage(Name, Pw, fileName string) error {
	Member := &database.Member{}
	err := database.DB.Model(Member).Where("Name = ? AND Pw = ?", Name, Pw).Update("image", fileName).Error
	return err
}
