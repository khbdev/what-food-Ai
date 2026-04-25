package models

// =========================
// CREATE USER
// =========================

type CreateUserRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Age     int32  `json:"age"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Image   string `json:"image"`
}

// =========================
// GET USER BY ID
// =========================

type GetUserByIDRequest struct {
	ID uint64 `json:"id"`
}

// =========================
// GET USER BY PHONE
// =========================

type GetUserByPhoneRequest struct {
	Phone string `json:"phone"`
}

// =========================
// UPDATE USER
// =========================

type UpdateUserRequest struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Age     int32  `json:"age"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Image   string `json:"image"`
}

// =========================
// DELETE USER
// =========================

type DeleteUserRequest struct {
	ID uint64 `json:"id"`
}