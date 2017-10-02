package main

import (
	"log"
	"net/http"

	"github.com/alanfoster/configurable-proxy/analytics"
	"github.com/alanfoster/configurable-proxy/clock"
	"github.com/alanfoster/configurable-proxy/config"
	"github.com/alanfoster/configurable-proxy/stubs"
)

func main() {
	log.Println("Started application")
	cm := config.Manager{
		Stubs: &[]config.Stub{
			{
				Request: config.Request{
					Pattern: "/",
					Headers: map[string]string{},
					Method:  "GET",
				},
				Response: config.Response{
					Status: http.StatusOK,
					Headers: map[string]string{
						"X-Foo-Header": "HelloWorld123",
					},
					Body: "{ mock: true }",
				},
			},
		},
	}

	stats := analytics.NewStats(clock.NewClock())
	statsMiddleware := analytics.NewMiddleware(stats)

	http.Handle("/", statsMiddleware.Middleware(stubs.Handler{
		ConfigManager: &cm,
	}))

	http.Handle("/stats", analytics.NewHandler(stats))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
