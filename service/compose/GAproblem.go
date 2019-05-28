package compose

import (
	"encoding/json"
	"github.com/qinhan-shu/gp-server/model/xorm"
)

//  遗传算法智能组卷问题
type GAproblem struct {
	id         int     //题目id
	difficulty int     //初始化难度(当提交次数小于10时使用)
	tags       []int64 //类型分布
	ac         int     //接受次数
	submit     int     //提交次数
	lastused   int64   //上一次提交时间
}

// 用于组卷初始化题目,这里保证不会出错,不需要返回error,获得题库中为id的问题
func (problem *GAproblem) Init(id int, problems []*model.Problem) {
	problem.id = int(problems[id-1].Id)
	problem.difficulty = problems[id-1].Difficulty
	tags := new([]int64)
	json.Unmarshal([]byte(problems[id-1].Tags), tags) //不会出错
	problem.tags = *tags
	problem.ac = problems[id-1].Ac
	problem.submit = problems[id-1].SubmitTime
	problem.lastused = problems[id-1].LastUsed
}
