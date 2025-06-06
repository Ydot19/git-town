package slice_test

import (
	"testing"

	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	"github.com/git-town/git-town/v20/internal/gohacks/slice"
	"github.com/shoenig/test/must"
)

func TestAppendAllMissing(t *testing.T) {
	t.Parallel()

	t.Run("aliased slice type", func(t *testing.T) {
		t.Parallel()
		list := gitdomain.SHAs{"111111", "222222"}
		give := gitdomain.SHAs{"333333", "444444"}
		have := slice.AppendAllMissing(list, give...)
		want := gitdomain.SHAs{"111111", "222222", "333333", "444444"}
		must.Eq(t, want, have)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		list := make([]string, 0)
		have := slice.AppendAllMissing(list, "one", "two", "three")
		want := []string{"one", "two", "three"}
		must.Eq(t, want, have)
	})

	t.Run("slice type", func(t *testing.T) {
		t.Parallel()
		list := []string{"one", "two", "three"}
		have := slice.AppendAllMissing(list, "two", "four", "five")
		want := []string{"one", "two", "three", "four", "five"}
		must.Eq(t, want, have)
	})

	t.Run("zero slice", func(t *testing.T) {
		t.Parallel()
		var list []string
		have := slice.AppendAllMissing(list, "one", "two", "three")
		want := []string{"one", "two", "three"}
		must.Eq(t, want, have)
	})
}
