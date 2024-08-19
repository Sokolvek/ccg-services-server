package feddis

import "sync"

type client interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
}

type Client struct {
	Map sync.Map
}

var (
	once     sync.Once
	instance *Client
)

func InitDB() {
	once.Do(func() {
		instance = NewClient()
	})
}

func GetClient() *Client {
	return instance
}

func NewClient() *Client {
	client := Client{sync.Map{}}

	return &client
}

func (c *Client) Set(key string, value interface{}) {
	c.Map.Store(key, value)
}

func (c *Client) Get(key string) (interface{}, error) {
	value, ok := c.Map.Load(key)

	if !ok {
		return nil, nil
	}

	return value, nil
}

func (c *Client) Has(key string) bool {
	_, ok := c.Map.Load(key)

	return ok
}
