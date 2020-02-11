package files

import (
	"github.com/pelletier/go-toml"
	"os"
	"path"
)

const CONF_NAME = ".task.toml"

type Conf struct {
	Environment *Environment
	AtCoder     *AtCoder
}

type AtCoder struct {
	TaskID     string
	TaskName   string
	TaskURL    string
	ContestID  string
	ContestURL string
}

type Environment struct {
	Language     string
	SrcName      string
	BuildCmd     string
	Cmd          string
	CleanCmd     string
	Template     string
	LanguageCode string
}

// NOTE: https://language-test-201603.contest.atcoder.jp/
var Environments = map[string]*Environment{
	"g++": {
		Language:     "c++",
		SrcName:      "main.cpp",
		BuildCmd:     "g++ -std=gnu++1y -O2 -o a.out main.cpp",
		Cmd:          "./a.out",
		CleanCmd:     "rm ./a.out",
		Template:     "internal/c++/main.cpp",
		LanguageCode: "3003",
	},
	"clang": {
		Language:     "c++",
		SrcName:      "main.cpp",
		BuildCmd:     "clang++ -I/usr/local/include/c++/v1 -L/usr/local/lib -std=c++14 -stdlib=libc++ -O2 -o a.out main.cpp",
		Cmd:          "./a.out",
		CleanCmd:     "rm ./a.out",
		Template:     "internal/c++/main.cpp",
		LanguageCode: "3003",
	},
	"go": {
		Language:     "go",
		SrcName:      "main.go",
		BuildCmd:     "go build -o ./binary main.go",
		Cmd:          "./binary",
		CleanCmd:     "rm ./binary",
		Template:     "internal/go/main.go",
		LanguageCode: "3013",
	},
	"python3": {
		Language:     "python3",
		SrcName:      "main.py",
		Cmd:          "python3 main.py",
		Template:     "internal/python3/main.py",
		LanguageCode: "3023",
	},
	"python2": {
		Language:     "python2",
		SrcName:      "main.py",
		Cmd:          "python2 main.py",
		Template:     "internal/python2/main.py",
		LanguageCode: "3022",
	},
}

func WriteConf(dir string, conf *Conf) error {
	confPath := path.Join(dir, CONF_NAME)
	file, err := createFile(confPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	return encoder.Encode(conf)
}

func LoadConf(dir string) (*Conf, error) {
	var conf Conf
	confPath := path.Join(dir, CONF_NAME)
	file, err := os.Open(confPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = toml.NewDecoder(file).Decode(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
