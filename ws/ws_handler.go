package ws

import (
	"net/http"
	"tracy-api/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateReqRoom struct{
	ID string `json:"id"`
}

func (h *Handler) CreateRoom(c *fiber.Ctx) error {
	var req CreateReqRoom

	if err := c.BodyParser(&req);  err != nil{
		response := helper.APIResponse("Failed to request chat", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	h.hub.Rooms[req.ID] = &Room{
		ID : req.ID,
		Clients : make(map[primitive.ObjectID]*Client),
	}

	response := helper.APIResponse("success to request chat", http.StatusOK, "success", req)
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func JoinRoom(hub *Hub) fiber.Handler{
	return websocket.New(func(c *websocket.Conn){
		
		roomId := c.Params("roomId")
		senderEmail := c.Query("senderEmail")
		receiverEmail := c.Query("receiverEmail")

		cl := &Client{
			Conn : c,
			Message : make(chan *Message, 10),
			RoomID : roomId,
			SenderEmail: senderEmail,
			ReceriverEmail: receiverEmail,
		}

		m := &Message{
			Content : "A new Request Chat",
			RoomID: roomId,
			SenderEmail: senderEmail,
			ReceriverEmail: receiverEmail,
		}

		// register a new client through register channel
		hub.Register <- cl


		// broadcast that message
		hub.Broadcast <- m

		go cl.writeMessage()
		cl.readMessage(hub)

	})


}