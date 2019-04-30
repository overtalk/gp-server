package db

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
)

// GetDifficultyAnalysis : analysis data by difficulty
func (m *MysqlDriver) GetDifficultyAnalysis(userID, startTime, endTime int64) (map[int64]float64, map[int64]int64, error) {
	userProblems := make([]*transform.UserProblemWithDifficulty, 0)
	if err := m.conn.
		Join("INNER", "problem", "user_problem.problem_id = problem.id").
		Where("user_id = ? and user_problem.submit_time > ? and user_problem.submit_time < ?", userID, startTime, endTime).
		Find(&userProblems); err != nil {
		logger.Sugar.Errorf("failed to get submit records : %v", err)
		return nil, nil, err
	}

	// get the difficulty num
	allDifficulty, err := m.GetAllDifficulty()
	if err != nil {
		logger.Sugar.Errorf("failed to get all difficulty : %v", err)
		return nil, nil, err
	}
	num := len(allDifficulty)

	line := make(map[int64]float64)
	pie := make(map[int64]int64)
	submitNum := make(map[int64]int64)
	for _, difficulty := range allDifficulty {
		line[int64(difficulty.Id)] = 0
		pie[int64(difficulty.Id)] = 0
		submitNum[int64(difficulty.Id)] = 0
	}

	for _, userProblem := range userProblems {
		// fmt.Printf("Difficulty = %d, isPass = %d\n", userProblem.Difficulty, userProblem.Ispass)
		submitNum[int64(userProblem.Difficulty)]++
		pie[int64(userProblem.Difficulty)] += int64(userProblem.Ispass)
	}
	// fmt.Println("passTime : ", pie)
	// fmt.Println("submitNum : ", submitNum)

	// get pass rate
	for i := 1; i <= num; i++ {
		if submitNum[int64(i)] == 0 {
			line[int64(i)] = 0
		} else {
			line[int64(i)] = math.Trunc((float64(pie[int64(i)])/float64(submitNum[int64(i)]))*1e2+0.5) * 1e-2
		}
	}

	return line, pie, nil
}

// GetTagsAnalysis : analysis data by tags
func (m *MysqlDriver) GetTagsAnalysis(userID, startTime, endTime int64, tags []int64) (map[int64]float64, map[int64]int64, error) {
	if len(tags) == 0 {
		return nil, nil, fmt.Errorf("no tags")
	}
	// get tags
	var tagArg []interface{}
	var like []string
	tagArg = append(tagArg, userID)
	tagArg = append(tagArg, startTime)
	tagArg = append(tagArg, endTime)
	for _, tag := range tags {
		like = append(like, " tags like ? ")
		like = append(like, " tags like ? ")
		like = append(like, " tags like ? ")
		like = append(like, " tags like ? ")
		tagArg = append(tagArg, "%"+fmt.Sprintf(",%d,", tag)+"%")
		tagArg = append(tagArg, fmt.Sprintf("[%d,", tag)+"%")
		tagArg = append(tagArg, "%"+fmt.Sprintf(",%d]", tag))
		tagArg = append(tagArg, fmt.Sprintf("[%d]", tag))
	}

	// get userProblems
	userProblems := make([]*transform.UserProblemWithTags, 0)
	if err := m.conn.
		Join("INNER", "problem", "user_problem.problem_id = problem.id").
		Where("user_id = ? and user_problem.submit_time > ? and user_problem.submit_time < ? and ( "+strings.Join(like, "||")+" )", tagArg...).
		Find(&userProblems); err != nil {
		return nil, nil, err
	}

	// get pass tate and pass num
	line := make(map[int64]float64)
	pie := make(map[int64]int64)
	submitNum := make(map[int64]int64)
	for _, tag := range tags {
		line[tag] = 0
		pie[tag] = 0
		submitNum[tag] = 0
	}

	// fmt.Println(len(userProblems))
	for _, userProblem := range userProblems {
		t := new([]int64)
		if err := json.Unmarshal([]byte(userProblem.Tags), t); err != nil {
			logger.Sugar.Error("failed to get tags of problem[ id = %d ] : %v", userProblem.ProblemId, err)
			continue
		}

		var requiredTag []int64
		for _, t1 := range tags {
			for _, t2 := range *t {
				if t1 == t2 {
					requiredTag = append(requiredTag, t1)
					break
				}
			}
		}

		// fmt.Println(userProblem.Id, userProblem.Tags)
		// fmt.Println("匹配上的tag : ", requiredTag)

		for _, tag := range requiredTag {
			submitNum[tag]++
			pie[tag] += int64(userProblem.Ispass)
		}
	}

	// get pass rate
	for _, tag := range tags {
		if submitNum[tag] == 0 {
			line[tag] = 0
		} else {
			line[tag] = math.Trunc((float64(pie[tag])/float64(submitNum[tag]))*1e2+0.5) * 1e-2
		}
	}

	return line, pie, nil
}
