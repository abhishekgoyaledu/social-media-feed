package main

import (
	"github.com/social-media/user-service/api"
	"github.com/social-media/user-service/conf"
)

func main() {
	api.BuildRouter(conf.Config.InternalHostAndPort)
}
