package module

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
)

var (
	DefaultRankUpdateInterval       = 10 * time.Second
	DefaultRankCleanInterval        = 24 * time.Hour
	RankRedisKey                    = "rank"
	MaxRanksNum               int64 = 200
)

// RankItem describes the pass number's item
type RankItem struct {
	UserID  int64 // user id
	PassNum int64 // pass num
}

// Rank : rank module
type Rank interface {
	GetRankList(r *http.Request) proto.Message
}
