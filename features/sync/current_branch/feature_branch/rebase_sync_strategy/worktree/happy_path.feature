Feature: sync a branch whose parent is active in another worktree

  Background:
    Given a Git repo with origin
    And the branches
      | NAME   | TYPE    | PARENT | LOCATIONS     |
      | parent | feature | main   | local, origin |
      | child  | feature | parent | local, origin |
    And the commits
      | BRANCH | LOCATION | MESSAGE              |
      | main   | local    | local main commit    |
      |        | origin   | origin main commit   |
      | parent | local    | local parent commit  |
      |        | origin   | origin parent commit |
      | child  | local    | local child commit   |
      |        | origin   | origin child commit  |
    And the current branch is "child"
    And branch "parent" is active in another worktree
    And Git setting "git-town.sync-feature-strategy" is "rebase"
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                                             |
      | child  | git fetch --prune --tags                            |
      |        | git checkout main                                   |
      | main   | git -c rebase.updateRefs=false rebase origin/main   |
      |        | git push                                            |
      |        | git checkout child                                  |
      | child  | git -c rebase.updateRefs=false rebase origin/parent |
      |        | git push --force-with-lease --force-if-includes     |
      |        | git -c rebase.updateRefs=false rebase origin/child  |
      |        | git -c rebase.updateRefs=false rebase origin/parent |
      |        | git push --force-with-lease --force-if-includes     |
    And these commits exist now
      | BRANCH | LOCATION                | MESSAGE              |
      | main   | local, origin, worktree | origin main commit   |
      |        |                         | local main commit    |
      | child  | local                   | origin parent commit |
      |        | local, origin           | origin child commit  |
      |        |                         | local child commit   |
      | parent | origin                  | origin parent commit |
      |        | worktree                | local parent commit  |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                                                                       |
      | child  | git reset --hard {{ sha-before-run 'local child commit' }}                                    |
      |        | git push --force-with-lease origin {{ sha-in-origin-before-run 'origin child commit' }}:child |
