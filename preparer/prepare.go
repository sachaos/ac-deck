package preparer

import (
	"github.com/sachaos/atcoder/files"
	"github.com/sachaos/atcoder/lib"
	"os"
	"path"
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

func Prepare(contest *lib.Contest, dir string, env *files.Environment, template *template.Template) error {
	dirPath := path.Join(dir, contest.ID)
	err := createDir(dirPath)
	if err != nil {
		return err
	}

	for _, task := range contest.Tasks {
		taskDir := path.Join(dirPath, task.ID)
		err := createDir(taskDir)
		if err != nil {
			return err
		}

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
				TaskName:   task.Name,
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
