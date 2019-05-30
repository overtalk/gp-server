package compose

import (
	"math/rand"
	"time"

	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

//随机组卷算法
func RandomAlgorithm(p *protocol.Paper, problems []*model.Problem) GAPaper {
	var paper GAPaper
	paper.id = 1
	var problem GAproblem
	//首先获取试卷要求参数
	difficulty := p.GetDifficulty() //试卷难度要求
	problemnum := p.GetProblemNum() //试卷数量要求
	tags := p.GetTags()             //试卷知识点要求
	diff := make([]int64, 5)        //每个级别难度数量
	num := len(problems)
	diff = GetDifficultyDistribution(difficulty, problemnum)
	rand.Seed(time.Now().UnixNano()) //设置时间种子
	for problemnum > 0 {
		r := rand.Intn(num)
		time.Sleep(1)
		problem.Init(r+1, problems)
		tag := problem.tags
		d := problem.difficulty
		t := problem.lastused
		//查重判断
		isrepeat := false
		for i := 0; i < len(paper.problems); i++ {
			if paper.problems[i].id == r+1 {
				isrepeat = true
				break
			}
		}
		if isrepeat == true {
			continue
		}
		//判断是否满足时间要求
		if !isRightTime(t) {
			continue
		}
		//判断是否满足知识点要求
		if !isRightTag(tags, tag) {
			continue
		}
		//判断是否满足难度要求
		switch d {
		case 1:
			if diff[0] > 0 {
				paper.problems = append(paper.problems, problem)
				diff[0]--
				problemnum--
			} else {
				continue
			}
		case 2:
			if diff[1] > 0 {
				paper.problems = append(paper.problems, problem)
				diff[1]--
				problemnum--
			} else {
				continue
			}
		case 3:
			if diff[2] > 0 {
				paper.problems = append(paper.problems, problem)
				diff[2]--
				problemnum--
			} else {
				continue
			}
		case 4:
			if diff[3] > 0 {
				paper.problems = append(paper.problems, problem)
				diff[3]--
				problemnum--
			} else {
				continue
			}
		case 5:
			if diff[4] > 0 {
				paper.problems = append(paper.problems, problem)
				diff[4]--
				problemnum--
			} else {
				continue
			}
		}
	}
	return paper
}

//获取试卷难度分布
func GetDifficultyDistribution(difficulty int64, problemnum int64) []int64 {
	diff := make([]int64, 5)
	var easy1, easy2, middle, difficult2, difficult1 int64
	switch difficulty {
	case 1:
		easy1 = int64(float64(problemnum) * 0.6)
		easy2 = int64(float64(problemnum) * 0.33)
		middle = problemnum - easy1 - easy2
	case 2:
		easy1 = int64(float64(problemnum) * 0.25)
		easy2 = int64(float64(problemnum) * 0.45)
		middle = int64(float64(problemnum) * 0.25)
		difficult2 = problemnum - easy1 - easy2 - middle
	case 3:
		easy1 = int64(float64(problemnum) * 0.05)
		easy2 = int64(float64(problemnum) * 0.25)
		middle = int64(float64(problemnum) * 0.4)
		difficult2 = int64(float64(problemnum) * 0.25)
		difficult1 = problemnum - easy1 - easy2 - middle - difficult2
	case 4:
		difficult1 = int64(float64(problemnum) * 0.25)
		difficult2 = int64(float64(problemnum) * 0.45)
		middle = int64(float64(problemnum) * 0.25)
		easy2 = problemnum - difficult1 - difficult2 - middle
	case 5:
		difficult1 = int64(float64(problemnum) * 0.6)
		difficult2 = int64(float64(problemnum) * 0.33)
		middle = problemnum - difficult1 - difficult2
	}
	diff[0] = easy1
	diff[1] = easy2
	diff[2] = middle
	diff[3] = difficult2
	diff[4] = difficult1
	return diff
}

/*是否满足知识点要求
tags:试卷知识点要求
tag：题目知识点
*/
func isRightTag(tags []int64, tag []int64) bool {
	var b bool
	b = false
	for i := 0; i < len(tag); i++ {
		for j := 0; j < len(tags); j++ {
			if tag[i] == tags[j] {
				b = true
			}
		}
	}
	return b
}

//对时间判断
func isRightTime(t int64) bool {
	nowtime := time.Now().Unix()
	oneyear := 31536000
	if int64(nowtime-t) > int64(oneyear) {
		return true
	}
	return false
}
