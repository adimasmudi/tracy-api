package ws

import (
	"net/http"
	"tracy-api/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
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
		Clients : make(map[string]*Client),
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

type RoomRes struct{
	ID string `json:"id"`
}

func (h *Handler) GetRooms(c *fiber.Ctx)error{
	rooms := make([]RoomRes, 0)

	for _, r := range(h.hub.Rooms){
		rooms = append(rooms, RoomRes{
			ID : r.ID,
		})
	}

	response := helper.APIResponse("success get rooms", http.StatusOK, "success", rooms)
	c.Status(http.StatusOK).JSON(response)
	return nil
}

type ClientRes struct{
	ID string `json:"id" bson:"_id"`
	SenderEmail string `json:"senderEmail"`
	ReceiverEmail string `json:"receiverEmail"`
}

func (h *Handler) GetClients(c *fiber.Ctx) error{
	var clients []ClientRes

	roomId := c.Params("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok{
		clients = make([]ClientRes, 0)
		response := helper.APIResponse("success get clients in room here", http.StatusOK, "success", clients)
		c.Status(http.StatusOK).JSON(response)
		return nil
	}

	for _, c := range(h.hub.Rooms[roomId].Clients){
		clients = append(clients, ClientRes{
			ID : c.ID,
			SenderEmail : c.SenderEmail,
			ReceiverEmail : c.ReceriverEmail,
		})
	}

	response := helper.APIResponse("success get clients in room", http.StatusOK, "success", clients)
	c.Status(http.StatusOK).JSON(response)
	return nil
}