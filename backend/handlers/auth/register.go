package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"toTheMoon/backend/cerror"
	"toTheMoon/backend/handlers/upbit"
)

type ReqRegister struct {
	UpbitAccessKey string `json:"upbit_access_key" binding:"required"`
	UpbitSecretKey string `json:"upbit_secret_key" binding:"required"`
}

type RespRegister struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
}


// @tags auth
// @Summary 회원가입
// @Description 회원가입
// @Accept json
// @Produce json
// @Param body body ReqRegister true "접근 키, 시크릿 키"
// @Success 200 {object} RespRegister
// @Failure 400 {object} RespRegister
// @Failure 401 {object} RespRegister
// @Failure 500 {object} RespRegister
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	var body ReqRegister
	if err := c.ShouldBind(&body); err != nil {
		panic(cerror.BadRequest())
	}

	var code int
	var message string

	u := upbit.NewUpbit(body.UpbitAccessKey,body.UpbitSecretKey)
	_, _, e := u.GetAccounts()

	if e != nil {
		code = http.StatusUnauthorized
		message = fmt.Sprintf("%s",e)
	} else {
		code = http.StatusOK
		message = fmt.Sprintf("인증이 완료 되었습니다.")

		// db transaction
		//user := models.User{
		//
		//}


	}
	c.JSON(code, message)
}