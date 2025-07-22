package order_handler

import (
	"category-crud/middleware"
	order_dto "category-crud/module/shop/order/dto"
	order_service "category-crud/module/shop/order/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type orderHandler struct {
	db           *gorm.DB
	log          *log.Logger
	orderService order_service.OrderService
}

func NewOrderHandler(group *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := &orderHandler{
		db:           db,
		log:          log,
		orderService: order_service.NewOrderService(db),
	}

	routeGroup := group.Group("/order", middleware.PermissionMiddleware())
	{
		routeGroup.GET("", handler.All)
		routeGroup.GET("/:id", handler.Show)
		routeGroup.POST("", handler.Create)
		routeGroup.PUT("/:id", handler.Update)
		routeGroup.DELETE("/:id", handler.Delete)
	}
}

func (h *orderHandler) All(c echo.Context) error {
	data, err := h.orderService.All(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *orderHandler) Show(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	data, err := h.orderService.Show(c, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *orderHandler) Create(c echo.Context) error {

	var req order_dto.OrderCreate

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	data, err := h.orderService.Create(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, data)
}

func (h *orderHandler) Update(c echo.Context) error {
	return nil
}

func (h *orderHandler) Delete(c echo.Context) error {
	return nil
}
