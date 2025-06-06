package testgit_test

import (
	"testing"

	"github.com/git-town/git-town/v20/internal/test/testgit"
	"github.com/shoenig/test/must"
)

func TestLocations(t *testing.T) {
	t.Parallel()

	t.Run("Contains", func(t *testing.T) {
		t.Parallel()
		t.Run("has the element", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationLocal, testgit.LocationOrigin}
			have := locations.Contains(testgit.LocationOrigin)
			must.True(t, have)
		})
		t.Run("does not have the element", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationLocal}
			have := locations.Contains(testgit.LocationOrigin)
			must.False(t, have)
		})
	})

	t.Run("Is", func(t *testing.T) {
		t.Parallel()
		t.Run("match with one element", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationOrigin}
			must.True(t, locations.Is(testgit.LocationOrigin))
		})
		t.Run("match with multiple elements", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationLocal, testgit.LocationOrigin}
			must.True(t, locations.Is(testgit.LocationLocal, testgit.LocationOrigin))
		})
		t.Run("wrong type", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationOrigin}
			must.False(t, locations.Is(testgit.LocationLocal))
		})
		t.Run("contains more elements", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationLocal, testgit.LocationOrigin}
			must.False(t, locations.Is(testgit.LocationLocal))
		})
		t.Run("contains fewer elements", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationLocal}
			must.False(t, locations.Is(testgit.LocationLocal, testgit.LocationOrigin))
		})
	})

	t.Run("Matches", func(t *testing.T) {
		t.Parallel()
		t.Run("has exactly the given elements", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationLocal, testgit.LocationOrigin}
			have := locations.Matches(testgit.LocationOrigin, testgit.LocationLocal)
			must.True(t, have)
		})
		t.Run("has the given elements and more", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationLocal, testgit.LocationOrigin, testgit.LocationCoworker}
			have := locations.Matches(testgit.LocationOrigin, testgit.LocationLocal)
			must.False(t, have)
		})
		t.Run("has not all of the given elements", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationLocal}
			have := locations.Matches(testgit.LocationOrigin, testgit.LocationLocal)
			must.False(t, have)
		})
		t.Run("has other elements", func(t *testing.T) {
			t.Parallel()
			locations := testgit.Locations{testgit.LocationLocal}
			have := locations.Matches(testgit.LocationOrigin)
			must.False(t, have)
		})
	})

	t.Run("NewLocations", func(t *testing.T) {
		t.Parallel()
		tests := map[string]testgit.Locations{
			"local":                   {testgit.LocationLocal},
			"origin":                  {testgit.LocationOrigin},
			"upstream":                {testgit.LocationUpstream},
			"local, origin":           {testgit.LocationLocal, testgit.LocationOrigin},
			"local, upstream":         {testgit.LocationLocal, testgit.LocationUpstream},
			"local, origin, upstream": {testgit.LocationLocal, testgit.LocationOrigin, testgit.LocationUpstream},
		}
		for give, want := range tests {
			have := testgit.NewLocations(give)
			must.Eq(t, want, have)
		}
	})
}
