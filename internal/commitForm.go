package internal

import (
	cli_styles "github.com/Uh-little-less-dum/dum-commit/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/spf13/viper"
)

type FormKey string

const (
	ScopeKey      FormKey = "scope"
	CommitTypeKey FormKey = "commitType"
	MsgKey        FormKey = "msg"
	DetailsKey    FormKey = "details"
)

type FormFields struct {
	scope, commitType, msgInput, detailsInput, confirm huh.Field
}

func (f *FormFields) GetFields(formOpts CommitOptions) {
	v := viper.GetViper()
	if len(formOpts.scope) == 0 {
		f.scope = huh.NewInput().Title("Scope").Placeholder("scope").Key(string(ScopeKey))
	} else {
		s := huh.NewSelect[string]()
		s.Options(formOpts.GetSelectOpts(formOpts.scope)...).Title("Scope").Key(string(ScopeKey)).WithHeight(v.GetInt("dum-commit.scope.height"))
		f.scope = s
	}

	if len(formOpts.commitType) == 0 {
		f.commitType = huh.NewInput().Title("Type").Placeholder("wip").Key(string(CommitTypeKey))
	} else {
		s := huh.NewSelect[string]()
		s.Options(formOpts.GetSelectOpts(formOpts.commitType)...).Title("Type").Key(string(CommitTypeKey)).WithHeight(v.GetInt("dum-commit.type.height"))
		f.commitType = s
	}
	f.msgInput = huh.NewInput().Placeholder("Message").Key(string(MsgKey))
	f.detailsInput = huh.NewText().Placeholder("Details").Key(string(DetailsKey))
	f.confirm = huh.NewConfirm().Title("Would you like to make this commit?")
}

func NewCommitForm() *huh.Form {
	opts := GetOptionsData()
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
