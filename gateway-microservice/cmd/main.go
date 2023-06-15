package main

import (
	"log"

	"github.com/MelomiLight/gateway-microservice/pkg/auth"
	"github.com/MelomiLight/gateway-microservice/pkg/config"
	"github.com/MelomiLight/gateway-microservice/pkg/order"
	"github.com/MelomiLight/gateway-microservice/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
