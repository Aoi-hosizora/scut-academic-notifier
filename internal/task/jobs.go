package task

import (
	"context"
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"github.com/Aoi-hosizora/ahlib/xgopool"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/model"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/logger"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/service"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/service/dao"
	"gopkg.in/tucnak/telebot.v2"
	"sync"
)

type JobSet struct {
	bw   *xtelebot.BotWrapper
	pool *xgopool.GoPool
}

func NewJobSet(bw *xtelebot.BotWrapper, pool *xgopool.GoPool) *JobSet {
	return &JobSet{bw: bw, pool: pool}
}

func (j *JobSet) foreachChats(chats []*model.Chat, fn func(chat *model.Chat)) {
	wg := sync.WaitGroup{}
	for _, chat := range chats {
		wg.Add(1)
		chat := chat
		ctx := context.WithValue(context.Background(), ctxFuncnameKey, "foreachChats")
		j.pool.CtxGo(ctx, func(_ context.Context) {
			defer wg.Done()
			fn(chat)
		})
	}
	wg.Wait()
}

func (j *JobSet) notifierJob() error {
	chats, _ := dao.QueryChats()
	if len(chats) == 0 {
		return nil
	}

	// get new items
	newItems, _ := service.GetNoticeItems()
	if len(newItems) == 0 {
		return nil
	}

	// foreach chat
	j.foreachChats(chats, func(chat *model.Chat) {
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

		// render and send
		rendered := service.RenderNoticeItems(diff, true)
		if rendered == "" {
			return
		}
		dest, err := j.bw.Bot().ChatByID(xnumber.I64toa(chat.ChatID))
		if err == nil {
			j.bw.SendTo(dest, rendered, telebot.ModeMarkdown)
		}
	})
	return nil
}
