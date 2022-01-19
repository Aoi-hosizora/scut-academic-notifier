package task

import (
	"context"
	"fmt"
	"github.com/Aoi-hosizora/ahlib-web/xtask"
	"github.com/Aoi-hosizora/ahlib-web/xtelebot"
	"github.com/Aoi-hosizora/ahlib/xcolor"
	"github.com/Aoi-hosizora/ahlib/xgopool"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/logger"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"log"
	"runtime"
)

type Task struct {
	task *xtask.CronTask
}

func NewTask(bw *xtelebot.BotWrapper) (*Task, error) {
	cr := cron.New(cron.WithSeconds())
	task := xtask.NewCronTask(cr)
	task.SetJobAddedCallback(func(j *xtask.FuncJob) {
		if config.IsDebugMode() {
			fmt.Printf("[Task-debug] %-45s --> %s (EntryID: %d)\n", xcolor.Blue.Sprintf("%s, %s", j.Title(), j.ScheduleExpr()), j.Funcname(), j.EntryID())
		}
	})
	pool := xgopool.New(int32(10 * runtime.NumCPU()))
	setupLoggers(task, pool)

	// tasks
	jobs := NewJobSet(bw, pool)
	cfg := config.Configs().Task
	_, err := task.AddJobByCronSpec("notifier", cfg.NotifierCron, jobs.notifierJob)
	if err != nil {
		return nil, err
	}

	t := &Task{task: task}
	return t, nil
}

func setupLoggers(task *xtask.CronTask, pool *xgopool.GoPool) {
	l := logger.Logger()
	task.SetJobScheduledCallback(func(j *xtask.FuncJob) {
		fields := logrus.Fields{"module": "task", "type": "execute", "task": j.Title()}
		l.WithFields(fields).Infof("[Task] Executing cron job `%s`", j.Title())
	})
	task.SetPanicHandler(func(j *xtask.FuncJob, v interface{}) {
		fields := logrus.Fields{"module": "task", "type": "panic", "task": j.Title(), "panic": fmt.Sprintf("%v", v)}
		l.WithFields(fields).Errorf("[Task] Job `%s` panics with `%v`", j.Title(), v)
	})
	pool.SetPanicHandler(func(ctx context.Context, v interface{}) {
		f := ctx.Value(ctxFuncnameKey)
		fields := logrus.Fields{"module": "task", "type": "panic", "function": f, "panic": fmt.Sprintf("%v", v)}
		l.WithFields(fields).Errorf("[Task] Function `%s` in job panics with `%v`", f, v)
	})
	task.SetErrorHandler(func(j *xtask.FuncJob, err error) {
		fields := logrus.Fields{"module": "task", "type": "error", "task": j.Title(), "error": err.Error()}
		l.WithFields(fields).Errorf("[Task] Job `%s` errors with `%v`", j.Title(), err)
	})
}

const ctxFuncnameKey = "funcname" // used by xgopool.Pool

func (t *Task) Start() {
	log.Printf("[Task] Starting %d cron jobs", len(t.task.Jobs()))
	t.task.Cron().Start() // run with goroutine
}

func (t *Task) Finish() {
	log.Printf("[Task] Stopping jobs...")
	<-t.task.Cron().Stop().Done()
	log.Println("[Task] Cron jobs are all finished successfully")
}
