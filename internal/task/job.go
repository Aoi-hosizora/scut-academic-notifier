package task

import (
	"context"
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"github.com/Aoi-hosizora/ahlib/xgopool"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/logger"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/service"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/service/dao"
	"gopkg.in/tucnak/telebot.v2"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type JobSet struct {
	bw   *xtelebot.BotWrapper
	pool *xgopool.GoPool

	notifierFlag int32
}

func NewJobSet(bw *xtelebot.BotWrapper, pool *xgopool.GoPool) *JobSet {
	return &JobSet{bw: bw, pool: pool}
}

func (j *JobSet) checkExclusive(flag *int32) (allow bool, release func()) {
	if !atomic.CompareAndSwapInt32(flag, 0, 1) {
		return false, nil
	}
	return true, func() {
		*flag = 0
	}
}

func (j *JobSet) foreachChat(chats []*model.Chat, fn func(chat *model.Chat)) {
	wg := sync.WaitGroup{}
	for _, chat := range chats {
		wg.Add(1)
		chat := chat
		ctx := context.WithValue(context.Background(), ctxFuncnameKey, "foreachChat")
		j.pool.CtxGo(ctx, func(_ context.Context) {
			defer wg.Done()
			fn(chat)
		})
	}
	wg.Wait()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (j *JobSet) notifierJob() {
	ok, release := j.checkExclusive(&j.notifierFlag)
	if !ok {
		return
	}
	defer release()
	chats, _ := dao.QueryChats()
	if len(chats) == 0 {
		return
	}

	// sleep randomly and get new items
	time.Sleep(time.Duration(rand.Intn(int(config.Configs().Task.NotifierTimeNoise))) * time.Second)
	newItems, _ := service.GetNoticeItems()
	if len(newItems) == 0 {
		return
	}

	// foreach chat
	j.foreachChat(chats, func(chat *model.Chat) {
		// get old items and calc diff
		oldItems, err := dao.GetNoticeItems(chat.ChatID)
		if err != nil {
			return
		}
		logger.Logger().Infof("Get old items: #%d | %d", len(oldItems), chat.ChatID)
		diff := model.DiffNoticeItemSlice(newItems, oldItems)
		logger.Logger().Infof("Get diff items: #%d | %d", len(diff), chat.ChatID)
		if len(diff) == 0 {
			return
		}

		// update old items
		err = dao.SetNoticeItems(chat.ChatID, newItems)
		if err != nil {
			return
		}
		logger.Logger().Infof("Set new items: #%d | %d", len(newItems), chat.ChatID)

		// format and send
		formatted := service.FormatNoticeItems(diff, true)
		if formatted == "" {
			return
		}
		dest, err := j.bw.Bot().ChatByID(xnumber.I64toa(chat.ChatID))
		if err == nil {
			j.bw.RespondSend(dest, formatted, telebot.ModeMarkdown)
		}
	})
}
