package cmd

import "github.com/sachaos/atcoder/lib/files"

func validateLanguage(lang string) bool {
	for key := range files.Environments {
		if key == lang {
			return true
		}
	}

	return false
}