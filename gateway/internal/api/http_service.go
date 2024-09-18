package api

import (
	"awesome-chat/gateway/internal/config"
	"awesome-chat/gateway/internal/model"
	"awesome-chat/gateway/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type HttpServer struct {
	e              *echo.Echo
	cfg            config.HttpConfig
	messageService service.MessageService
}

func NewHttpServer(cfg config.HttpConfig, messageService service.MessageService) *HttpServer {
	return &HttpServer{
		cfg:            cfg,
		messageService: messageService,
	}
}

func (server *HttpServer) Run() {
	e := echo.New()
	group := e.Group("/message")
	group.POST("", server.Create)
	err := e.Start(server.cfg.GetAddress())
	if err != nil {
		log.Fatalf("cannot start http server: %v", err)
	}
}

func (server *HttpServer) Create(c echo.Context) error {
	text := c.QueryParam("msg_text")
	m := model.Message{Text: text}
	err := server.messageService.Create(context.Background(), m)
	if err != nil {
		c.String(http.StatusInternalServerError, "some shit happened")
	}
	c.String(http.StatusOK, "created")
	return nil
}

// getRoutes for debug
func getRoutes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", " ")
	if err != nil {
		log.Fatalf("cannot get routes %v", err)
		return
	}
	fmt.Printf("available routes: %v", string(data))
}
