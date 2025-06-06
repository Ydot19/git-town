package bitbucketcloud_test

import (
	"testing"

	"github.com/git-town/git-town/v20/internal/forge/bitbucketcloud"
	"github.com/git-town/git-town/v20/internal/git/giturl"
	"github.com/shoenig/test/must"
)

func TestDetect(t *testing.T) {
	t.Parallel()
	tests := map[string]bool{
		"username@bitbucket.org:git-town/docs.git": true,  // SAAS URL
		"git@custom-url.com:git-town/docs.git":     false, // custom URL
		"git@github.com:git-town/git-town.git":     false, // other hosting service URL
	}
	for give, want := range tests {
		url, has := giturl.Parse(give).Get()
		must.True(t, has)
		have := bitbucketcloud.Detect(url)
		must.EqOp(t, want, have)
	}
}
