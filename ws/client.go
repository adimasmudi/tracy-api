package ws

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Conn           *websocket.Conn
	Message        chan *Message
	ID             string `json:"id" bson:"_id"`
	RoomID         string             `json:"roomId"`
	SenderEmail    string             `json:"senderEmail"`
	ReceriverEmail string             `json:"receiverEmail"`
}

type Message struct {
	Content        string `json:"content"`
	RoomID         string `json:"roomId"`
	SenderEmail    string `json:"senderEmail"`
	ReceriverEmail string `json:"receriverEmail"`
}

func (c *Client) writeMessage(){
	defer func(){
		c.Conn.Close()
	}()

	for {
		message, ok := <- c.Message

		if !ok{
			return 
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub){
	defer func(){
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()

		if err != nil{
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content : string(m),
			RoomID: c.RoomID,
			SenderEmail: c.SenderEmail,
			ReceriverEmail: c.ReceriverEmail,
		}

		hub.Broadcast <- msg
	}
}