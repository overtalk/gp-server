package announcement

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// GetAnnouncements : get announcements
func (a *Announcement) GetAnnouncements(r *http.Request) proto.Message {
	req := &protocol.AnnouncementsReq{}
	resp := &protocol.AnnouncementsResp{Status: &protocol.Status{}}
	// get token and data
	data, token, err := utils.GetReqAndToken(r)
	if err != nil {
		logger.Sugar.Error(err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = err.Error()
		return resp
	}
	if err := proto.Unmarshal(data, req); err != nil {
		logger.Sugar.Errorf("failed to unmarshal request body : %v", err)
		resp.Status.Code = protocol.Code_DATA_LOSE
		resp.Status.Message = "failed to unmarshal request body"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Errorf("failed to get token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	announcements, err := a.db.GetGlobalAnnouncements(req.PageNum, req.PageIndex)
	if err != nil {
		logger.Sugar.Errorf("failed to get global announcements : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "ailed to get global announcements"
		return resp
	}
	total, err := a.db.GetGlobalAnnouncementsNum()
	if err != nil {
		logger.Sugar.Errorf("failed to get the num of global announcements : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "ailed to get the num of global announcements"
		return resp
	}

	for _, announcement := range announcements {
		resp.Announcements = append(resp.Announcements, transform.AnnouncementToProto(announcement))
	}

	resp.Total = total
	resp.PageIndex = req.PageIndex
	resp.PageNum = req.PageNum
	return resp
}

// GetDetail : get the detail message of announcement
func (a *Announcement) GetDetail(r *http.Request) proto.Message {
	return nil
}

// AddAnnouncement : add announcement
func (a *Announcement) AddAnnouncement(r *http.Request) proto.Message {
	return nil
}

// EditAnnouncement : modify announcement
func (a *Announcement) EditAnnouncement(r *http.Request) proto.Message {
	return nil
}

// DelAnnouncement : del announcement
func (a *Announcement) DelAnnouncement(r *http.Request) proto.Message {
	return nil
}
