package environment

var DefaultEnvironmentSelector = NewEnvironmentSelector()
var DefaultOldEnvironmentSelector = NewEnvironmentSelector()

func init() {
	for key, environment := range OldEnvironments {
		err := DefaultOldEnvironmentSelector.Add(key, environment)
		if err != nil {
			panic(err.Error())
		}
	}

	for aliasKey, key := range OldAliases {
		err := DefaultOldEnvironmentSelector.AddAlias(aliasKey, key)
		if err != nil {
			panic(err.Error())
		}
	}

	for key, environment := range Environments {
		err := DefaultEnvironmentSelector.Add(key, environment)
		if err != nil {
			panic(err.Error())
		}
	}

	for aliasKey, key := range Aliases {
		err := DefaultEnvironmentSelector.AddAlias(aliasKey, key)
		if err != nil {
			panic(err.Error())
		}
	}
}
