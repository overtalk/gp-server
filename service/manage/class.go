package manage

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils/transform/xorm"
)

// GetClasses : get all classes
func (m *BackStageManage) GetClasses(r *http.Request) proto.Message {
	req := &protocol.GetClassesReq{}
	resp := &protocol.GetClassesResp{Status: &protocol.Status{}}

	status := m.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	classes, err := m.db.GetClasses(req.PageNum, req.PageIndex)
	if err != nil {
		logger.Sugar.Errorf("failed to get classes : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get classes"
		return resp
	}

	for _, class := range classes {
		resp.Classes = append(resp.Classes, transform.TurnIntactClassToMinProto(class))
	}

	// get all number
	classesNum, err := m.db.GetClassNum()
	if err != nil {
		logger.Sugar.Errorf("failed to get the number of classes : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get the number of classes"
		return resp
	}

	resp.Total = classesNum
	resp.PageIndex = req.PageIndex
	resp.PageNum = req.PageNum
	return resp
}
