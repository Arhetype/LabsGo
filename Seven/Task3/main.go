package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	activeConnections sync.WaitGroup
)

func main() {
	port := ":8080"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Ошибка при создании слушателя:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Сервер запущен и слушает на порту", port)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		fmt.Println("\nПолучен сигнал завершения. Завершение работы сервера...")
		cancel()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				fmt.Println("Сервер завершает работу...")
				return
			default:
				fmt.Println("Ошибка при принятии соединения:", err)
				continue
			}
		}

		activeConnections.Add(1)

		go handleConnections(ctx, conn)
	}
}

func handleConnections(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	defer activeConnections.Done()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
		return
	}

	message := string(buffer[:n])
	fmt.Println("Получено сообщение:", message)

	response := "Сообщение получено: " + message
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Ошибка при отправке ответа:", err)
		return
	}
}

func waitForShutdown() {
	activeConnections.Wait()
	fmt.Println("Все соединения завершены. Сервер остановлен.")
	os.Exit(0)
}
