package internal

import (
	"github.com/Uh-little-less-dum/dum-commit/internal/findup"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type CommitOptions struct {
	scope        []string
	commitType   []string
	SelectHeight int64
}

// Convers a string into a huh.Option
func getSelectOption(val string) huh.Option[string] {
	return huh.Option[string]{
		Value: val,
		Key:   val,
	}
}

// Returns an array of huh.Option for an array of strings
func (c *CommitOptions) GetSelectOpts(optStrings []string) []huh.Option[string] {
	var opts []huh.Option[string]
	for _, o := range optStrings {
		opts = append(opts, getSelectOption(o))
	}
	return opts
}

func (c *CommitOptions) getOptions(pkgJsonPath string) {
	v := viper.GetViper()

	scopeOpts := v.GetStringSlice("dum-commit.scope.opts")
	typeOpts := v.GetStringSlice("dum-commit.type.opts")
	if len(scopeOpts) > 0 {
		c.scope = scopeOpts
	}

	if len(typeOpts) > 0 {
		c.commitType = typeOpts
	}

}

// TODO: Implement viper flags here.
// Returns options for structured commit based on the root package.json file.
func GetOptionsData() CommitOptions {
	pkgJson, err := findup.Find("package.json")
	if err != nil {
		log.Fatal(err)
	}
	opts := CommitOptions{}
	opts.getOptions(pkgJson)
	return opts
}
