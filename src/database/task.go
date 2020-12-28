package database

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-db/xredis"
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
	conn, err := Redis.Dial()
	if err != nil {
		return nil, false
	}
	defer conn.Close()

	pattern := getOldDataPattern(xnumber.I64toa(chatId), "*", "*")
	keys, err := redis.Strings(conn.Do("KEYS", pattern))
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
	conn, err := Redis.Dial()
	if err != nil {
		return false
	}
	defer conn.Close()

	pattern := getOldDataPattern(xnumber.I64toa(chatId), "*", "*")
	tot, del, err := xredis.WithConn(conn).DeleteAll(pattern)
	if err != nil || (tot != 0 && del == 0) {
		return false
	}

	keys := make([]string, 0)
	values := make([]string, 0)
	for _, item := range items {
		id := xnumber.I64toa(chatId)
		pattern := getOldDataPattern(id, item.Type, item.Title)
		keys = append(keys, pattern)
		values = append(values, id)
	}

	tot, add, err := xredis.WithConn(conn).SetAll(keys, values)
	return err == nil && (tot == 0 || add >= 1)
}
