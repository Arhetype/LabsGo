package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

type ChatServer struct {
	clients   map[*Client]bool
	broadcast chan []byte
	mu        sync.Mutex
}

func newChatServer() *ChatServer {
	return &ChatServer{
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
	}
}

func (cs *ChatServer) run() {
	for {
		msg := <-cs.broadcast
		cs.mu.Lock()
		for client := range cs.clients {
			select {
			case client.send <- msg:
			default:
				close(client.send)
				delete(cs.clients, client)
			}
		}
		cs.mu.Unlock()
	}
}

func (c *Client) readMessages(cs *ChatServer) {
	defer func() {
		c.conn.Close()
		cs.mu.Lock()
		delete(cs.clients, c)
		cs.mu.Unlock()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		cs.broadcast <- msg
	}
}

func (c *Client) writeMessages() {
	for msg := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.conn.Close()
}

func chatHandler(cs *ChatServer, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Ошибка при подключении:", err)
		return
	}

	client := &Client{conn: conn, send: make(chan []byte)}
	cs.mu.Lock()
	cs.clients[client] = true
	cs.mu.Unlock()

	go client.readMessages(cs)
	go client.writeMessages()
}

func main() {
	chatServer := newChatServer()
	go chatServer.run()

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		chatHandler(chatServer, w, r)
	})

	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}

/*
Структура ChatServer: Хранит список клиентов и канал для рассылки сообщений.
Метод run: Запускает бесконечный цикл, который слушает сообщения в канале broadcast и рассылает их всем подключенным клиентам.
Методы readMessages и writeMessages: Обрабатывают чтение и запись сообщений для клиента. readMessages читает сообщения от клиента и отправляет их в broadcast, а writeMessages отправляет сообщения из канала send обратно клиенту.
Функция chatHandler: Обрабатывает подключение нового клиента, обновляет список клиентов и запускает горутины для чтения и записи сообщений.
*/
