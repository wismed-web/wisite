package main

import (
	"github.com/labstack/echo/v4"
	ad "github.com/wismed-web/wisite/server/api/admin"
	si "github.com/wismed-web/wisite/server/api/sign-in"
	so "github.com/wismed-web/wisite/server/api/sign-out"
	su "github.com/wismed-web/wisite/server/api/sign-up"
	"github.com/wismed-web/wisite/server/ws"
)

// path: handler
var mGET = map[string]echo.HandlerFunc{
	// web socket for message
	"/ws/msg": ws.WSMsg,

	// admin
	"/api/admin/users":       ad.ListUser,
	"/api/admin/onlineusers": ad.ListOnlineUser,

	// sign-in
	"/api/sign-in/signin": si.SignIn,

	// sign-out
	"/api/sign-out/signout": so.SignOut,
}

var mPOST = map[string]echo.HandlerFunc{
	// sign-up
	"/api/sign-up/new":          su.NewUser,
	"/api/sign-up/verify-email": su.VerifyEmail,
}

var mPUT = map[string]echo.HandlerFunc{
	"/api/test": nil,
}

var mDELETE = map[string]echo.HandlerFunc{
	"/api/test": nil,
}

var mPATCH = map[string]echo.HandlerFunc{
	"/api/test": nil,
}

// ---------------------------------------- //

func hookPathHandler(e *echo.Echo) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

	var mRegAPIs = map[string]map[string]echo.HandlerFunc{
		"GET":    mGET,
		"POST":   mPOST,
		"PUT":    mPUT,
		"DELETE": mDELETE,
		"PATCH":  mPATCH,
		// others...
	}

	type echoRoute func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	mRegMethod := map[string]echoRoute{
		"GET":    e.GET,
		"POST":   e.POST,
		"PUT":    e.PUT,
		"DELETE": e.DELETE,
		"PATCH":  e.PATCH,
		// others...
	}

	for _, m := range methods {
		mAPI := mRegAPIs[m]
		method := mRegMethod[m]
		for path, handler := range mAPI {
			if handler == nil {
				continue
			}
			method(path, handler)
		}
	}
}
