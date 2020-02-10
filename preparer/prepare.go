package preparer

import (
	"fmt"
	"github.com/rakyll/statik/fs"
	"github.com/sachaos/atcoder/files"
	"github.com/sachaos/atcoder/lib"
	_ "github.com/sachaos/atcoder/statik"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

func createDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0777)
		if err != nil {
			return err
		}
	}

	return nil
}

func createFile(fpath string) (*os.File, error) {
	return os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE, 0666)
}

type TemplateData struct {
	Task *lib.Task
}

func prepareTemplate(p string) (*template.Template, error) {
	split := strings.Split(p, "/")
	var file io.ReadCloser
	var err error
	if split[0] == "internal" {
		f, err := fs.New()
		if err != nil {
			return nil, err
		}

		file, err = f.Open(strings.TrimPrefix(p, "internal"))
		if err != nil {
			return nil, fmt.Errorf("internal not found: %w", err)
		}
		defer file.Close()
	} else {
		file, err = os.Open(p)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return template.New("src").Parse(string(all))
}

func Prepare(contest *lib.Contest, dir string, env *files.Environment) error {
	template, err := prepareTemplate(env.Template)
	if err != nil {
		return err
	}

	dirPath := path.Join(dir, contest.ID)
	err = createDir(dirPath)
	if err != nil {
		return err
	}

	for _, task := range contest.Tasks {
		taskDir := path.Join(dirPath, task.ID)
		err := createDir(taskDir)
		if err != nil {
			return err
		}

		fmt.Printf("Generating testdata & source on %s\n", taskDir)

		err = files.WriteTestData(taskDir, task.Examples)
		if err != nil {
			return err
		}

		srcFile, err := createFile(path.Join(taskDir, env.SrcName))
		if err != nil {
			return err
		}
		defer srcFile.Close()

		err = files.WriteConf(taskDir, &files.Conf{
			Environment: env,
			AtCoder: &files.AtCoder{
				TaskID:     task.ID,
				TaskName:   task.Name,
				ContestID:  contest.ID,
				ContestURL: contest.URL,
			},
		})
		if err != nil {
			return err
		}

		err = template.Execute(srcFile, TemplateData{Task: task})
		if err != nil {
			return err
		}
	}

	return nil
}
