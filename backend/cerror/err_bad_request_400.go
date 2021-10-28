package cerror

import (
	"net/http"
)

const (

	// auth
	ErrPasswordNotMatched         = "휴대폰 번호와 비밀번호를 다시 한 번 확인해주세요."
	ErrRefreshTokenInvalid        = "인증정보가 만료되었습니다. 다시 로그인을 시도해주세요."
	ErrVerificationCodeNotMatched = "인증번호가 일치하지 않습니다."
	ErrNumPasswordFailExceed      = "비밀번호 최대 실패 횟수를 초과하였습니다."

	// unique
	ErrPhoneNumberUnique = "이미 등록된 전화번호입니다."
	ErrUserIdUnique      = "이미 등록된 아이디입니다."

	// not found
	RecordNotFound = "존재하지 않는 기록입니다."
)

/*
보내온 데이터가 잘못됐을 경우  or 잘못된 방식으로 URL을 호출한 경우
- frontend 로직의 문제
- 사용자가 잘못된 데이터로 api를 직접 호출한 경우

The server cannot or will not process the request due to an apparent client error
(e.g., malformed request syntax, size too large, invalid request message framing, or deceptive request routing)
*/

func BadRequest() CustomError {
	return newCustomError(
		http.StatusBadRequest,
		http.StatusText(http.StatusBadRequest))
}

func BadRequestWithMsg(message string) CustomError {
	return newCustomError(
		http.StatusBadRequest,
		message,
	)
}

func RecordNotFoundError() CustomError {
	return newCustomError(
		http.StatusBadRequest,
		RecordNotFound,
		//http.StatusText(http.StatusInternalServerError),
	)
}
