package gate

import (
	"fmt"
	"net/http"
	"os"
)

func (s *Service) route() http.Handler {
	httpMux := http.NewServeMux()

	httpMux.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		// HTTP CROS Policy
		origin := os.Getenv("Access_Control_Allow_Origin")
		if origin == "" {
			origin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)

		fmt.Fprint(w, "Reload success")
	})

	return httpMux
}
