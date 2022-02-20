package dao

import (
	"github.com/Aoi-hosizora/ahlib-db/xgorm"
	"github.com/Aoi-hosizora/ahlib/xstatus"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/database"
)

func QueryChats() ([]*model.Chat, error) {
	chats := make([]*model.Chat, 0)
	rdb := database.GormDB().Model(&model.Chat{}).Find(&chats)
	if rdb.Error != nil {
		return nil, rdb.Error
	}
	return chats, nil
}

func CreateChat(chatID int64) (xstatus.DbStatus, error) {
	chat := &model.Chat{ChatID: chatID}
	rdb := database.GormDB().Model(&model.Chat{}).Create(chat)
	return xgorm.CreateErr(rdb)
}

func DeleteChat(chatID int64) (xstatus.DbStatus, error) {
	rdb := database.GormDB().Model(&model.Chat{}).Where("chat_id = ?", chatID).Delete(&model.Chat{})
	return xgorm.DeleteErr(rdb)
}
