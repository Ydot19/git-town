Feature: cannot ship a feature branch that is active in another worktree

  Background:
    Given a Git repo with origin
    And the branches
      | NAME    | TYPE    | PARENT | LOCATIONS     |
      | feature | feature | main   | local, origin |
      | other   | feature | main   | local, origin |
    And the commits
      | BRANCH  | LOCATION      | MESSAGE        | FILE NAME        |
      | feature | local, origin | feature commit | conflicting_file |
    And the current branch is "other"
    And branch "feature" is active in another worktree
    And Git setting "git-town.ship-strategy" is "always-merge"
    When I run "git-town ship feature"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                  |
      | other  | git fetch --prune --tags |
    And Git Town prints the error:
      """
      branch "feature" is active in another worktree
      """

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And Git Town prints:
      """
      nothing to undo
      """
