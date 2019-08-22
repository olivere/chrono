package server

import (
	"net/http"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/olivere/httputil"
)

// homeHandler renders the home page.
func homeHandler(logger log.Logger) http.HandlerFunc {
	hostname, _ := os.Hostname()
	type response struct {
		Hostname string    `json:"hostname"`
		Env      []string  `json:"env"`
		Time     time.Time `json:"time"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		defer httputil.RecoverJSON(w, r)

		resp := response{
			Hostname: hostname,
			Env:      os.Environ(),
			Time:     time.Now().UTC(),
		}
		httputil.WriteJSON(w, resp)
	}
}
