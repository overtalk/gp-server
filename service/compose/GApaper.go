package compose

import (
	//"encoding/json"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
	//"github.com/qinhan-shu/gp-server/utils"
	"math"
	"math/rand"
	"time"
)

//试卷，里面包含多个问题
type GAPaper struct {
	id              int         //试卷id
	fitness         float64     //适应度函数
	totaldifficulty float64     //整体难度系数
	problems        []GAproblem //试卷包含试题
}

//初始化试卷
//TODO: 连接数据库的数据
func (paper *GAPaper) init(id int, p *protocol.Paper, problems []*model.Problem) {
	paper.id = id
	var papernum, num int //papernum:要求数量,num:题库数量
	papernum = getNum(p)
	num = len(problems)
	r := generateRandomNumber(1, num, papernum, p, problems) //从题库中随机生成要求数量的题目id
	for i := 0; i < papernum; i++ {
		var b GAproblem
		b.Init(r[i], problems)
		paper.problems = append(paper.problems, b)
	}
	paper.totaldifficulty = paper.gettotaldifficulty()
	paper.fitness = paper.getfitness(p)
}

/*
求试卷整体的难度=难度/数量
*/
func (paper *GAPaper) gettotaldifficulty() float64 {
	difficulty := paper.getdifficulty() / float64(len(paper.problems))
	return difficulty
}

/*获得适应度函数
 */
func (paper *GAPaper) getfitness(p *protocol.Paper) float64 {
	var fitness float64
	var diff, diffend, difffront float64 //老师要求试卷难度的期望值,和组成卷子的难度

	switch p.GetDifficulty() {
	case 1:
		difffront = 0.1
	case 2:
		difffront = 0.3
	case 3:
		difffront = 0.5
	case 4:
		difffront = 0.7
	case 5:
		difffront = 0.9
	}
	diffend = paper.gettotaldifficulty()
	diff = math.Abs(diffend - difffront)
	fitness = 1.0 - diff
	return fitness
}

/*
获得难度
1. 难度系数为每题题目难度系数相加
2. 提交数量小于10用默认的难度系数，提交数量大于10用ac/submit
*/
func (paper *GAPaper) getdifficulty() float64 {
	difficulty := 0.0
	for i := 0; i < len(paper.problems); i++ {
		if paper.problems[i].submit < 10 {
			difficulty += float64(paper.problems[i].difficulty)/5.0 - 0.1
		} else {
			d := 1.0 - float64(paper.problems[i].ac/paper.problems[i].submit)
			difficulty += d
		}
	}
	return difficulty
}

//获取试卷题目数量
func getNum(p *protocol.Paper) int {
	return int(p.GetProblemNum())
}

//获取数据库中题目数量
func getSQLNum(problems []*model.Problem) int {
	return len(problems)
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int, p *protocol.Paper, problems []*model.Problem) []int {
	var tags []int64
	tags = p.GetTags()
	var n int //试卷要求的数量
	n = getNum(p)
	num1 := int(n / len(tags)) //均匀分配的数量
	coun := 0
	var t []int //用于存储题目的分段点，例如1 2 |3 4| 5 6 存储2 4,6
	for i := 0; i < len(tags)-1; i++ {
		coun += num1
		t = append(t, coun)
	}
	t = append(t, n)
	n1 := 0
	n2 := len(tags) - 1
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !contain(num, tags[n2], problems) {
			exist = true
		}
		if !exist {
			nums = append(nums, num)
			n1++
			if n1 >= t[len(tags)-1-n2] {
				n2--
			}
		}
	}
	return nums
}

//id1在tag中
func contain(id int, tag int64, problems []*model.Problem) bool {
	var pro GAproblem
	pro.Init(id, problems)
	if isContain(tag, pro.tags) {
		nowtime := time.Now().Unix()
		oneyear := 31536000
		if (nowtime - pro.lastused) > int64(oneyear) {
			return true
		}
	}
	return false
}

func isContain(tag int64, tags []int64) bool {
	for i := 0; i < len(tags); i++ {
		if tags[i] == tag {
			return true
		}
	}
	return false
}
