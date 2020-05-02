package files

import (
	"github.com/pelletier/go-toml"
	"github.com/sachaos/ac-deck/lib/environment"
	"os"
	"path"
)

const CONF_NAME = ".task.toml"

type Conf struct {
	Environment *environment.Environment
	AtCoder     *AtCoder
}

type AtCoder struct {
	TaskID     string
	TaskName   string
	TaskURL    string
	ContestID  string
	ContestURL string
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
