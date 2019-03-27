package source

// ErrGetConfFail : failed to get config details, and pring details info
type ErrGetConfFail struct {
	info string
}

func (e ErrGetConfFail) Error() string {
	return "failed to get config from github, error message [ " + e.info + " ]"
}
