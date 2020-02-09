package files

import "os"

func createFile(fpath string) (*os.File, error) {
	return os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE, 0666)
}

