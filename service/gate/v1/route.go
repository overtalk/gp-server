package gate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/golang/protobuf/proto"

	"github.com/QHasaki/gp-server/logger"
	"github.com/QHasaki/gp-server/protocol/v1"
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

		result, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Sugar.Errorf("failed to get resp.body : %v", err)
		}

		if r.Method == "POST" {
			logger.Sugar.Debug("POST")

			req := &protocol.TestRequest{}

			if err := proto.Unmarshal(result, req); err != nil {
				logger.Sugar.Errorf("failed to unmarshal : %v", err)
			} else {
				logger.Sugar.Infof("request from client : %v", req)
			}

			// TODO: get resp, and marshal

			

			fmt.Fprint(w, "Reload success")
		}

	})

	return httpMux
}
