package config

import (
	"errors"
	"os"
	"strings"
)

const (
	ModeDev  int = 0
	ModeProd int = 1
)

var (
	CookieName = "rebootli-session-123"
	Port       = "8080"
	// production is the default because it is safer
	// e.g. error pages are not shown in the browser
	Mode = ModeProd
)

func init() {
	if value, ok := os.LookupEnv("REBOOTLI_COOKIE_NAME"); ok {
		CookieName = value
	}
	if value, ok := os.LookupEnv("REBOOTLI_PORT"); ok {
		Port = value
	}
	if value, ok := os.LookupEnv("REBOOTLI_MODE"); ok {
		mode := strings.ToUpper(value)
		switch mode {
		case "DEV":
			Mode = ModeDev
		case "PROD":
			Mode = ModeProd
		default:
			panic(errors.New("invalid mode: " + value))
		}
	}
}
