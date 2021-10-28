package cerror

import "net/http"

const (
	FCMErrTitle     = "푸쉬 알람 전송에 문제가 발생했습니다."
	UnknownErrTitle = "일시적인 문제가 발생하였습니다. 잠시 후 다시 시도해주세요."
)

/*
A generic error message,
given when an unexpected condition was encountered and no more specific message is suitable.
*/
func DBErr(err error) CustomError {
	return newCustomError(
		http.StatusInternalServerError,
		UnknownErrTitle,
		//err.Error(),
		// http.StatusText(http.StatusInternalServerError),
	)
}

func Marshal(err error) CustomError {
	return newCustomError(
		http.StatusInternalServerError,
		UnknownErrTitle,
		//err.Error(),
		// http.StatusText(http.StatusInternalServerError),
	)
}

func Unknown(err error) CustomError {
	return newCustomError(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
	)
}

func S3Error(err error) CustomError {
	return newCustomError(
		http.StatusInternalServerError,
		UnknownErrTitle,
		//http.StatusText(http.StatusInternalServerError),
	)
}

func SecretsManagerError(err error) CustomError {
	return newCustomError(
		http.StatusInternalServerError,
		UnknownErrTitle,
		//http.StatusText(http.StatusInternalServerError),
	)
}

func UnMarshallError(err error) CustomError {
	return newCustomError(
		http.StatusInternalServerError,
		UnknownErrTitle,
		//http.StatusText(http.StatusInternalServerError),
	)
}

func SendSMSError() CustomError {
	return newCustomError(
		http.StatusInternalServerError,
		UnknownErrTitle,
		//http.StatusText(http.StatusInternalServerError),
	)
}

func FCMError(err error) CustomError {
	return newCustomError(
		http.StatusInternalServerError,
		FCMErrTitle,
		//http.StatusText(http.StatusInternalServerError),
	)
}
