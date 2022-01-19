package dao

import (
	"github.com/Aoi-hosizora/ahlib-db/xgorm"
	"github.com/Aoi-hosizora/ahlib/xstatus"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/database"
)

func QueryUsers() []*model.User {
	users := make([]*model.User, 0)
	database.GormDB().Model(&model.User{}).Find(&users)
	return users
}

func CreateUser(chatID int64) xstatus.DbStatus {
	rdb := database.GormDB().Create(&model.User{ChatID: chatID})
	sts, _ := xgorm.CreateErr(rdb)
	return sts
}

func DeleteUser(chatID int64) xstatus.DbStatus {
	rdb := database.GormDB().Model(&model.User{}).Where("chat_id = ?", chatID).Delete(&model.User{})
	sts, _ := xgorm.DeleteErr(rdb)
	return sts
}
