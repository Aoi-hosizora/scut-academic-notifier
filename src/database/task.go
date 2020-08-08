package database

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-web/xredis"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/scut-academic-notifier/src/model"
	"github.com/gomodule/redigo/redis"
	"strings"
)

const MagicToken = "$$"

func getOldDataPattern(chatId, tag, title string) string {
	tag = strings.ReplaceAll(tag, "-", MagicToken)
	title = strings.ReplaceAll(title, "-", MagicToken)
	return fmt.Sprintf("ah-scut-%s-%s-%s", chatId, tag, title)
}

func parseOldDataPattern(key string) (chatId int64, tag, title string) {
	sp := strings.Split(key, "-")
	chatId, _ = xnumber.ParseInt64(sp[2], 10)
	tag = strings.ReplaceAll(sp[3], MagicToken, "-")
	title = strings.ReplaceAll(sp[4], MagicToken, "-")
	return
}

func GetOldData(chatId int64) ([]*model.Item, bool) {
	pattern := getOldDataPattern(xnumber.FormatInt64(chatId, 10), "*", "*")
	redisMu.Lock()
	keys, err := redis.Strings(Conn.Do("KEYS", pattern))
	redisMu.Unlock()
	if err != nil {
		return nil, false
	}

	items := make([]*model.Item, len(keys))
	for idx := range items {
		_, tag, title := parseOldDataPattern(keys[idx])
		items[idx] = &model.Item{Type: tag, Title: title}
	}
	return items, true

}

func SetOldData(chatId int64, items []*model.Item) bool {
	pattern := getOldDataPattern(xnumber.FormatInt64(chatId, 10), "*", "*")
	redisMu.Lock()
	tot, del, err := xredis.WithConn(Conn).DeleteAll(pattern)
	redisMu.Unlock()
	if err != nil || (tot != 0 && del == 0) {
		return false
	}

	keys := make([]string, 0)
	values := make([]string, 0)
	for _, item := range items {
		id := xnumber.FormatInt64(chatId, 10)
		pattern := getOldDataPattern(id, item.Type, item.Title)
		keys = append(keys, pattern)
		values = append(values, id)
	}

	redisMu.Lock()
	tot, add, err := xredis.WithConn(Conn).SetAll(keys, values)
	redisMu.Unlock()
	return err == nil && (tot == 0 || add >= 1)
}
