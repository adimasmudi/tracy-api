package ws

import (
	"net/http"
	"tracy-api/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
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

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request)bool{
		return true
	},
}

func (h *Handler) JoinRoom(c *fiber.Ctx) error{
	var w http.ResponseWriter
	var r *http.Request


	conn, err := upgrader.Upgrade(w, r, nil)
	
	if err != nil{
		response := helper.APIResponse("Failed to join chat", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	roomId := c.Params("roomId")
	senderEmail := c.Query("senderEmail")
	receiverEmail := c.Query("receiverEmail")

	cl := &Client{
		Conn : conn,
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
	h.hub.Register <- cl

	// broadcast that message
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)

	return nil

}