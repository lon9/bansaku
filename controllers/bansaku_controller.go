package controllers

import (
	"github.com/Rompei/zepher-bansaku/db"
	"github.com/Rompei/zepher-bansaku/models"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
	"net/http"
)

// BansakuIndex is root of Bansaku button
func BansakuIndex(c *echo.Context) error {
	data := make(map[string]interface{})
	return c.Render(http.StatusOK, "bansaku", data)
}

// BansakuClient is client of bansaku
type BansakuClient struct {
	ID             int
	context        *echo.Context
	removeClientCh chan *BansakuClient
	bansakuCh      chan string
}

// NewBansakuClient returns client
func NewBansakuClient(c *echo.Context, remove chan *BansakuClient, bansaku chan string) *BansakuClient {
	return &BansakuClient{
		context:        c,
		removeClientCh: remove,
		bansakuCh:      bansaku,
	}
}

// Start is called websocket is opened.
func (client *BansakuClient) Start() {
	for {
		var bansaku string
		err := websocket.Message.Receive(client.context.Socket(), &bansaku)
		if err != nil {
			client.removeClientCh <- client
			return
		}
		client.bansakuCh <- bansaku
	}
}

// Send sends Bansaku count.
func (client *BansakuClient) Send(bansaku *models.Bansaku) {
	err := websocket.JSON.Send(client.context.Socket(), bansaku)
	if err != nil {
		panic(err)
	}
}

// Close close websocket
func (client *BansakuClient) Close() {
	client.context.Socket().Close()
}

// BansakuServer is server of websocket.
type BansakuServer struct {
	clientCount    int
	clients        map[int]*BansakuClient
	addClientCh    chan *BansakuClient
	removeClientCh chan *BansakuClient
	bansakuCh      chan string
}

// NewBansakuServer returns server
func NewBansakuServer() *BansakuServer {
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

	c := db.GetRedis()
	count, err := redis.Int64(c.Do("get", "count"))
	if err != nil {
		count = 0
	}
	client.Send(&models.Bansaku{
		Count: count,
	})
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
				count = 1
			} else {
				count++
			}
			c.Do("set", "count", count)
			bansakuCount := models.Bansaku{
				Count: count,
			}
			server.sendCount(&bansakuCount)
		}
	}
}

// BansakuSocketHandler is handler for treat socket
func (server *BansakuServer) BansakuSocketHandler() echo.HandlerFunc {
	return echo.HandlerFunc(func(c *echo.Context) (err error) {
		client := NewBansakuClient(c, server.removeClientCh, server.bansakuCh)
		server.addClientCh <- client
		client.Start()
		return
	})
}
