package opcodes

import (
	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	"github.com/git-town/git-town/v20/internal/vm/shared"
)

// CheckoutIfExists does the same as Checkout
// but only if that branch actually exists.
type CheckoutIfExists struct {
	Branch                  gitdomain.LocalBranchName
	undeclaredOpcodeMethods `exhaustruct:"optional"`
}

func (self *CheckoutIfExists) Run(args shared.RunArgs) error {
	existingBranch, err := args.Git.CurrentBranch(args.Backend)
	if err != nil {
		return err
	}
	if existingBranch == self.Branch {
		return nil
	}
	if !args.Git.BranchExists(args.Backend, self.Branch) {
		return nil
	}
	args.PrependOpcodes(&CheckoutIfNeeded{Branch: self.Branch})
	return nil
}
