package services

import (
	"golang-microservices/mvc/domain"
	"golang-microservices/mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)

}
