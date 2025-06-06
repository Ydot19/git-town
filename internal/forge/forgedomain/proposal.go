package forgedomain

import "github.com/git-town/git-town/v20/internal/git/gitdomain"

// Proposal contains information about a change request on a forge.
// Alternative names are "pull request" or "merge request".
type Proposal struct {
	// whether this proposal can be merged via the API
	MergeWithAPI bool

	// the number used to identify the proposal on the forge
	Number int

	// name of the source branch ("head") of this proposal
	Source gitdomain.LocalBranchName

	// name of the target branch ("base") of this proposal
	Target gitdomain.LocalBranchName

	// textual title of the proposal
	Title string

	// the URL of this proposal
	URL string
}
