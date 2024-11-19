package internal

import (
	cli_styles "github.com/Uh-little-less-dum/dum-commit/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

const (
	ScopeKey      string = "scope"
	CommitTypeKey string = "commitType"
	MsgKey        string = "msg"
	DetailsKey    string = "details"
)

type FormFields struct {
	scope, commitType, msgInput, detailsInput, confirm huh.Field
}

func (f *FormFields) GetFields(formOpts CommitOptions) {
	v := viper.GetViper()
	if len(formOpts.scope) == 0 {
		f.scope = huh.NewInput().Title("Scope").Placeholder("scope").Key(ScopeKey)
	} else {
		s := huh.NewSelect[string]()
		s.Options(formOpts.GetSelectOpts(formOpts.scope)...).Title("Scope").Key(ScopeKey).WithHeight(v.GetInt("dum-commit.scope.height"))
		f.scope = s
	}

	if len(formOpts.commitType) == 0 {
		f.commitType = huh.NewInput().Title("Type").Placeholder("wip").Key(CommitTypeKey)
	} else {
		s := huh.NewSelect[string]()
		s.Options(formOpts.GetSelectOpts(formOpts.commitType)...).Title("Type").Key(CommitTypeKey).WithHeight(v.GetInt("dum-commit.type.height"))
		f.commitType = s
	}
	f.msgInput = huh.NewInput().Placeholder("Message").Key(MsgKey)
	f.detailsInput = huh.NewText().Placeholder("Details").Key(DetailsKey)
	f.confirm = huh.NewConfirm().Title("Would you like to make this commit?")
}

func NewCommitForm() *huh.Form {
	opts := GetOptionsData()
	log.Info(opts)
	fields := FormFields{}
	fields.GetFields(opts)
	return huh.NewForm(
		huh.NewGroup(
			fields.commitType,
			fields.scope,
			fields.msgInput,
			fields.detailsInput,
		),
		huh.NewGroup(
			fields.confirm,
		),
	).WithTheme(cli_styles.GetHuhTheme()).WithProgramOptions(tea.WithAltScreen())
}
