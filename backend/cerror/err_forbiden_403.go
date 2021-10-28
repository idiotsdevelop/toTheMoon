package cerror

import (
	"net/http"
)

const (
	ForbiddenErrTitle = "BOSS Admin"
	ForbiddenErrBody  = "You do not have authority to edit this page."
)

/*
Authorized (Bearer) token이 요청 메세지에 포함되었으나, 해당 token의 유저가 특정 resource(api)에 권한이 없는 경우.

The request contained valid data and was understood by the server,
but the server is refusing action.
This may be due to the user not having the necessary permissions for a resource or needing an account of some sort,
or attempting a prohibited action (e.g. creating a duplicate record where only one is allowed).
This code is also typically used if the request provided authentication via the WWW-Authenticate header field,
but the server did not accept that authentication. The request SHOULD NOT be repeated.
*/
func Forbidden() CustomError {
	return newCustomError(
		http.StatusForbidden,
		ForbiddenErrBody,
	)
}

func ForbiddenWithMsg(msg string) CustomError {
	return newCustomError(
		http.StatusForbidden,
		msg,
	)
}
