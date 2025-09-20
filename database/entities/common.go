package entities

import "time"

// Timestamp is a struct that contains common fields for tracking creation and update times.
type Timestamp struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}


type Pagination struct {
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int64 `json:"total_pages"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type Authorization struct {
	Token string `json:"token"`
	Role  string `json:"role" binding:"required,oneof=admin author user"`
}