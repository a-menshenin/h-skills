package main

import (
	"fmt"
	"sync"
	"time"
)

type RWMutex struct{
    readCh chan struct{}
    writeCh chan struct{}
}

func NewRWMutex() *RWMutex {
    readCh := make(chan struct{})
    writeCh := make(chan struct{})
    
    go func (){
        readCh <- struct{}{}
        writeCh <- struct{}{}
    }()
    
    return &RWMutex{
        readCh: readCh,
        writeCh: writeCh,
    }
}

func (m *RWMutex) Lock() {
    <-m.readCh
    <-m.writeCh
}

func (m *RWMutex) Unlock() {
    go func (){
        m.readCh <- struct{}{}
        m.writeCh <- struct{}{}
    }()
}

func (m *RWMutex) RLock() {
    <-m.writeCh
}

func (m *RWMutex) RUnlock() {
    go func (){
        m.writeCh <- struct{}{}
    }()
}

func main() {
    rwMutex := NewRWMutex()
    var wg sync.WaitGroup

	var a int

    // Пример параллельного чтения
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            rwMutex.RLock()
            fmt.Printf("Reader %d reading\n", id)
            time.Sleep(time.Duration(i) * time.Second)
            fmt.Printf("Reader %d done reading\n", id)
			b := a
			_=b
            rwMutex.RUnlock()
        }(i)
    }

    // Пример параллельной записи
    wg.Add(1)
    go func() {
        defer wg.Done()
        time.Sleep(500 * time.Millisecond) // даем читателям начать
        rwMutex.Lock()
        fmt.Println("Writer writing")
        time.Sleep(2 * time.Second)
        fmt.Println("Writer done writing")
		a++
        rwMutex.Unlock()
    }()

    wg.Wait()
    fmt.Println("All operations completed")
}
