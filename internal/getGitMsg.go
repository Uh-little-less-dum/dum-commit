package internal

import (
	"github.com/charmbracelet/log"
)

func GetGitCommit() map[FormKey]string {
	form := NewCommitForm()
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	var data = make(map[FormKey]string)
	data[CommitTypeKey] = form.GetString(string(CommitTypeKey))
	data[ScopeKey] = form.GetString(string(ScopeKey))
	data[MsgKey] = form.GetString(string(MsgKey))
	data[DetailsKey] = form.GetString(string(DetailsKey))
	return data
}
