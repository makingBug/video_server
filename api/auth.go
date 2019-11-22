package main

import (
	"net/http"
	"video_server/api/defs"
	"video_server/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELE_UNAME = "X-User-Name"

//检查用户的session是否是合法的
func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0{
		return false
	}

	uname,ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FIELE_UNAME,uname)
	return true
}

func ValidateUser(w http.ResponseWriter,r *http.Request) bool{
	uname := r.Header.Get(HEADER_FIELE_UNAME)
	if len(uname) == 0{
		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return false
	}
	return true
}