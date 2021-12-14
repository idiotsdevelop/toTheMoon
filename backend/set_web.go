package main

import (
	"github.com/gin-gonic/gin"
	"toTheMoon/backend/handlers-web/bithumb"
	"toTheMoon/backend/handlers-web/socket"
	"toTheMoon/backend/handlers-web/test"
)

func setWeb(rWeb *gin.RouterGroup) {

	rTest := rWeb.Group("/test")
	{
		rTest.GET("", test.Test)
	}

	rSocket := rWeb.Group("/socket")
	{
		rSocket.GET("", socket.Call)
	}

	//rUpbit := rWeb.Group("/upbit")
	{

	}

	rBithumb := rWeb.Group("/bithumb")
	{
		rBithumb.POST("/candle_stick", bithumb.CandleStick)
	}

}
