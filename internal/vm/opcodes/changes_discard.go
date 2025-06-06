package opcodes

import "github.com/git-town/git-town/v20/internal/vm/shared"

// ChangesDiscard resets the branch to the last commit, discarding uncommitted changes.
type ChangesDiscard struct {
	undeclaredOpcodeMethods `exhaustruct:"optional"`
}

func (self *ChangesDiscard) Run(args shared.RunArgs) error {
	return args.Git.DiscardOpenChanges(args.Frontend)
}
