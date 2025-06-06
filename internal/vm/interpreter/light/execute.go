package light

import (
	"fmt"

	"github.com/git-town/git-town/v20/internal/cli/colors"
	"github.com/git-town/git-town/v20/internal/cli/dialog/components"
	"github.com/git-town/git-town/v20/internal/config"
	"github.com/git-town/git-town/v20/internal/forge/forgedomain"
	"github.com/git-town/git-town/v20/internal/git"
	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	"github.com/git-town/git-town/v20/internal/gohacks/stringslice"
	"github.com/git-town/git-town/v20/internal/vm/program"
	"github.com/git-town/git-town/v20/internal/vm/shared"
	. "github.com/git-town/git-town/v20/pkg/prelude"
)

func Execute(args ExecuteArgs) {
	for {
		nextStep := args.Prog.Pop()
		if nextStep == nil {
			return
		}
		err := nextStep.Run(shared.RunArgs{
			Backend:                         args.Backend,
			BranchInfos:                     None[gitdomain.BranchInfos](),
			Config:                          NewMutable(&args.Config),
			Connector:                       args.Connector,
			DialogTestInputs:                components.NewTestInputs(),
			FinalMessages:                   args.FinalMessages,
			Frontend:                        args.Frontend,
			Git:                             args.Git,
			PrependOpcodes:                  args.Prog.Prepend,
			RegisterUndoablePerennialCommit: nil,
			UpdateInitialSnapshotLocalSHA:   nil,
		})
		if err != nil {
			fmt.Println(colors.Red().Styled("NOTICE: " + err.Error()))
		}
	}
}

type ExecuteArgs struct {
	Backend       gitdomain.RunnerQuerier
	Config        config.ValidatedConfig
	Connector     Option[forgedomain.Connector]
	FinalMessages stringslice.Collector
	Frontend      gitdomain.Runner
	Git           git.Commands
	Prog          program.Program
}
