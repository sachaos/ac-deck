package environment

var DefaultEnvironmentSelector = NewEnvironmentSelector()


func init() {
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
