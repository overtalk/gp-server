package db

import (
	"encoding/json"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
)

// GetDifficultyAnalysis : analysis data by difficulty
func (m *MysqlDriver) GetDifficultyAnalysis(userID, startTime, endTime int64) ([]int64, []int64, error) {
	userProblems := make([]*transform.UserProblemWithDifficulty, 0)
	if err := m.conn.
		Join("INNER", "problem", "user_problem.problem_id = problem.id").
		Where("user_id = ? and submit_time > ? and submit_time < ?", userID, startTime, endTime).
		Find(&userProblems); err != nil {
		return nil, nil, err
	}
	logger.Sugar.Error(userProblems)
	return nil, nil, nil
}

// GetTagsAnalysis : analysis data by tags
func (m *MysqlDriver) GetTagsAnalysis(userID, startTime, endTime int64, tags []int64) ([]int64, []int64, error) {
	userProblems := make([]*transform.UserProblemWithTags, 0)
	if err := m.conn.
		Join("INNER", "problem", "user_problem.problem_id = problem.id").
		Where("user_id = ? and submit_time > ? and submit_time < ?", userID, startTime, endTime).
		Find(&userProblems); err != nil {
		return nil, nil, err
	}
	logger.Sugar.Error(userProblems)
	for _, userProblem := range userProblems {
		tags := new([]int64)
		if err := json.Unmarshal([]byte(userProblem.Tags), tags); err != nil {
			logger.Sugar.Error("failed to get tags of problem[ id = %d ] : %v", userProblem.ProblemId, err)
			continue
		}
		// TODO: analysis
	}
	return nil, nil, nil
}
