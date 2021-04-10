package model

import (
	"github.com/Aoi-hosizora/ahlib-db/xgorm"
	"time"
)

type User struct {
	Id     uint32 `gorm:"primary_key; auto_increment"`
	ChatID int64  `gorm:"not_null; unique_index:uk_chat_id"`

	xgorm.GormTime2
	DeletedAt *time.Time `gorm:"default:'1970-01-01 00:00:01'; unique_index:uk_chat_id"`
}

/*
CREATE TABLE IF NOT EXISTS "tbl_user" (
    "id"         integer primary key autoincrement,
    "chat_id"    bigint,
    "created_at" datetime,
    "updated_at" datetime,
    "deleted_at" datetime DEFAULT '1970-01-01 00:00:01'
);
CREATE UNIQUE INDEX uk_chat_id ON "tbl_user" (chat_id, deleted_at);
*/
