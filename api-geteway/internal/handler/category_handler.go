package handler

import (
	"errors"
	"net/http"
	"strconv"

	"api-geteway/internal/models"
	"api-geteway/internal/service"
	"api-geteway/pkg/response"

	categorypb "github.com/khbdev/what-food-proto/proto/"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	svc *service.CategoryService
}

// =========================
// INIT
// =========================

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: s}
}

// =========================
// CREATE CATEGORY
// =========================

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req models.Category

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.CreateCategory(c.Request.Context(), &categorypb.CreateCategoryRequest{
		Name: req.Name,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// GET BY ID
// =========================

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {

	idStr := c.Param("id")
	if idStr == "" {
		response.Fail(c, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	res, err := h.svc.GetCategoryByID(c.Request.Context(), &categorypb.GetByIDRequest{
		Id: id,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// GET ALL
// =========================

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {

	res, err := h.svc.GetAllCategories(c.Request.Context(), &categorypb.GetAllRequest{})
	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// UPDATE CATEGORY
// =========================

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {

	idStr := c.Param("id")
	if idStr == "" {
		response.Fail(c, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var req models.Category

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.UpdateCategory(c.Request.Context(), &categorypb.UpdateCategoryRequest{
		Id:   id,
		Name: req.Name,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// DELETE CATEGORY
// =========================

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {

	idStr := c.Param("id")
	if idStr == "" {
		response.Fail(c, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	res, err := h.svc.DeleteCategory(c.Request.Context(), &categorypb.DeleteCategoryRequest{
		Id: id,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// GET ALL WITH USER PRODUCTS
// =========================

func (h *CategoryHandler) GetAllWithUserProducts(c *gin.Context) {

	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		response.Fail(c, http.StatusBadRequest, errors.New("user_id is required"))
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid user_id"))
		return
	}

	res, err := h.svc.GetAllWithUserProducts(c.Request.Context(), &categorypb.GetAllWithUserProductsRequest{
		UserId: userID,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}