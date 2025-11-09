package cmd

import (
	cfg "poteto-backyard-be/services/core/infra/config"

	"github.com/kataras/i18n"
)

type Config struct {
	Http *cfg.Client `config:"http"`
	// microservices
	User *cfg.Client `config:"user"`

	// infra
	I18n *i18n.Config `config:"i18n"`

	// MessageQueue *pubsub.Config `config:"message_queue"`
	// Websocket    *ws.WSConfig   `config:"websocket"`
	// Proxy        *proxy.Config  `config:"proxy"`
}
