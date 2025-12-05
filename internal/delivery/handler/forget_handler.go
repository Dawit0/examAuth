package handler

import (
	"net/http"

	"github.com/Dawit0/examAuth/internal/delivery/dto"
	"github.com/Dawit0/examAuth/internal/service"
	"github.com/gin-gonic/gin"
)

type ForgetPasswordHandler struct {
	resetUserService *service.ResetUserService
}

func NewForgetPasswordHandler(resetUserService *service.ResetUserService) *ForgetPasswordHandler {
	return &ForgetPasswordHandler{resetUserService: resetUserService}
}

func (h *ForgetPasswordHandler) RequestResetPasswordEmail(c *gin.Context) {
	var req dto.RequestResetDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.resetUserService.RequestResetPasswordEmail(req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password reset email sent"})
}

func (h *ForgetPasswordHandler) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.resetUserService.ResetPassword(req.Email, req.OTP, req.NewPassword); err != nil {
		switch err.Error() {
		case "password reset expired":
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "OTP expired",
			})
		case "invalid credentials", "record not found":
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid OTP",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to reset password",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password reset successfully"})
}
