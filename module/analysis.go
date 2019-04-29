package module

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Analysis : analysis module
type Analysis interface {
	DifficultyAnalysis(r *http.Request) proto.Message
	TagsAnalysis(r *http.Request) proto.Message
}
