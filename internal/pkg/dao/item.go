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

func concatPattern(chatID, tag, title string) string {
	tag = strings.ReplaceAll(tag, "-", MagicToken)
	title = strings.ReplaceAll(title, "-", MagicToken)
	return fmt.Sprintf("ah-scut-%s-%s-%s", chatID, tag, title)
}

func parsePattern(key string) (chatID int64, tag, title string) {
	sp := strings.Split(key, "-")
	chatID, _ = xnumber.ParseInt64(sp[2], 10)
	tag = strings.ReplaceAll(sp[3], MagicToken, "-")
	title = strings.ReplaceAll(sp[4], MagicToken, "-")
	return
}

func GetPostItems(chatID int64) ([]*model.PostItem, bool) {
	pattern := concatPattern(xnumber.I64toa(chatID), "*", "*")
	keys, err := database.Redis().Keys(context.Background(), pattern).Result()
	if err != nil {
		return nil, false
	}

	items := make([]*model.PostItem, 0, len(keys))
	for _, key := range keys {
		_, tag, title := parsePattern(key)
		items = append(items, &model.PostItem{Type: tag, Title: title})
	}
	return items, true
}

func SetPostItems(chatID int64, items []*model.PostItem) bool {
	pattern := concatPattern(xnumber.I64toa(chatID), "*", "*")
	_, err := xredis.DelAll(database.Redis(), context.Background(), pattern)
	if err != nil {
		return false
	}

	keys := make([]string, 0)
	values := make([]string, 0)
	for _, item := range items {
		id := xnumber.I64toa(chatID)
		pattern = concatPattern(id, item.Type, item.Title)
		keys = append(keys, pattern)
		values = append(values, id)
	}

	_, err = xredis.SetAll(database.Redis(), context.Background(), keys, values)
	return err == nil
}
