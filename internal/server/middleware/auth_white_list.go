package middleware

var authWhiteList = map[string][]string{
	"GET":  {},
	"POST": {
		// "/identity/user",
	},
	"PATCH":  {},
	"UPDATE": {},
	"DELETE": {},
}
