package opcodes

import (
	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	"github.com/git-town/git-town/v20/internal/vm/shared"
)

// CreateBranch creates a new branch but leaves the current branch unchanged.
type BranchLocalRename struct {
	NewName                 gitdomain.LocalBranchName
	OldName                 gitdomain.LocalBranchName
	undeclaredOpcodeMethods `exhaustruct:"optional"`
}

func (self *BranchLocalRename) Run(args shared.RunArgs) error {
	return args.Git.RenameBranch(args.Frontend, self.OldName, self.NewName)
}
