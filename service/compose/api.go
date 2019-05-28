package compose

import (
	//"encoding/json"
	"errors"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
	// "github.com/qinhan-shu/gp-server/logger"
)

var (
	// ErrIllegalRange num is too big
	ErrIllegalRange = errors.New("illegal random number range [ papernum > num ]")
)

func (c *Compose) IntelligentCompose(problems []*model.Problem, paper *protocol.Paper) ([]*model.PaperProblem, error) {
	// logger.Sugar.Errorf("meiyou tag...")
	// return nil, err
	tag := paper.GetTags()
	nums := 20 * len(tag)
	if nums < int(paper.GetProblemNum()) {
		return nil, ErrIllegalRange
	}
	p := make([]*model.PaperProblem, 0)
	var population GAPopulation
	var ga GA
	population.init(50, paper, problems) //初始化种群，设定为50个试卷
	for i := 0; i < 100; i++ {
		if population.getfitness(paper).fitness > 0.9 {
			break
		}
		ga.evolvePopulation(population, paper, problems)
	}
	newpaper := population.papers[population.getfitnessID(paper)]
	for i := 0; i < len(newpaper.problems); i++ {
		p = append(p, &model.PaperProblem{
			Index:     i + 1,
			ProblemId: int64(newpaper.problems[i].id),
		})
	}

	//s := new([]int64)
	//json.Unmarshal(problems[0].tags, s)
	return p, nil
}
