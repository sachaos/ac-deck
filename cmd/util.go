package cmd

import "github.com/sachaos/atcoder/files"

func validateLanguage(lang string) bool {
	for key := range files.Environments {
		if key == lang {
			return true
		}
	}

	return false
}