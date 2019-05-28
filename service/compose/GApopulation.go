package compose

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
	"time"
)

//种群，包含多份试卷
type GAPopulation struct {
	papers []GAPaper
}

//初始化种群
func (population *GAPopulation) init(size int, p *protocol.Paper, problems []*model.Problem) {
	for i := 0; i < size; i++ {
		time.Sleep(1)
		var paper GAPaper
		paper.init(i+1, p, problems)
		population.papers = append(population.papers, paper)
	}
}

//获得种群中最优秀的个体
func (population *GAPopulation) getfitness(p *protocol.Paper) GAPaper {
	//没有试卷生成
	var best int
	var bestfitness float64
	best = 0
	bestfitness = population.papers[0].getfitness(p)
	for i := 0; i < len(population.papers); i++ {
		if bestfitness < population.papers[i].getfitness(p) {
			bestfitness = population.papers[i].getfitness(p)
			best = i
		}
	}
	return population.papers[best]
}

//获得种群中最优秀的id
func (population *GAPopulation) getfitnessID(p *protocol.Paper) int {
	var best int
	var bestfitness float64
	best = 0
	bestfitness = population.papers[0].getfitness(p)
	for i := 0; i < len(population.papers); i++ {
		if bestfitness < population.papers[i].getfitness(p) {
			bestfitness = population.papers[i].getfitness(p)
			best = i
		}
	}
	return best
}

//获得种群长度
func (population *GAPopulation) getSize() int {
	return len(population.papers)
}
