package handler

import (
	"errors"
	"net/http"
	"strconv"

	"api-geteway/internal/models"
	"api-geteway/internal/service"
	"api-geteway/pkg/response"

	ingredientpb "github.com/khbdev/what-food-proto/proto/incrideats"

	"github.com/gin-gonic/gin"
)

type IngredientHandler struct {
	svc *service.ProductService
}

// =========================
// INIT
// =========================

func NewIngredientHandler(s *service.ProductService) *IngredientHandler {
	return &IngredientHandler{svc: s}
}

// =========================
// CREATE
// =========================

func (h *IngredientHandler) CreateIngredient(c *gin.Context) {

	var req models.Ingredient

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.CreateIngredient(c.Request.Context(), &ingredientpb.CreateIngredientRequest{
		UserId:     req.UserID,
		Name:       req.Name,
		Quantity:   req.Quantity,
		CategoryId: req.CategoryID,
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

func (h *IngredientHandler) GetIngredientByID(c *gin.Context) {

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

	res, err := h.svc.GetIngredientByID(c.Request.Context(), &ingredientpb.GetIngredientByIDRequest{
		Id:     id,
		UserId: userID,
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

func (h *IngredientHandler) GetAllIngredients(c *gin.Context) {

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

	res, err := h.svc.GetAllIngredients(c.Request.Context(), &ingredientpb.GetAllIngredientsRequest{
		UserId: userID,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// UPDATE
// =========================

func (h *IngredientHandler) UpdateIngredient(c *gin.Context) {

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

	var req models.Ingredient

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.UpdateIngredient(c.Request.Context(), &ingredientpb.UpdateIngredientRequest{
		Id:         id,
		UserId:     req.UserID,
		Name:       req.Name,
		Quantity:   req.Quantity,
		CategoryId: req.CategoryID,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// DELETE
// =========================

func (h *IngredientHandler) DeleteIngredient(c *gin.Context) {

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

	res, err := h.svc.DeleteIngredient(c.Request.Context(), &ingredientpb.DeleteIngredientRequest{
		Id:     id,
		UserId: userID,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}