package configdomain_test

import (
	"testing"

	"github.com/git-town/git-town/v20/internal/config/configdomain"
	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	"github.com/shoenig/test/must"
)

func TestPerennialRegex(t *testing.T) {
	t.Parallel()

	t.Run("only characters, no wildcards matches all branch names that contain that phrase", func(t *testing.T) {
		t.Parallel()
		perennialRegexOpt, err := configdomain.ParsePerennialRegex("release")
		must.NoError(t, err)
		perennialRegex := perennialRegexOpt.GetOrPanic()
		tests := map[string]bool{
			"":                false,
			"release":         true,
			"release-1":       true,
			"another-release": true,
			"main":            false,
		}
		for give, want := range tests {
			have := perennialRegex.MatchesBranch(gitdomain.LocalBranchName(give))
			must.Eq(t, want, have)
		}
	})

	t.Run("with wildcards", func(t *testing.T) {
		t.Parallel()
		perennialRegexOpt, err := configdomain.ParsePerennialRegex("release-.*")
		must.NoError(t, err)
		perennialRegex := perennialRegexOpt.GetOrPanic()
		tests := map[string]bool{
			"":                false,
			"release":         false,
			"release-1":       true,
			"release-2":       true,
			"release-30":      true,
			"another-release": false,
			"main":            false,
		}
		for give, want := range tests {
			have := perennialRegex.MatchesBranch(gitdomain.LocalBranchName(give))
			must.Eq(t, want, have)
		}
	})
}
