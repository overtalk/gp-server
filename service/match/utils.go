package match

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

func (m *Match) checkArgsandAuth(r *http.Request, req proto.Message) (*model.User, *protocol.Status) {
	data, token, err := utils.GetReqAndToken(r)
	if err != nil {
		return nil, &protocol.Status{
			Code:    protocol.Code_DATA_LOSE,
			Message: err.Error(),
		}
	}
	if err := proto.Unmarshal(data, req); err != nil {
		return nil, &protocol.Status{
			Code:    protocol.Code_DATA_LOSE,
			Message: "failed to unmarshal request body",
		}
	}

	// check token
	userID, err := m.cache.GetUserIDByToken(token)
	if err != nil {
		return nil, &protocol.Status{
			Code:    protocol.Code_UNAUTHORIZATED,
			Message: "invalid token",
		}
	}

	// get user from db, and get the operation auth of the player
	user, err := m.db.GetUserByID(userID)
	if err != nil {
		return nil, &protocol.Status{
			Code:    protocol.Code_UNAUTHORIZATED,
			Message: "failed to get user info",
		}
	}

	if user.Role != int(protocol.Role_MANAGER) {
		return nil, &protocol.Status{
			Code:    protocol.Code_PERMISSION_DENIED,
			Message: "permission denied",
		}
	}
	return user, nil
}

func (m *Match) newPaper(paper *transform.Paper) {
	p := make([]*model.PaperProblem, 0)
	p = append(p, &model.PaperProblem{
		Index:     1,
		ProblemId: 1,
	})
	p = append(p, &model.PaperProblem{
		Index:     2,
		ProblemId: 1,
	})
}
