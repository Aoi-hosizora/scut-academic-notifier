package dao

import (
	"github.com/Aoi-hosizora/ahlib-db/xgorm"
	"github.com/Aoi-hosizora/ahlib/xstatus"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/database"
)

func QueryChats() []*model.Chat {
	users := make([]*model.Chat, 0)
	database.GormDB().Model(&model.Chat{}).Find(&users)
	return users
}

func CreateChat(chatID int64) xstatus.DbStatus {
	chat := &model.Chat{ChatID: chatID}
	rdb := database.GormDB().Model(&model.Chat{}).Create(chat)
	sts, _ := xgorm.CreateErr(rdb)
	return sts
}

func DeleteChat(chatID int64) xstatus.DbStatus {
	rdb := database.GormDB().Model(&model.Chat{}).Where("chat_id = ?", chatID).Delete(&model.Chat{})
	sts, _ := xgorm.DeleteErr(rdb)
	return sts
}
