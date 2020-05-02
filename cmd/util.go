package cmd

import (
	"github.com/sachaos/ac-deck/lib/environment"
)

func validateLanguage(lang string) bool {
	return environment.DefaultEnvironmentSelector.Has(lang) || environment.DefaultOldEnvironmentSelector.Has(lang)
}