package command

import (
	"testing"

	"github.com/hashicorp/consul/command/base"
	"github.com/mitchellh/cli"
)

func testOperatorRaftCommand(t *testing.T) (*cli.MockUi, *OperatorRaftCommand) {
	ui := cli.NewMockUi()
	return ui, &OperatorRaftCommand{
		Command: base.Command{
			UI:    ui,
			Flags: base.FlagSetHTTP,
		},
	}
}

func TestOperator_Raft_Implements(t *testing.T) {
	t.Parallel()
	var _ cli.Command = &OperatorRaftCommand{}
}
