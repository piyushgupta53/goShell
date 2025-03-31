package parser

type Redirection struct {
	Fd     int // 1 = stdout, 2 = stderr
	Append bool
	Target string
}

type Command struct {
	Name         string
	Args         []string
	Redirections []Redirection
}
