package configdomain

import (
	"github.com/git-town/git-town/v21/internal/git/gitdomain"
	. "github.com/git-town/git-town/v21/pkg/prelude"
)

// UnvalidatedConfigData is the Git Town configuration as read from disk.
// It might be lacking essential information in case Git metadata and config files don't contain it.
// If you need this information, validate it into a ValidatedConfig.
type UnvalidatedConfigData struct {
	GitUserEmail Option[gitdomain.GitUserEmail]
	GitUserName  Option[gitdomain.GitUserName]
	MainBranch   Option[gitdomain.LocalBranchName]
}

// IsMainBranch indicates whether the branch with the given name
// is the main branch of the repository.
func (self *UnvalidatedConfigData) IsMainBranch(branch gitdomain.LocalBranchName) bool {
	if mainBranch, hasMainBranch := self.MainBranch.Get(); hasMainBranch {
		return branch == mainBranch
	}
	return false
}

// indicates the branch type of the given branch, if it can determine it
func (self *UnvalidatedConfigData) PartialBranchType(branch gitdomain.LocalBranchName) Option[BranchType] {
	if self.IsMainBranch(branch) {
		return Some(BranchTypeMainBranch)
	}
	return None[BranchType]()
}

func (self *UnvalidatedConfigData) ToValidatedConfig(defaults ValidatedConfigData) ValidatedConfigData {
	return ValidatedConfigData{
		GitUserEmail: self.GitUserEmail.GetOrElse(defaults.GitUserEmail),
		GitUserName:  self.GitUserName.GetOrElse(defaults.GitUserName),
		MainBranch:   self.MainBranch.GetOrElse(defaults.MainBranch),
	}
}
