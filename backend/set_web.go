package main

import (
	"github.com/gin-gonic/gin"
	"toTheMoon/backend/handlers-web/test"
)

func setWeb(rWeb *gin.RouterGroup) {

	rTest := rWeb.Group("/test")
	{
		rTest.GET("", test.Test)
	}

}
