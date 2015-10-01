package controllers

import (
	"github.com/Rompei/zepher/db"
	"github.com/Rompei/zepher/models"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
	"net/http"
)

func BansakuIndex(c *echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// BansakuClient is client of bansaku
type BansakuClient struct {
	ID             int
	ws             *websocket.Conn
	removeClientCh chan *BansakuClient
	bansakuCh      chan string
}

// NewClient returns client
func NewClient(ws *websocket.Conn, remove chan *BansakuClient, bansaku chan string) *BansakuClient {
	return &BansakuClient{
		ws:             ws,
		removeClientCh: remove,
		bansakuCh:      bansaku,
	}
}

// Start is called websocket is opened.
func (client *BansakuClient) Start() {
	for {
		var bansaku string
		err := websocket.Message.Receive(client.ws, &bansaku)
		if err != nil {
			client.removeClientCh <- client
			return
		}
		client.bansakuCh <- bansaku
	}
}

// Send sends Bansaku count.
func (client *BansakuClient) Send(bansaku *models.Bansaku) {
	err := websocket.JSON.Send(client.ws, bansaku)
	if err != nil {
		panic(err)
	}
}

// Close close websocket
func (client *BansakuClient) Close() {
	client.ws.Close()
}

// BansakuServer is server of websocket.
type BansakuServer struct {
	clientCount    int
	clients        map[int]*BansakuClient
	addClientCh    chan *BansakuClient
	removeClientCh chan *BansakuClient
	bansakuCh      chan string
}

// NewServer returns server
func NewServer() *BansakuServer {
	return &BansakuServer{
		clientCount:    0,
		clients:        map[int]*BansakuClient{},
		addClientCh:    make(chan *BansakuClient),
		removeClientCh: make(chan *BansakuClient),
		bansakuCh:      make(chan string),
	}
}

func (server *BansakuServer) addClient(client *BansakuClient) {
	server.clientCount++
	client.ID = server.clientCount
	server.clients[client.ID] = client
}

func (server *BansakuServer) removeClient(client *BansakuClient) {
	delete(server.clients, client.ID)
}

func (server *BansakuServer) sendCount(bansaku *models.Bansaku) {
	for _, client := range server.clients {
		c := client
		go func() {
			c.Send(bansaku)
		}()
	}
}

// Start opens websocket
func (server *BansakuServer) Start() {
	for {
		select {
		case client := <-server.addClientCh:
			server.addClient(client)
		case client := <-server.removeClientCh:
			server.removeClient(client)
		case <-server.bansakuCh:
			c := db.GetRedis()
			count, err := redis.Int64(c.Do("get", "count"))
			if err != nil {
				panic(err)
			}
			count++
			c.Do("set", "count", count)
			bansakuCount := models.Bansaku{
				Count: count,
			}
			server.sendCount(&bansakuCount)
		}
	}
}
