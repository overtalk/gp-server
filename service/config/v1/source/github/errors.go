package source

type ErrGetConfFail struct {
	info string
}

func (e ErrGetConfFail) Error() string {
	return "failed to get config from github, error message [ " + e.info + " ]"
}
