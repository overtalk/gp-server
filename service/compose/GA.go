package compose

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
	"math/rand"
	"time"
)

type GA struct {
	mutationRate float64 //变异概率
	tournameSize int     //每次进化淘汰数组大小
}

//进化操作，采用轮盘赌策略和最优选择策略的方法
//最优选择策略：适应度最高的个体不参与到下一代的繁殖
func (ga *GA) evolvePopulation(population GAPopulation, p *protocol.Paper, problems []*model.Problem) GAPopulation {
	ga.mutationRate = 0.01
	var newpopulation GAPopulation
	newpopulation.papers = append(newpopulation.papers, population.getfitness(p)) //将适应度最高的个体添加到下一轮进化的第一个位置
	newpopulation.papers[0].id = 0
	//进行交叉操作
	for i := 1; i < len(population.papers); i++ {
		paper1 := selectCross(population, p)
		time.Sleep(1)
		paper2 := selectCross(population, p)
		//两张交叉的试卷重复,则交叉失效，重新进行选择
		for paper1.id == paper2.id {
			time.Sleep(1)
			paper2 = selectCross(population, p)
		}
		paper1 = cross(paper1, paper2, p)
		//将新的试卷更改适应度和难度系数
		paper1.fitness = paper1.getfitness(p)
		paper1.totaldifficulty = paper1.gettotaldifficulty()
		newpopulation.papers = append(newpopulation.papers, paper1)
		newpopulation.papers[i].id = i
	}
	//进行变异操作 to do
	for i := 1; i < len(population.papers); i++ {
		time.Sleep(1)
		r := rand.Float64()
		//如果随机生成的数小于变异概率则发生变异
		if r < ga.mutationRate {
			newpopulation.papers[i] = mutate(newpopulation.papers[i], problems)
			newpopulation.papers[i].fitness = newpopulation.papers[i].getfitness(p)
			newpopulation.papers[i].totaldifficulty = newpopulation.papers[i].gettotaldifficulty()
		}
	}
	return newpopulation
}

//选择操作，使用轮盘赌策略
func selectCross(population GAPopulation, p *protocol.Paper) GAPaper {
	var paper GAPaper
	fitness := make([]float64, len(population.papers))
	fitnessid := population.getfitnessID(p)
	for i := 0; i < len(population.papers); i++ {
		fitness[i] = population.papers[i].getfitness(p)
	}
	selectid := RouletteWheelSelection(fitnessid, fitness)
	paper = population.papers[selectid-1]
	return paper
}

/*
轮盘赌算法
id表示被最优选择排除的试卷个体
*/
func RouletteWheelSelection(id int, fitness []float64) int {
	var selectid int
	var num float64
	selectid = 1
	selection := make([]float64, len(fitness)-1)
	for i := 0; i < len(fitness)-1; i++ {
		if i+1 < id {
			num += fitness[i]
		} else {
			num += fitness[i+1]
		}
		selection[i] = num
	}
	for i := 0; i < len(selection); i++ {
		selection[i] = selection[i] / num
	}
	rand.Seed(time.Now().UnixNano()) //设置时间种子
	a := rand.Float64()
	for i := 0; i < len(selection); i++ {
		if a > selection[i] {
			if i < id {
				selectid = i + 1
			} else {
				selectid = i + 2
			}
		}
	}
	return selectid
}

//交叉运算，使用单点交叉运算进行操作
func cross(paper1 GAPaper, paper2 GAPaper, p *protocol.Paper) GAPaper {
	var p1, p2 GAPaper
	var prob GAproblem
	p1 = paper1
	p2 = paper2

	tags := p.GetTags()
	num := len(tags)

	r := rand.Intn(num)
	var num1 int

	num1 = len(paper1.problems) / num
	num1 = num1 + r*num1

	if r == num-1 {
		num1=rand.Intn(len(p1.problems))
	}

	//交叉运算
	for i := 0; i <= num1; i++ {
		prob = p1.problems[i]
		p1.problems[i] = p2.problems[i]
		p2.problems[i] = prob
	}

	return p1
}

//进行变异操作
func mutate(paper GAPaper, problems []*model.Problem) GAPaper {
	var newpaper GAPaper                 //表示变异后新产生的试卷
	var oldproblem, newproblem GAproblem //表示变异后产生新的试题
	var tag int64                        //表示变异位置的tag

	newpaper = paper
	//用随机数产生变异的为止
	size := len(newpaper.problems)
	r := rand.Intn(size - 1)
	oldproblem = paper.problems[r]
	tag = oldproblem.tags[len(oldproblem.tags)-1]
	//保证题目在tag所在的范围内，且不会重复出现
	newproblem = getNewproblem(oldproblem, tag, problems)
	newpaper.problems[r] = newproblem
	return newpaper
}

//生成新的题目，且不和之前题目重复,且时间在一年之内没有出现过
func getNewproblem(oldproblem GAproblem, tag int64, problems []*model.Problem) GAproblem {
	newproblem := oldproblem
	num := len(problems)
	r := rand.Intn(num)
	newproblem.Init(r+1, problems)
	nowtime := time.Now().Unix()
	for newproblem.id == oldproblem.id || newproblem.tags[len(newproblem.tags)-1] != oldproblem.tags[len(oldproblem.tags)-1] || nowtime-newproblem.lastused < 31536000 {
		time.Sleep(1)
		r = rand.Intn(num)
		newproblem.Init(r+1, problems)
	}
	return newproblem
}
