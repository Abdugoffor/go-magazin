package category_handler

import (
	category_dto "category-crud/module/category/dto"
	category_service "category-crud/module/category/service"
	"category-crud/module/middleware"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type categoryHandler struct {
	db              *gorm.DB
	log             *log.Logger
	categoryService category_service.CategoryService
}

func NewCategoryHandler(group *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := &categoryHandler{
		db:              db,
		log:             log,
		categoryService: category_service.NewCategoryService(db),
	}

	group.GET("/category", handler.All)

	routeGroup := group.Group("/category", middleware.PermissionMiddleware())
	{
		// routeGroup.GET("", handler.All)
		routeGroup.GET("/:id", handler.Show)
		routeGroup.POST("", handler.Create)
		routeGroup.PUT("/:id", handler.Update)
		routeGroup.DELETE("/:id", handler.Delete)
	}
}

func (h *categoryHandler) All(c echo.Context) error {
	data, err := h.categoryService.All(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *categoryHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	data, err := h.categoryService.Show(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *categoryHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req category_dto.Create

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	data, err := h.categoryService.Create(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, data)
}

func (h *categoryHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	var req category_dto.Update
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	data, err := h.categoryService.Update(ctx, uint(id), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *categoryHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID"})
	}

	if err := h.categoryService.Delete(ctx, uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "deleted"})
}
