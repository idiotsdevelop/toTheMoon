package cerror

import "fmt"

/*
https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
*/
type CustomError struct {
	// timestamp
	// id
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func newCustomError(statusCode int, errorMessage string) CustomError {
	return CustomError{
		StatusCode: statusCode,
		Message:    errorMessage,
	}
}

func (error *CustomError) Error() string {
	return fmt.Sprintf("[%d]%s", error.StatusCode, error.Message)
}

type CustomError400 struct {
	StatusCode int    `json:"status_code" example:"400"`
	Message    string `json:"message" example:"Request body is malformed!"`
}

type CustomError401 struct {
	StatusCode int    `json:"status_code" example:"401"`
	Message    string `json:"message" example:"SignIn required or Refresh token required or Bad SingIn Credential!"`
}

type CustomError403 struct {
	StatusCode int    `json:"status_code" example:"403"`
	Message    string `json:"message" example:"You do not have authority to call this api!"`
}

type CustomError404 struct {
	StatusCode int    `json:"status_code" example:"404"`
	Message    string `json:"message" example:"resource not found"`
}

type CustomError409 struct {
	StatusCode int    `json:"status_code" example:"409"`
	Message    string `json:"message" example:"Unique constraints violation!"`
}

type CustomError500 struct {
	StatusCode int    `json:"status_code" example:"500"`
	Message    string `json:"message" example:"Unexpected internal server error!"`
}
