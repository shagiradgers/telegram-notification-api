package server

import (
	desc "telegram-notification-api/api"
	"telegram-notification-api/internal/clients"
	"telegram-notification-api/internal/dao"
)

type server struct {
	dao     dao.DAO
	clients clients.Clients

	desc.UnimplementedTelegramNotificationServiceServer
}

func (s *server) mustEmbedUnimplementedTelegramNotificationServiceServer() {}

func NewServer(dao dao.DAO, clients clients.Clients) desc.TelegramNotificationServiceServer {
	s := &server{
		dao:     dao,
		clients: clients,
	}
	return s
}
