package environment

import (
	"fmt"
	"sort"
)

type EnvironmentSelector struct {
	contents map[string]*Environment
	alias    map[string]string
	revAlias map[string][]string
}

func NewEnvironmentSelector() *EnvironmentSelector {
	return &EnvironmentSelector{contents: map[string]*Environment{}, alias: map[string]string{}, revAlias: map[string][]string{}}
}

func (e *EnvironmentSelector) Select(key string) *Environment {
	environment, ok := e.contents[key]
	if ok {
		return environment
	}

	aliasDst, ok := e.alias[key]
	if ok {
		return e.contents[aliasDst]
	}

	return nil
}

func (e *EnvironmentSelector) Add(key string, environment *Environment) error {
	if e.Has(key) {
		return fmt.Errorf("duplicate environment key exists: %s", key)
	}

	e.contents[key] = environment
	return nil
}

func (e *EnvironmentSelector) AddAlias(aliasKey, key string) error {
	if !e.Has(key) {
		return fmt.Errorf("the key is not exists: %s", key)
	}

	if e.Has(aliasKey) {
		return fmt.Errorf("duplicate environment key exists: %s", aliasKey)
	}

	e.alias[aliasKey] = key
	e.revAlias[key] = append(e.revAlias[key], aliasKey)

	return nil
}

func (e *EnvironmentSelector) Has(key string) bool {
	_, ok := e.contents[key]
	_, aok := e.alias[key]
	return ok || aok
}

func (e *EnvironmentSelector) Keys() []string {
	environments := []*Environment{}
	for _, env := range e.contents {
		environments = append(environments, env)
	}

	sort.Slice(environments, func(i, j int) bool {
		return environments[i].LanguageCode < environments[j].LanguageCode
	})

	keys := []string{}
	for _, environment := range environments {
		keys = append(keys, environment.Key)
	}

	return keys
}

func (e *EnvironmentSelector) Aliases(key string) []string {
	return e.revAlias[key]
}
