package pool

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/fukaldev/task/pkg/taskdef"
	"github.com/redis/go-redis/v9"
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
	pool.appName = appName
	return pool
}

func (p *Pool) CreatePool(wg *sync.WaitGroup, taskQueue *redis.Client) {
	wg.Add(p.poolSize)
	for i := 0; i < p.poolSize; i++ {
		go func(id int, appName string) {
			for {
				fmt.Printf("Task %d is ready\n", id)
				ctx := context.Background()
				for {
					receivedTask, err := taskQueue.BRPop(ctx, 1<<63-1, p.appName).Result()
					if err != nil {
						wg.Done()
						log.Fatal("failed to read task:", err)
					}
					fmt.Printf("Task %d starting to run given function\n", id)
					fmt.Println("Received task", receivedTask)
					fmt.Printf("Task %d finished\n", id)
				}
			}
		}(i, p.appName)
	}
}

func (p *Pool) GetPoolSize() int {
	return p.poolSize
}
