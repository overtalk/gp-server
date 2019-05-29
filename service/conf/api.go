package conf

import (
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/protocol"
	"github.com/qinhan-shu/gp-server/utils"
)

// GetConfig : get all config
func (a *Conf) GetConfig(r *http.Request) proto.Message {
	resp := &protocol.Config{Status: &protocol.Status{}}
	token, err := utils.GetToken(r)
	if err != nil {
		logger.Sugar.Infof("missing token : %v", err)
		resp.Status.Code = protocol.Code_NO_TOKEN
		resp.Status.Message = "missing token"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Infof("invalid token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	difficulty, err := a.db.GetAllDifficulty()
	if err != nil {
		logger.Sugar.Errorf("failed to get all difficulty : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get all difficulty"
		return resp
	}
	d := make(map[int64]string)
	for _, v := range difficulty {
		d[int64(v.Id)] = v.Detail
	}
	resp.Difficulty = d

	tags, err := a.db.GetAllTag()
	if err != nil {
		logger.Sugar.Errorf("failed to get all tags : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get all tags"
		return resp
	}
	t := make(map[int64]string)
	for _, v := range tags {
		t[int64(v.Id)] = v.Detail
	}
	resp.Tags = t

	return resp
}

// GetUserRole : get user role
func (a *Conf) GetUserRole(r *http.Request) proto.Message {
	resp := &protocol.UserRole{Status: &protocol.Status{}}
	token, err := utils.GetToken(r)
	if err != nil {
		logger.Sugar.Infof("missing token : %v", err)
		resp.Status.Code = protocol.Code_NO_TOKEN
		resp.Status.Message = "missing token"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Infof("invalid token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	roles, err := a.db.GetAllRole()
	if err != nil {
		logger.Sugar.Errorf("failed to get all difficulty : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get all difficulty"
		return resp
	}
	d := make(map[int64]string)
	for _, v := range roles {
		d[int64(v.Id)] = v.Detail
	}
	resp.Role = d

	return resp
}

// GetAllLanguage : get all languages
func (a *Conf) GetAllLanguage(r *http.Request) proto.Message {
	resp := &protocol.JudgeLanguage{Status: &protocol.Status{}}
	token, err := utils.GetToken(r)
	if err != nil {
		logger.Sugar.Infof("missing token : %v", err)
		resp.Status.Code = protocol.Code_NO_TOKEN
		resp.Status.Message = "missing token"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Infof("invalid token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	languages, err := a.db.GetAllLanguage()
	if err != nil {
		logger.Sugar.Errorf("failed to get all languages : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "failed to get all languages"
		return resp
	}
	d := make(map[int64]string)
	for _, v := range languages {
		d[int64(v.Id)] = v.Detail
	}
	resp.Language = d

	return resp
}

// GetJudgeResult : get all judge result
func (a *Conf) GetJudgeResult(r *http.Request) proto.Message {
	resp := &protocol.JudgeResults{Status: &protocol.Status{}}
	token, err := utils.GetToken(r)
	if err != nil {
		logger.Sugar.Infof("missing token : %v", err)
		resp.Status.Code = protocol.Code_NO_TOKEN
		resp.Status.Message = "missing token"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Infof("invalid token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	results := make(map[int64]string)

	results[0] = "SUCCESS"
	results[-1] = "WRONG_ANSWER"
	results[1] = "CPU_TIME_LIMIT_EXCEEDED"
	results[2] = "REAL_TIME_LIMIT_EXCEEDED"
	results[3] = "MEMORY_LIMIT_EXCEEDED"
	results[4] = "RUNTIME_ERROR"
	results[5] = "SYSTEM_ERROR"
	results[6] = "COMPILE_ERROR"

	resp.JudgeResults = results

	return resp
}

// GetAlgorithm : get all algorithms
func (a *Conf) GetAlgorithm(r *http.Request) proto.Message {
	resp := &protocol.PaperComposeAlgorithm{Status: &protocol.Status{}}
	token, err := utils.GetToken(r)
	if err != nil {
		logger.Sugar.Infof("missing token : %v", err)
		resp.Status.Code = protocol.Code_NO_TOKEN
		resp.Status.Message = "missing token"
		return resp
	}

	// check token
	_, err = a.cache.GetUserIDByToken(token)
	if err != nil {
		logger.Sugar.Infof("invalid token : %v", err)
		resp.Status.Code = protocol.Code_UNAUTHORIZATED
		resp.Status.Message = "invalid token"
		return resp
	}

	algorithms, err := a.db.GetAlgorithm()
	if err != nil {
		logger.Sugar.Infof("failed to get all algorithms : %v", err)
		resp.Status.Code = protocol.Code_INTERNAL
		resp.Status.Message = "invalid token"
		return resp
	}

	temp := make(map[int64]string)
	for _, v := range algorithms {
		temp[int64(v.Id)] = v.Detail
	}

	resp.Algorithm = temp

	return resp
}
