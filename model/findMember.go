package model

import "RandomChatting_Server/database"

// FindMember : 멤버 찾기
func FindMember(Name, Pw string) error {
	Member := &database.Member{}
	err := database.DB.Where("Name = ? AND Pw = ?", Name, Pw).Find(Member).Error
	return err
}
