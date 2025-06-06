Feature: syncing a local feature branch using --no-push

  Background:
    Given a Git repo with origin
    And the commits
      | BRANCH | LOCATION | MESSAGE            |
      | main   | local    | local main commit  |
      |        | origin   | origin main commit |
    And the branches
      | NAME    | TYPE    | PARENT | LOCATIONS |
      | feature | feature | main   | local     |
    And the commits
      | BRANCH  | LOCATION | MESSAGE              |
      | feature | local    | local feature commit |
    And the current branch is "feature"
    And Git setting "git-town.sync-feature-strategy" is "rebase"
    When I run "git-town sync --no-push"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                           |
      | feature | git fetch --prune --tags                          |
      |         | git checkout main                                 |
      | main    | git -c rebase.updateRefs=false rebase origin/main |
      |         | git checkout feature                              |
      | feature | git -c rebase.updateRefs=false rebase main        |
    And these commits exist now
      | BRANCH  | LOCATION      | MESSAGE              |
      | main    | local, origin | origin main commit   |
      |         | local         | local main commit    |
      | feature | local         | local feature commit |
    And the initial branches and lineage exist now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                           |
      | feature | git reset --hard {{ sha 'local feature commit' }} |
      |         | git checkout main                                 |
      | main    | git reset --hard {{ sha 'local main commit' }}    |
      |         | git checkout feature                              |
    And the initial commits exist now
    And the initial branches and lineage exist now
