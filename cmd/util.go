package cmd

import (
	"github.com/sachaos/atcoder/lib/environment"
)

func validateLanguage(lang string) bool {
	return environment.DefaultEnvironmentSelector.Has(lang)
}