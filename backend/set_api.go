package main

import (
	"github.com/gin-gonic/gin"
	"toTheMoon/backend/constants"
	"toTheMoon/backend/handlers/auth"
)

func setApi(rApi *gin.RouterGroup) {

	// api/auth
	rAuth := rApi.Group(constants.Auth)
	{
		rAuth.POST(constants.Register, auth.Register)
	}


	//rExchange := rUpbit.Group(fmt.Sprintf("%s", constants.ApiSectionExchange))
	//{
	//	rExchange.GET(fmt.Sprintf("%s",constants.RouteAccounts))
	//}
	//
	//rQuotation := rUpbit.Group(fmt.Sprintf("%s", constants.ApiSectionQuotation))
	//{
	//	rQuotation.GET()
	//}
}
