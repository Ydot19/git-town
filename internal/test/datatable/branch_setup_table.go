package datatable

import (
	"github.com/cucumber/godog"
	"github.com/git-town/git-town/v20/internal/config/configdomain"
	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	"github.com/git-town/git-town/v20/internal/test/testgit"
	"github.com/git-town/git-town/v20/pkg/asserts"
	. "github.com/git-town/git-town/v20/pkg/prelude"
)

type BranchSetup struct {
	BranchType Option[configdomain.BranchType]
	Locations  testgit.Locations
	Name       gitdomain.LocalBranchName
	Parent     Option[gitdomain.LocalBranchName]
}

func ParseBranchSetupTable(table *godog.Table) []BranchSetup {
	result := make([]BranchSetup, 0, len(table.Rows)-1)
	headers := table.Rows[0]
	lastLocations := testgit.Locations{}
	for _, row := range table.Rows[1:] {
		name := None[gitdomain.LocalBranchName]()
		branchType := None[configdomain.BranchType]()
		parent := None[gitdomain.LocalBranchName]()
		locations := testgit.Locations{}
		for c, cell := range row.Cells {
			switch headers.Cells[c].Value {
			case "NAME":
				name = Some(gitdomain.NewLocalBranchName(cell.Value))
			case "TYPE":
				branchType = asserts.NoError1(configdomain.ParseBranchType(cell.Value))
			case "PARENT":
				if len(cell.Value) > 0 {
					parent = Some(gitdomain.NewLocalBranchName(cell.Value))
				}
			case "LOCATIONS":
				if cell.Value == "" {
					if len(lastLocations) > 0 {
						locations = lastLocations
					} else {
						panic("branch table does not provide locations")
					}
				} else {
					locations = testgit.NewLocations(cell.Value)
					lastLocations = locations
				}

			default:
				panic("unknown branch table header: " + cell.Value)
			}
		}
		if len(locations) == 0 {
			panic("branch table doesn't define locations")
		}
		result = append(result, BranchSetup{
			BranchType: branchType,
			Locations:  locations,
			Name:       name.GetOrPanic(),
			Parent:     parent,
		})
	}
	return result
}
