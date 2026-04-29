package handler

import (
	"errors"
	"net/http"
	"strconv"

	"api-geteway/internal/service"
	"api-geteway/pkg/response"

	ingredientpb "github.com/khbdev/what-food-proto/proto/incrideats"

	"github.com/gin-gonic/gin"
)

type IngredientHandler struct {
	svc *service.ProductService
}

func NewIngredientHandler(s *service.ProductService) *IngredientHandler {
	return &IngredientHandler{svc: s}
}

// tokendan user_id olish uchun helper
func getUserID(c *gin.Context) (int64, error) {
	val, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("unauthorized")
	}
	uid, ok := val.(uint64)
	if !ok || uid == 0 {
		return 0, errors.New("invalid user_id")
	}
	return int64(uid), nil
}

// =========================
// CREATE
// =========================

func (h *IngredientHandler) CreateIngredient(c *gin.Context) {

	userID, err := getUserID(c)
	if err != nil {
		response.Fail(c, http.StatusUnauthorized, err)
		return
	}

	var req struct {
		Name       string  `json:"name"`
		Quantity   float64 `json:"quantity"`
		CategoryID int64   `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.CreateIngredient(c.Request.Context(), &ingredientpb.CreateIngredientRequest{
		UserId:     userID,
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

	userID, err := getUserID(c)
	if err != nil {
		response.Fail(c, http.StatusUnauthorized, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
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

	userID, err := getUserID(c)
	if err != nil {
		response.Fail(c, http.StatusUnauthorized, err)
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

	userID, err := getUserID(c)
	if err != nil {
		response.Fail(c, http.StatusUnauthorized, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var req struct {
		Name       string  `json:"name"`
		Quantity   float64 `json:"quantity"`
		CategoryID int64   `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.UpdateIngredient(c.Request.Context(), &ingredientpb.UpdateIngredientRequest{
		Id:         id,
		UserId:     userID,
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

	userID, err := getUserID(c)
	if err != nil {
		response.Fail(c, http.StatusUnauthorized, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
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