package main

import "fmt"

type Task struct {
	id       int
	priority string
}

func worker(highPrio, lowPrio <-chan Task, done chan<- bool) {
	b := make(chan struct{})

	go func(){
		for task := range highPrio {
			fmt.Printf("task #%d with high priority handled\n", task.id)
		}

		b <- struct{}{}
	}()

	go func(){
		<-b
		for task := range lowPrio {
			fmt.Printf("task #%d with low priority handled\n", task.id)
		}

		done <- true
		close(b)
	}()
}

func main() {
	highPrio := make(chan Task, 5)
	lowPrio := make(chan Task, 5)
	done := make(chan bool)

	// Запуск воркера
	go worker(highPrio, lowPrio, done)

	// Добавляем задачи в очереди
	for i := 1; i <= 3; i++ {
		highPrio <- Task{id: i, priority: "high"}
	}
	for i := 4; i <= 6; i++ {
		lowPrio <- Task{id: i, priority: "low"}
	}

	// Закрываем каналы после добавления задач
	close(highPrio)
	close(lowPrio)

	// Ожидание завершения работы воркера
	<-done
	close(done)
	fmt.Println("All tasks processed")
}
