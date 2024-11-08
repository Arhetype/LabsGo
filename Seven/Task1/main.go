package main

import (
	"fmt"
	"net"
	"os"
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

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии соединения:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

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
