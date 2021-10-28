package middlewares

import (
	"bytes"
	"fmt"
	"github.com/getsentry/sentry-go"
	"toTheMoon/backend/cerror"
	"toTheMoon/backend/config"
	"toTheMoon/backend/constants"

	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)

func WrapMiddleware(c *gin.Context) {
	// set header
	if !strings.Contains(c.Request.URL.Path, "swagger") && !strings.Contains(c.Request.URL.Path, "static") {
		c.Writer.Header().Set("Content-Type", "application/json")
	}
	c.Writer.Header().Set(constants.HeaderCacheControl, constants.CacheNoStore)

	//store request body to use when err occur
	bodyByte, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))
	//c.Set(constants.CtxRequestBody, bodyByte)
	defer func() {
		//reqBodyByte, _ := c.Get(constants.CtxRequestBody)
		//reqBody := reqBodyByte.([]byte)

		// error Handling
		if r := recover(); r != nil {
			var customError cerror.CustomError
			switch r.(type) {
			case error:
				customError = cerror.CustomError{
					StatusCode: http.StatusInternalServerError,
					Message:    http.StatusText(http.StatusInternalServerError),
				}
			case string:
				customError = cerror.CustomError{
					StatusCode: http.StatusInternalServerError,
					Message:    http.StatusText(http.StatusInternalServerError),
				}
			case cerror.CustomError:
				customError = r.(cerror.CustomError)
			default:
				customError = cerror.CustomError{
					StatusCode: http.StatusInternalServerError,
					Message:    http.StatusText(http.StatusInternalServerError),
				}
			}
			c.JSON(customError.StatusCode, customError)

			sentry.CaptureException(&customError)
			//if !config.GetBool(config.IsProd) {
			errorStack := string(debug.Stack())
			if config.IsLocal() {
				log.Println(errorStack)
				log.Println(customError)
			} else {
				if c.ContentType() == "multipart/form-data" {
					log.Println(errorStack)
					log.Println(customError)
				} else {
					errorStack = fmt.Sprintf("Request Body: %s //////// %s", string(bodyByte), errorStack)
					log.Println(strings.ReplaceAll(errorStack, "\n", "\t"))
					log.Println(customError)
				}
			}
			c.Abort()
		}
	}()

	c.Next()
}
