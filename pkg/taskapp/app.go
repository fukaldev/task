package taskapp

import (
	"sync"

	"github.com/fukaldev/task/internal/pool"
	"github.com/fukaldev/task/pkg/taskdef"
)

type App struct {
	appName string
	tasks   map[string]taskdef.Callable
	pool    pool.Pool
}

func NewApp(appName string, threadCount int) *App {
	app := new(App)
	app.appName = appName
	app.tasks = make(map[string]taskdef.Callable)
	app.pool = *pool.NewPool(threadCount, appName, &app.tasks)
	return app
}

func (a *App) RegisterTask(id string, task taskdef.Callable) {
	a.tasks[id] = task
}

func (a *App) Start() {
	wg := sync.WaitGroup{}
	a.pool.CreatePool(&wg)
	wg.Wait()
}
