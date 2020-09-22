package model

import "RandomChatting_Server/database"

// FindMember : 멤버 찾기
func FindMember(Name, Pw string) (*database.Member, error) {
	Member := &database.Member{}
	err := database.DB.Where("Name = ? AND Pw = ?", Name, Pw).Find(Member).Error
	return Member, err
}
