package model

import "RandomChatting_Server/database"

// FindMember : 멤버 찾기
func FindMember(ID, Pw string) (*database.Member, error) {
	Member := &database.Member{}
	err := database.DB.Where("ID = ? AND Pw = ?", ID, Pw).Find(Member).Error
	return Member, err
}
