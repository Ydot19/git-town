package gitdomain_test

import (
	"testing"

	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	"github.com/shoenig/test/must"
)

func TestSHAs(t *testing.T) {
	t.Parallel()

	t.Run("First", func(t *testing.T) {
		t.Parallel()
		shas := gitdomain.NewSHAs("111111", "222222")
		have := shas.First()
		want := shas[0]
		must.Eq(t, want, have)
	})

	t.Run("Join", func(t *testing.T) {
		t.Parallel()
		t.Run("contains elements", func(t *testing.T) {
			t.Parallel()
			give := gitdomain.SHAs{"111111", "222222"}
			have := give.Join(", ")
			want := "111111, 222222"
			must.EqOp(t, want, have)
		})
		t.Run("empty list", func(t *testing.T) {
			t.Parallel()
			give := gitdomain.SHAs{}
			have := give.Join(", ")
			want := ""
			must.EqOp(t, want, have)
		})
	})

	t.Run("Last", func(t *testing.T) {
		t.Parallel()
		t.Run("multiple elements", func(t *testing.T) {
			t.Parallel()
			shas := gitdomain.NewSHAs("111111", "222222")
			have := shas.Last()
			want := shas[1]
			must.Eq(t, want, have)
		})
		t.Run("one element", func(t *testing.T) {
			t.Parallel()
			shas := gitdomain.NewSHAs("111111")
			have := shas.Last()
			want := shas[0]
			must.Eq(t, want, have)
		})
	})

	t.Run("NewSHAs", func(t *testing.T) {
		t.Parallel()
		have := gitdomain.NewSHAs("111111", "222222", "333333")
		want := gitdomain.SHAs{"111111", "222222", "333333"}
		must.Eq(t, want, have)
	})

	t.Run("Strings", func(t *testing.T) {
		t.Parallel()
		t.Run("contains elements", func(t *testing.T) {
			t.Parallel()
			give := gitdomain.SHAs{"111111", "222222"}
			have := give.Strings()
			want := []string{"111111", "222222"}
			must.Eq(t, want, have)
		})
		t.Run("empty list", func(t *testing.T) {
			t.Parallel()
			give := gitdomain.SHAs{}
			have := give.Strings()
			want := []string{}
			must.Eq(t, want, have)
		})
	})
}
