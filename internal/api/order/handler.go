package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	appOrder "github.com/peterprospl12/logi-flow/internal/app/order"
)

type Handler struct {
	svc appOrder.Service
}

func NewHandler(svc appOrder.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	grp := r.Group("/api/v1/orders")
	{
		grp.POST("", h.create)
		grp.GET("", h.list)
	}
}

func (h *Handler) create(c *gin.Context) {
	var req CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := req.ToCommand()
	o, err := h.svc.CreateOrder(context.Background(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	location := fmt.Sprintf("/api/v1/orders/%s", o.ID.String())
	c.Header("Location", location)
	c.JSON(http.StatusCreated, ToGetOrderDTO(o))
}

func (h *Handler) list(c *gin.Context) {
	orders, err := h.svc.GetOrders(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	out := make([]GetOrderDTO, 0, len(orders))
	for _, o := range orders {
		out = append(out, ToGetOrderDTO(o))
	}

	c.JSON(http.StatusOK, out)
}
