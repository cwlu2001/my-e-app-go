package myip

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cwlu2001/my-e-app-go/errorHandler"
)

func extractLine(s string, matches ...string) string {
	var r strings.Builder
	lines := strings.Split(s, "\n")
	for _, match := range matches {
		for _, line := range lines {
			if strings.HasPrefix(line, match) {
				r.WriteString(line)
				r.WriteString(", ")
				break
			}
		}
	}
	return strings.TrimSuffix(r.String(), ",")
}

func Myip() {
	resp, err := http.Get("https://www.cloudflare.com/cdn-cgi/trace")
	errorHandler.Handler(err)

	body, err := io.ReadAll(resp.Body)
	errorHandler.Handler(err)

	msg := extractLine(string(body), "ip", "loc", "ts")
	fmt.Println(msg)
}
