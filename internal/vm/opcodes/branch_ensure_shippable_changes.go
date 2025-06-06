package opcodes

import (
	"fmt"

	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	"github.com/git-town/git-town/v20/internal/messages"
	"github.com/git-town/git-town/v20/internal/vm/shared"
)

// BranchEnsureShippableChanges asserts that the branch has unique changes not on the main branch.
type BranchEnsureShippableChanges struct {
	Branch                  gitdomain.LocalBranchName
	Parent                  gitdomain.LocalBranchName
	undeclaredOpcodeMethods `exhaustruct:"optional"`
}

func (self *BranchEnsureShippableChanges) AutomaticUndoError() error {
	return fmt.Errorf(messages.ShipBranchNothingToDo, self.Branch)
}

func (self *BranchEnsureShippableChanges) Run(args shared.RunArgs) error {
	hasUnmergedChanges, err := args.Git.BranchHasUnmergedChanges(args.Backend, self.Branch, self.Parent)
	if err != nil {
		return err
	}
	if !hasUnmergedChanges {
		return fmt.Errorf(messages.ShipBranchNothingToDo, self.Branch)
	}
	return nil
}

func (self *BranchEnsureShippableChanges) ShouldUndoOnError() bool {
	return true
}
