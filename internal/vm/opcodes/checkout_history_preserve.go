package opcodes

import (
	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	"github.com/git-town/git-town/v20/internal/gohacks/slice"
	"github.com/git-town/git-town/v20/internal/vm/shared"
	. "github.com/git-town/git-town/v20/pkg/prelude"
)

// CheckoutHistoryPreserve does stuff.
type CheckoutHistoryPreserve struct {
	PreviousBranchCandidates []Option[gitdomain.LocalBranchName]
	undeclaredOpcodeMethods  `exhaustruct:"optional"`
}

func (self *CheckoutHistoryPreserve) Run(args shared.RunArgs) error {
	if !args.Git.CurrentBranchCache.Initialized() {
		// the branch cache is not initialized --> there were no branch changes --> no need to restore the branch history
		return nil
	}
	currentBranch := args.Git.CurrentBranchCache.Value()
	actualPreviousBranch := args.Git.CurrentBranchCache.Previous()
	// remove the current branch from the list of previous branch candidates because the current branch should never also be the previous branch
	candidates := slice.GetAll(self.PreviousBranchCandidates)
	candidatesWithoutCurrent := slice.Remove(candidates, currentBranch)
	expectedPreviousBranch, hasExpectedPreviousBranch := args.Git.FirstExistingBranch(args.Backend, candidatesWithoutCurrent...).Get()
	if !hasExpectedPreviousBranch || actualPreviousBranch == expectedPreviousBranch {
		return nil
	}
	// We	need to ignore errors here because failing to set the Git branch history
	// is not an error condition.
	// This operation can fail for a number of reasons like the previous branch being
	// checked out in another worktree, or concurrent Git access
	args.PrependOpcodes(
		&CheckoutUncached{Branch: expectedPreviousBranch},
		&CheckoutUncached{Branch: currentBranch},
	)
	return nil
}
