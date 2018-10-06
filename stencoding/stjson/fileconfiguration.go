package stjson

import (
	"encoding/json"
	"fmt"
	"os"
)

type Client struct {
	consulIP   string
	connString string
}

func (c *Client) String() string {
	return fmt.Sprintf("ConsulIP: %s, Connection String: %s", c.consulIP, c.connString)
}

var defaultClient = Client{
	consulIP:   "localhost:9000",
	connString: "postgres:localhost:5432",
}

type ConfigFunc func(opt *Client)

func FromFile(path string) ConfigFunc {
	return func(opt *Client) {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		decoder := json.NewDecoder(f)

		fop := struct {
			ConsulIP string `json:"consul_ip"`
		}{}
		err = decoder.Decode(&fop)

		if err != nil {
			panic(err)
		}

		opt.consulIP = fop.ConsulIP
	}
}

func FromEnv() ConfigFunc {
	return func(opt *Client) {
		connsStr, exist := os.LookupEnv("CONN_DB")
		if exist {
			opt.connString = connsStr
		}
	}
}

func NewClient(opts ...ConfigFunc) *Client {
	client := defaultClient
	for _, val := range opts {
		val(&client)
	}

	return &client
}

func FileConfigulation() {
	client := NewClient(FromFile("config.json"), FromEnv())
	fmt.Println(client.String())
}
