package main

import (
	"github.com/gin-gonic/gin"
	apiOrder "github.com/peterprospl12/logi-flow/internal/api/order"
	appOrder "github.com/peterprospl12/logi-flow/internal/app/order"
	"github.com/peterprospl12/logi-flow/internal/infra/db/memory"
)

func main() {
	store := memory.NewStore()
	orderRepo := memory.NewOrderRepository(store)

	orderService := appOrder.NewService(orderRepo)
	orderHandler := apiOrder.NewHandler(orderService)

	r := gin.Default()
	orderHandler.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
