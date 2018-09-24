package stubs

import (
	"net/http"

	"time"

	"github.com/alanfoster/configurable-proxy/config"
)

type Handler struct {
	ConfigManager *config.Manager
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handled := false

	time.Sleep(1*time.Second + 500*time.Millisecond)

	for _, stub := range *h.ConfigManager.Stubs {
		if !handled {
			handled = true
			response := stub.Response

			w.WriteHeader(response.Status)
			for header, value := range response.Headers {
				w.Header().Add(header, value)
			}
			w.Write([]byte(response.Body))
		}
	}
}
