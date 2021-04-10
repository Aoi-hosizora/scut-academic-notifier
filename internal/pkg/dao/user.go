package dao

import (
	"github.com/Aoi-hosizora/ahlib-db/xgorm"
	"github.com/Aoi-hosizora/ahlib/xstatus"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/database"
)

func GetUsers() []*model.User {
	users := make([]*model.User, 0)
	database.DB().Model(&model.User{}).Find(&users)
	return users
}

func GetUser(chatId int64) *model.User {
	user := &model.User{}
	rdb := database.DB().Model(&model.User{}).Where(&model.User{ChatID: chatId}).First(user)
	if rdb.RecordNotFound() {
		return nil
	}
	return user
}

func AddUser(user *model.User) xstatus.DbStatus {
	rdb := database.DB().Create(user)
	sts, _ := xgorm.CreateErr(rdb)
	return sts
}

func DeleteUser(chatId int64) xstatus.DbStatus {
	rdb := database.DB().Model(&model.User{}).Where(&model.User{ChatID: chatId}).Delete(&model.User{})
	sts, _ := xgorm.DeleteErr(rdb)
	return sts
}
