package debug

import (
	"os"

	"github.com/git-town/git-town/v20/internal/cli/dialog"
	"github.com/git-town/git-town/v20/internal/cli/dialog/components"
	"github.com/git-town/git-town/v20/internal/config/configdomain"
	. "github.com/git-town/git-town/v20/pkg/prelude"
	"github.com/spf13/cobra"
)

func enterGitHubToken() *cobra.Command {
	return &cobra.Command{
		Use: "github-token",
		RunE: func(_ *cobra.Command, _ []string) error {
			dialogInputs := components.LoadTestInputs(os.Environ())
			_, _, err := dialog.GitHubToken(None[configdomain.GitHubToken](), dialogInputs.Next())
			return err
		},
	}
}
