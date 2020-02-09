package files

import (
	"github.com/sachaos/atcoder/lib"
	"gopkg.in/yaml.v2"
	"os"
	"path"
)

const TESTDATA_NAME = "testdata.yaml"

func WriteTestData(dir string, examples []*lib.Example) error {
	p := path.Join(dir, TESTDATA_NAME)
	file, err := createFile(p)
	if err != nil {
		return err
	}
	defer file.Close()

	return yaml.NewEncoder(file).Encode(examples)
}

func LoadTestData(dir string) ([]*lib.Example, error) {
	p := path.Join(dir, TESTDATA_NAME)
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var examples []*lib.Example
	err = yaml.NewDecoder(file).Decode(&examples)
	if err != nil {
		return nil, err
	}

	return examples, nil
}
