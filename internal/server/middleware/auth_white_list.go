package middleware

var authWhiteList = map[string][]string{
	"GET": {},
	"POST": {
		"/authn/login",
	},
	"PATCH":  {},
	"UPDATE": {},
	"DELETE": {},
}
