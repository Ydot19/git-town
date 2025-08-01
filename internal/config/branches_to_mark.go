package config

import (
	"errors"

	"github.com/git-town/git-town/v21/internal/config/configdomain"
	"github.com/git-town/git-town/v21/internal/git/gitdomain"
	"github.com/git-town/git-town/v21/internal/messages"
	. "github.com/git-town/git-town/v21/pkg/prelude"
)

// provides the branches to make contribution, observed, parked, or prototype
func BranchesToMark(args []string, branchesSnapshot gitdomain.BranchesSnapshot, config UnvalidatedConfig) (branchesToMark configdomain.BranchesAndTypes, branchToCheckout Option[gitdomain.LocalBranchName], err error) {
	branchesToMark = configdomain.BranchesAndTypes{}
	switch len(args) {
	case 0:
		currentBranch, hasCurrentBranch := branchesSnapshot.Active.Get()
		if !hasCurrentBranch {
			return branchesToMark, branchToCheckout, errors.New(messages.CurrentBranchCannotDetermine)
		}
		branchesToMark.Add(currentBranch, &config)
		branchToCheckout = None[gitdomain.LocalBranchName]()
	case 1:
		branch := gitdomain.NewLocalBranchName(args[0])
		branchesToMark.Add(branch, &config)
		branchInfo, hasBranchInfo := branchesSnapshot.Branches.FindByRemoteName(branch.TrackingBranch(config.NormalConfig.DevRemote)).Get()
		if hasBranchInfo && branchInfo.SyncStatus == gitdomain.SyncStatusRemoteOnly {
			branchToCheckout = Some(branch)
		} else {
			branchToCheckout = None[gitdomain.LocalBranchName]()
		}
	default:
		branchesToMark.AddMany(gitdomain.NewLocalBranchNames(args...), &config)
		branchToCheckout = None[gitdomain.LocalBranchName]()
	}
	return branchesToMark, branchToCheckout, nil
}
