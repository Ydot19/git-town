package configdomain

import (
	"slices"

	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	. "github.com/git-town/git-town/v20/pkg/prelude"
	"github.com/git-town/git-town/v20/pkg/set"
)

// configuration settings that exist in both UnvalidatedConfig and ValidatedConfig
type NormalConfigData struct {
	Aliases                  Aliases
	BitbucketAppPassword     Option[BitbucketAppPassword]
	BitbucketUsername        Option[BitbucketUsername]
	BranchTypeOverrides      BranchTypeOverrides
	CodebergToken            Option[CodebergToken]
	ContributionRegex        Option[ContributionRegex]
	DefaultBranchType        BranchType
	DevRemote                gitdomain.Remote
	FeatureRegex             Option[FeatureRegex]
	ForgeType                Option[ForgeType] // Some = override by user, None = auto-detect
	GitHubToken              Option[GitHubToken]
	GitLabToken              Option[GitLabToken]
	GiteaToken               Option[GiteaToken]
	HostingOriginHostname    Option[HostingOriginHostname]
	Lineage                  Lineage
	NewBranchType            Option[BranchType]
	ObservedRegex            Option[ObservedRegex]
	Offline                  Offline
	PerennialBranches        gitdomain.LocalBranchNames
	PerennialRegex           Option[PerennialRegex]
	PushHook                 PushHook
	ShareNewBranches         ShareNewBranches
	ShipDeleteTrackingBranch ShipDeleteTrackingBranch
	ShipStrategy             ShipStrategy
	SyncFeatureStrategy      SyncFeatureStrategy
	SyncPerennialStrategy    SyncPerennialStrategy
	SyncPrototypeStrategy    SyncPrototypeStrategy
	SyncTags                 SyncTags
	SyncUpstream             SyncUpstream
}

func (self *NormalConfigData) IsOnline() bool {
	return self.Online().IsTrue()
}

func (self *NormalConfigData) NoPushHook() NoPushHook {
	return self.PushHook.Negate()
}

func (self *NormalConfigData) Online() Online {
	return self.Offline.ToOnline()
}

func (self *NormalConfigData) PartialBranchType(branch gitdomain.LocalBranchName) BranchType {
	// check the branch type overrides
	if branchTypeOverride, hasBranchTypeOverride := self.BranchTypeOverrides[branch]; hasBranchTypeOverride {
		return branchTypeOverride
	}
	// check the configured branch lists
	if slices.Contains(self.PerennialBranches, branch) {
		return BranchTypePerennialBranch
	}
	// check if a regex matches
	if regex, has := self.ContributionRegex.Get(); has && regex.MatchesBranch(branch) {
		return BranchTypeContributionBranch
	}
	if regex, has := self.FeatureRegex.Get(); has && regex.MatchesBranch(branch) {
		return BranchTypeFeatureBranch
	}
	if regex, has := self.ObservedRegex.Get(); has && regex.MatchesBranch(branch) {
		return BranchTypeObservedBranch
	}
	if regex, has := self.PerennialRegex.Get(); has && regex.MatchesBranch(branch) {
		return BranchTypePerennialBranch
	}
	// branch doesn't match any of the overrides --> default branch type
	return self.DefaultBranchType
}

func (self *NormalConfigData) PartialBranchesOfType(branchType BranchType) gitdomain.LocalBranchNames {
	matching := set.New[gitdomain.LocalBranchName]()
	switch branchType {
	case BranchTypeContributionBranch:
	case BranchTypeFeatureBranch:
	case BranchTypeMainBranch:
		// main branch is stored in ValidatedConfig
	case BranchTypeObservedBranch:
	case BranchTypeParkedBranch:
	case BranchTypePerennialBranch:
		matching.Add(self.PerennialBranches...)
	case BranchTypePrototypeBranch:
	}
	for key, value := range self.BranchTypeOverrides {
		if value == branchType {
			matching.Add(key)
		}
	}
	return matching.Values()
}

func DefaultNormalConfig() NormalConfigData {
	return NormalConfigData{
		Aliases:                  Aliases{},
		BitbucketAppPassword:     None[BitbucketAppPassword](),
		BitbucketUsername:        None[BitbucketUsername](),
		BranchTypeOverrides:      BranchTypeOverrides{},
		CodebergToken:            None[CodebergToken](),
		ContributionRegex:        None[ContributionRegex](),
		DefaultBranchType:        BranchTypeFeatureBranch,
		DevRemote:                gitdomain.RemoteOrigin,
		FeatureRegex:             None[FeatureRegex](),
		ForgeType:                None[ForgeType](),
		GitHubToken:              None[GitHubToken](),
		GitLabToken:              None[GitLabToken](),
		GiteaToken:               None[GiteaToken](),
		HostingOriginHostname:    None[HostingOriginHostname](),
		Lineage:                  NewLineage(),
		NewBranchType:            None[BranchType](),
		ObservedRegex:            None[ObservedRegex](),
		Offline:                  false,
		PerennialBranches:        gitdomain.LocalBranchNames{},
		PerennialRegex:           None[PerennialRegex](),
		PushHook:                 true,
		ShareNewBranches:         ShareNewBranchesNone,
		ShipDeleteTrackingBranch: true,
		ShipStrategy:             ShipStrategyAPI,
		SyncFeatureStrategy:      SyncFeatureStrategyMerge,
		SyncPerennialStrategy:    SyncPerennialStrategyRebase,
		SyncPrototypeStrategy:    SyncPrototypeStrategyRebase,
		SyncTags:                 true,
		SyncUpstream:             true,
	}
}
