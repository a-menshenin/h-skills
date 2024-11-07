package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

type MyWaitGroup struct {
	res int32
}

func NewMyWaitGroup() *MyWaitGroup {
	return &MyWaitGroup{res: 0}
}

func (wg *MyWaitGroup) Add(delta int) {
	atomic.AddInt32(&wg.res, 1)
}

func (wg *MyWaitGroup) Done() {
	atomic.AddInt32(&wg.res, -1)
}

func (wg *MyWaitGroup) Wait() {
	for {
		if atomic.LoadInt32(&wg.res) == 0 {
			break
		}
		runtime.Gosched()
	}
}

func task(id int, wg *MyWaitGroup) {
    // Имитация работы
    fmt.Printf("Task %d started\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Task %d completed\n", id)

    // Завершаем задачу
    wg.Done()
}

func main() {
    // Создаем свой WaitGroup
    wg := NewMyWaitGroup()

    // Создаем и запускаем несколько горутин
    for i := 1; i <= 3; i++ {
        wg.Add(1) // Увеличиваем счетчик на 1 для каждой горутины
        go task(i, wg)
    }

    // Ожидаем завершения всех горутин
    wg.Wait()

    fmt.Println("All tasks completed.")
}

