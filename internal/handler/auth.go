package handler

import (
	"Afisha/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(ctx *gin.Context) {
	var input domain.UserToLogin
	if err := ctx.ShouldBind(&input); err != nil {

		return
	}
}

func (h *Handler) SignOut(ctx *gin.Context) {

}

func (h *Handler) SignUp(ctx *gin.Context) {

}
