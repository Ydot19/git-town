Feature: detached sync the current feature branch using the "rebase" feature sync strategy

  Background:
    Given a Git repo with origin
    And the branches
      | NAME    | TYPE    | PARENT | LOCATIONS     |
      | feature | feature | main   | local, origin |
    And the commits
      | BRANCH  | LOCATION | MESSAGE               |
      | main    | local    | local main commit     |
      |         | origin   | origin main commit    |
      | feature | local    | local feature commit  |
      |         | origin   | origin feature commit |
    And the current branch is "feature"
    And Git setting "git-town.sync-feature-strategy" is "rebase"
    When I run "git-town sync --detached"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                              |
      | feature | git fetch --prune --tags                             |
      |         | git -c rebase.updateRefs=false rebase main           |
      |         | git push --force-with-lease --force-if-includes      |
      |         | git -c rebase.updateRefs=false rebase origin/feature |
      |         | git -c rebase.updateRefs=false rebase main           |
      |         | git push --force-with-lease --force-if-includes      |
    And these commits exist now
      | BRANCH  | LOCATION      | MESSAGE               |
      | main    | local         | local main commit     |
      |         | origin        | origin main commit    |
      | feature | local, origin | origin feature commit |
      |         |               | local feature commit  |
      |         | origin        | local main commit     |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                                                                           |
      | feature | git reset --hard {{ sha-before-run 'local feature commit' }}                                      |
      |         | git push --force-with-lease origin {{ sha-in-origin-before-run 'origin feature commit' }}:feature |
    And the initial commits exist now
    And the initial branches and lineage exist now
