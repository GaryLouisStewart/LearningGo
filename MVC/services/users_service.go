package services

import (
	"github.com/GaryLouisStewart/LearningGo/MVC/domain"
	"github.com/GaryLouisStewart/LearningGo/MVC/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
