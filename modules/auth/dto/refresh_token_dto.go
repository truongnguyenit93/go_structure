package dto

const (
	MESSAGE_SUCCESS_REFRESH_TOKEN        = "Successfully refreshed token"
	MESSAGE_FAILED_REFRESH_TOKEN         = "Failed to refresh token"
	MESSAGE_FAILED_INVALID_REFRESH_TOKEN = "Invalid refresh token"
	MESSAGE_FAILED_EXPIRED_REFRESH_TOKEN = "Refresh token has expired"
	MESSAGE_FAILED_PROCESS_REQUEST		 = "Failed to process request"
	MESSAGE_FAILED_TOKEN_NOT_FOUND		 = "Token not found"
	MESSAGE_FAILED_TOKEN_NOT_VALID		 = "Token not valid"
	MESSAGE_FAILED_DENIED_ACCESS		 = "Denied access"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}