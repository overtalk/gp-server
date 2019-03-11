package model

// UserProblem : teble `user_problem`
type UserProblem struct {
	UserID      int    `gorm:"type : int(20); not null"`
	ProblemID   string `gorm:"type : varchar(100); not null"`
	PassTime    int    `gorm:"type : int(20); not null; default : 0"`
	RefusedTime int    `gorm:"type : int(20); not null; default : 0"`
}
