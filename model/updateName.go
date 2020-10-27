package model

import (
	"RandomChatting_Server/database"
)

// UpdateName : 닉네임 수정
func UpdateName(ID, Name string) error {
	Member := &database.Member{}
	err := database.DB.Model(Member).Where("id = ?", ID).Update("name", Name).Error
	return err
}
