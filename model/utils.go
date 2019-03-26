package model

import (
	"github.com/qinhan-shu/gp-server/model/xorm"
)

// IntactProblem : intact problem
type IntactProblem struct {
	Detail          *model.Problem
	InAndOutExample []*model.TestData
	Tags            []*model.ProblemTag
}

// IntactClass : intact class
type IntactClass struct {
	Class         *model.Class
	Announcements []*model.Announcement
	TutorName     string
	Publisher     []string
}
