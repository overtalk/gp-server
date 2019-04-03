package announcement

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/protocol"
)

// GetAnnouncements : get announcements
func (a *Announcement) GetAnnouncements(r *http.Request) proto.Message {
	req := &protocol.AnnouncementsReq{}
	resp := &protocol.AnnouncementsResp{Status: &protocol.Status{}}

	_, status := a.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	announcements, err := a.db.GetGlobalAnnouncements(req.PageNum, req.PageIndex)
	if err != nil {
		logger.Sugar.Errorf("failed to get global announcements : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get global announcements"
		return resp
	}
	total, err := a.db.GetGlobalAnnouncementsNum()
	if err != nil {
		logger.Sugar.Errorf("failed to get the num of global announcements : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get the num of global announcements"
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
	req := &protocol.AnnouncementDetailReq{}
	resp := &protocol.AnnouncementDetailResp{Status: &protocol.Status{}}

	_, status := a.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	announcement, err := a.db.GetAnnouncementDetail(req.Id)
	if err != nil {
		logger.Sugar.Errorf("failed to get announcement detail by id : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get announcement detail by id"
		return resp
	}

	resp.Announcement = transform.AnnouncementToProto(announcement)
	return resp
}

// AddAnnouncement : add announcement
func (a *Announcement) AddAnnouncement(r *http.Request) proto.Message {
	req := &protocol.AddAnnouncementReq{}
	resp := &protocol.AddAnnouncementResp{Status: &protocol.Status{}}

	user, status := a.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	announcement := transform.ProtoToAnnouncement(req.Announcement)
	announcement.CreateTime = time.Now().Unix()
	announcement.LastUpdateTime = time.Now().Unix()
	announcement.Publisher = user.Id
	if err := a.db.AddAnnouncement(announcement); err != nil {
		logger.Sugar.Errorf("failed to add global announcement : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to add global announcement"
		return resp
	}

	resp.IsSuccess = true
	return resp
}

// EditAnnouncement : modify announcement
func (a *Announcement) EditAnnouncement(r *http.Request) proto.Message {
	req := &protocol.EditAnnouncementReq{}
	resp := &protocol.EditAnnouncementResp{Status: &protocol.Status{}}

	user, status := a.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	announcement := transform.ProtoToAnnouncement(req.Announcement)
	announcement.LastUpdateTime = time.Now().Unix()
	announcement.Publisher = user.Id
	if err := a.db.EditAnnouncement(announcement); err != nil {
		logger.Sugar.Errorf("failed to edit global announcement : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to edit global announcement"
		return resp
	}

	resp.IsSuccess = true
	return resp
}

// DelAnnouncement : del announcement
func (a *Announcement) DelAnnouncement(r *http.Request) proto.Message {
	req := &protocol.DelAnnouncementReq{}
	resp := &protocol.DelAnnouncementResp{Status: &protocol.Status{}}

	_, status := a.checkArgsandAuth(r, req)
	if status != nil {
		logger.Sugar.Error(status.Message)
		resp.Status = status
		return resp
	}

	if err := a.db.DelAnnouncement(req.Id); err != nil {
		logger.Sugar.Errorf("failed to del global announcement : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to del global announcement"
		return resp
	}

	resp.IsSuccess = true
	return resp
}
