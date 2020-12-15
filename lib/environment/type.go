package environment

type Environment struct {
	Key          string
	Language     string
	SrcName      string
	Template     string
	LanguageCode string

	BuildCmd string
	Cmd      string
	CleanCmd string

	DockerImage      string
	BuildCmdOnDocker string
	CmdOnDocker      string

	WorkingDir string
	SrcDir string

	Note string `yaml:"-"`
}

func (e *Environment) GetWorkingDir() string {
	if e.WorkingDir != "" {
		return e.WorkingDir
	}

	return "/src"
}

func (e *Environment) GetSrcDir() string {
	if e.SrcDir != "" {
		return e.SrcDir
	}

	return "/src"
}
