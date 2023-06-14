package order

import (
	"github.com/MelomiLight/go-grpc-api-gateway/pkg/auth"
	"github.com/MelomiLight/go-grpc-api-gateway/pkg/config"
	"github.com/MelomiLight/go-grpc-api-gateway/pkg/order/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	route := r.Group("/order")
	route.Use(a.AuthRequired)
	route.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
