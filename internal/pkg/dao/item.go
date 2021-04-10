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

const MagicToken = "$$"

func concatPattern(chatId, tag, title string) string {
	tag = strings.ReplaceAll(tag, "-", MagicToken)
	title = strings.ReplaceAll(title, "-", MagicToken)
	return fmt.Sprintf("ah-scut-%s-%s-%s", chatId, tag, title)
}

func parsePattern(key string) (chatId int64, tag, title string) {
	sp := strings.Split(key, "-")
	chatId, _ = xnumber.ParseInt64(sp[2], 10)
	tag = strings.ReplaceAll(sp[3], MagicToken, "-")
	title = strings.ReplaceAll(sp[4], MagicToken, "-")
	return
}

func GetOldItems(chatId int64) ([]*model.PostItem, bool) {
	pattern := concatPattern(xnumber.I64toa(chatId), "*", "*")
	keys, err := database.Redis().Keys(context.Background(), pattern).Result()
	if err != nil {
		return nil, false
	}

	items := make([]*model.PostItem, len(keys))
	for idx := range items {
		_, tag, title := parsePattern(keys[idx])
		items[idx] = &model.PostItem{Type: tag, Title: title}
	}
	return items, true

}

func SetOldItems(chatId int64, items []*model.PostItem) bool {
	pattern := concatPattern(xnumber.I64toa(chatId), "*", "*")
	_, err := xredis.DelAll(database.Redis(), context.Background(), pattern)
	if err != nil {
		return false
	}

	keys := make([]string, 0)
	values := make([]string, 0)
	for _, item := range items {
		id := xnumber.I64toa(chatId)
		pattern := concatPattern(id, item.Type, item.Title)
		keys = append(keys, pattern)
		values = append(values, id)
	}

	_, err = xredis.SetAll(database.Redis(), context.Background(), keys, values)
	return err == nil
}
