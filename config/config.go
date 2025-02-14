package config

import "os"

var CookieName = "rebootli-session-123"
var Port = "8080"

func init() {
	if os.Getenv("REBOOTLI_COOKIE_NAME") != "" {
		CookieName = os.Getenv("REBOOTLI_COOKIE_NAME")
	}
	if os.Getenv("REBOOTLI_PORT") != "" {
		Port = os.Getenv("REBOOTLI_PORT")
	}
}
