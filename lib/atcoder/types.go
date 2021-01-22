package atcoder

const LangNew = "lang"

type Contest struct {
	ID      string
	URL     string
	LangVer string
	Tasks   []*Task
}

type Task struct {
	ID       string
	Name     string
	URL      string
	Examples []*Example
}

type Example struct {
	In  string
	Exp string
}

type Status struct {
	SubmissionDate string
	Problem        string
	Language       string
	Point          string
	CodeLength     string
	Result         string
	ElapsedTime    string
	Memory         string
}
