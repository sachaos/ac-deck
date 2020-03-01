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

	Note string `yaml:"-"`
}

