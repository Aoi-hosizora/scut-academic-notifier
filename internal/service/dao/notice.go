package dao

import (
	"context"
	"fmt"
	"github.com/Aoi-hosizora/ahlib-db/xredis"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/database"
	"strings"
)

const magicToken = "$$"

func concatPattern(chatID, typ, title string) string {
	typ = strings.ReplaceAll(typ, "-", magicToken)
	title = strings.ReplaceAll(title, "-", magicToken)
	return fmt.Sprintf("scut-notice-%s-%s-%s", chatID, typ, title)
	//                                      2  3  4
}

func parsePattern(key string) (chatID int64, typ, title string) {
	sp := strings.Split(key, "-")
	chatID, _ = xnumber.ParseInt64(sp[2], 10)
	typ = strings.ReplaceAll(sp[3], magicToken, "-")
	title = strings.ReplaceAll(sp[4], magicToken, "-")
	return
}

func GetNoticeItems(chatID int64) ([]*model.NoticeItem, error) {
	pattern := concatPattern(xnumber.I64toa(chatID), "*", "*")
	keys, err := database.RedisClient().Keys(context.Background(), pattern).Result()
	if err != nil {
		return nil, err
	}

	items := make([]*model.NoticeItem, 0, len(keys))
	for _, key := range keys {
		_, typ, title := parsePattern(key)
		items = append(items, &model.NoticeItem{Type: typ, Title: title})
	}
	return items, nil
}

func SetNoticeItems(chatID int64, items []*model.NoticeItem) error {
	pattern := concatPattern(xnumber.I64toa(chatID), "*", "*")
	_, err := xredis.DelAll(context.Background(), database.RedisClient(), pattern)
	if err != nil {
		return err
	}

	kvs := make([]interface{}, 0, len(items)*2)
	for _, item := range items {
		idStr := xnumber.I64toa(chatID)
		pattern = concatPattern(idStr, item.Type, item.Title)
		kvs = append(kvs, pattern, idStr)
	}
	err = database.RedisClient().MSet(context.Background(), kvs...).Err()
	return err
}
