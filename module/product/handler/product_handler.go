package product_handler

import (
	"category-crud/module/middleware"
	product_dto "category-crud/module/product/dto"
	product_service "category-crud/module/product/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type productHandler struct {
	db             *gorm.DB
	log            *log.Logger
	productService product_service.ProductService
}

func NewProductHandler(group *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := &productHandler{
		db:             db,
		log:            log,
		productService: product_service.NewProductService(db),
	}

	productGroup := group.Group("/product", middleware.PermissionMiddleware())
	{
		productGroup.GET("", handler.All)
		productGroup.GET("/:id", handler.Show)
		productGroup.POST("", handler.Create)
		productGroup.PUT("/:id", handler.Update)
		productGroup.DELETE("/:id", handler.Delete)
	}
}

func (h *productHandler) All(c echo.Context) error {

	data, err := h.productService.All(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *productHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	data, err := h.productService.Show(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *productHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req product_dto.Create

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	data, err := h.productService.Create(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, data)
}

func (h *productHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	var req product_dto.Update
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	data, err := h.productService.Update(ctx, uint(id), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *productHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	if err := h.productService.Delete(ctx, uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "deleted"})
}
