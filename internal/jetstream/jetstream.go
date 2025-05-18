package jetstream

import (
	"encoding/json"
	"log"
	"net/url"
	"time"

	"github.com/agentio/atiquette/api/xyz/statusphere"
	"github.com/agentio/statusphere/internal/storage"
	"github.com/gorilla/websocket"
)

const addr = "jetstream2.us-west.bsky.network"

func Listen() {

	u := url.URL{Scheme: "wss", Host: addr, Path: "/subscribe", RawQuery: "wantedCollections=xyz.statusphere.status"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		var m Message
		json.Unmarshal(message, &m)
		if m.Kind == "commit" {
			log.Printf("%+v\n", m)
			storage.SaveStatus(&storage.Status{
				Uri:       "at://" + m.Did + "/" + m.Commit.Collection + "/" + m.Commit.RKey,
				AuthorDid: m.Did,
				Status:    m.Commit.Record.Status,
				CreatedAt: m.Commit.Record.CreatedAt,
				UpdatedAt: time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
			})
		}
	}
}

type Message struct {
	Did    string `json:"did"`
	Time   int64  `json:"time_us"`
	Kind   string `json:"kind"`
	Commit struct {
		Rev        string             `json:"rev"`
		Operation  string             `json:"operation"`
		Collection string             `json:"collection"`
		RKey       string             `json:"rkey"`
		Record     statusphere.Status `json:"record"`
		Cid        string             `json:"cid"`
	} `json:"commit"`
}
