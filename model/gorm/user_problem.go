package model

// UserProblem : teble `user_problem`
type UserProblem struct {
	UserID      int64 `gorm:"type : int(20); not null"`
	ProblemID   int64 `gorm:"type : int(20); not null"`
	PassTime    int   `gorm:"type : int(20); not null; default : 0"`
	RefusedTime int   `gorm:"type : int(20); not null; default : 0"`
}
