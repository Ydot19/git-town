Feature: ship a local feature branch using the fast-forward strategy

  Background:
    Given a Git repo with origin
    And the branches
      | NAME    | TYPE    | PARENT | LOCATIONS |
      | feature | feature | main   | local     |
    And the commits
      | BRANCH  | LOCATION | MESSAGE        |
      | feature | local    | feature commit |
    And the current branch is "feature"
    And Git setting "git-town.ship-strategy" is "fast-forward"
    When I run "git-town ship"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH  | COMMAND                     |
      | feature | git fetch --prune --tags    |
      |         | git checkout main           |
      | main    | git merge --ff-only feature |
      |         | git push                    |
      |         | git branch -D feature       |
    And the branches are now
      | REPOSITORY    | BRANCHES |
      | local, origin | main     |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE        |
      | main   | local, origin | feature commit |
    And no lineage exists now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                       |
      | main   | git branch feature {{ sha 'feature commit' }} |
      |        | git checkout feature                          |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE        |
      | main   | local, origin | feature commit |
    And the initial branches and lineage exist now
