package atcoder

type Contest struct {
	ID string
	URL string
	Tasks []*Task
}

type Task struct {
	ID string
	Name string
	URL  string
	Examples []*Example
}

type Example struct {
	In string
	Exp string
}
