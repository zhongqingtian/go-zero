package svc

import (
	"shorturl/api/internal/config"
	"shorturl/rpc/transform/transformer"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
//		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
	}
}
