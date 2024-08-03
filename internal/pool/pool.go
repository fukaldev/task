package pool

import (
	"fmt"
	"sync"

	"github.com/fukaldev/task/pkg/taskdef"
)

type Pool struct {
	poolSize int
	appName  string
	tasks    *map[string]taskdef.Callable
}

func NewPool(poolSize int, appName string, tasks *map[string]taskdef.Callable) *Pool {
	pool := new(Pool)
	pool.poolSize = poolSize
	pool.tasks = tasks
	return pool
}

func (p *Pool) CreatePool(wg *sync.WaitGroup) {
	for i := 0; i < p.poolSize; i++ {
		go func(id int, appName string) {
			for {
				fmt.Printf("Task %d is ready\n", id)
				fmt.Printf("Task %d starting to run given function\n", id)
				// task.Call()
				fmt.Printf("Task %d finished\n", id)
				wg.Done()
			}
		}(i, p.appName)
	}
}

func (p *Pool) GetPoolSize() int {
	return p.poolSize
}
