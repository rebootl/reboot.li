package config

import "os"

var CookieName = "rebootli-session-123"

func init() {
	if os.Getenv("REBOOTLI_COOKIE_NAME") != "" {
		CookieName = os.Getenv("REBOOTLI_COOKIE_NAME")
	}
}
