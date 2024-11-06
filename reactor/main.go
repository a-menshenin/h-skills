package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

type EventHandler interface {
  HandleEvent(conn net.Conn, message string)
}

type FooHandler struct {}

func (h *FooHandler) HandleEvent(conn net.Conn, message string) {
	// отправляем ответ клиенту
	_, err := conn.Write([]byte("Foo Event handled: " + message + "\n"))
	if err != nil {
		fmt.Errorf("HandleEvent conn.Write err: %w", err)
	}
}

type Reactor struct {
	wg *sync.WaitGroup
	closeCh chan struct{}
}

func NewReactor() *Reactor {
	wg := &sync.WaitGroup{}
	return &Reactor{
		wg: wg,
	}
}

func (r *Reactor) ListenAndServe(port string) error {
	addr, err := net.ResolveTCPAddr("tcp", "localhost" + port)
    if err != nil {
        fmt.Println("Error resolving address:", err)

        return err
    }
	
	listener, err := net.ListenTCP("tcp", addr)
    if err != nil {
        fmt.Println("Error listening:", err)

        return err
    }
	fmt.Println("Server is listening on " + port)

	r.wg.Add(1)
	go func(){
		defer r.wg.Done()
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting:", err.Error())
				os.Exit(1)
			}

			fmt.Println("Connected with", conn.RemoteAddr().String())
			r.wg.Add(1)
			go r.handleRequest(conn)
		}
	}()

	go func() {
		r.wg.Wait()
		listener.Close()
		r.closeCh <- struct{}{}
	}()

	return nil
}

func (r *Reactor) handleRequest(conn net.Conn) {
    defer func() {
		conn.Close()
		r.wg.Done()
	}()

    // читаем данные от клиента
	// buffer := make([]byte, 1024)
	// n, err := conn.Read(buffer)
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("conn.Read err: %w", err).Error())

	// 	return
	// }

	// var msg []byte
	clientMessage, err :=  bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("ReadString err " + err.Error())

		return
	}

	// clientMessage := string(msg)

	fmt.Printf("Received from client: %s\n", clientMessage)

	switch clientMessage {
	case "foo":
		h := FooHandler{}
		h.HandleEvent(conn, clientMessage)
	default:
		return
	}
}

func (r *Reactor) Wait() {
	<-r.closeCh
}

func main() {
  reactor := NewReactor()

  go func() {
    if err := reactor.ListenAndServe(":8080"); err != nil {
      fmt.Println("Ошибка сервера:", err)
    }
  }()

  // Ожидаем завершения обработки соединений
  reactor.Wait()
}
