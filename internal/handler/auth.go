package handler

import (
	"Afisha/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) SignIn(ctx *gin.Context) {
	var input domain.UserToLogin
	if err := ctx.ShouldBind(&input); err != nil {
		logrus.Error(err)
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.IUserService.Verify(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ip := ctx.ClientIP()
	tokens, err := h.service.ITokenService.GenerateTokens(ip)
	if err != nil {
		logrus.Error(err)
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.ITokenService.AddToken(user.Id, ip, tokens["refreshToken"])
	if err != nil {
		logrus.Error(err)
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"User":         user,
		"AccessToken":  tokens["accessToken"],
		"RefreshToken": tokens["refreshToken"],
	})
}

func (h *Handler) SignOut(ctx *gin.Context) {

}

func (h *Handler) SignUp(ctx *gin.Context) {
	var input domain.UserRegister
	if err := ctx.ShouldBind(&input); err != nil {
		logrus.Error(err)
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.IUserService.Create(input)
	if err != nil {
		logrus.Error(err)
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
