Feature: swapping a branch that is deleted at the remote

  Background:
    Given a Git repo with origin
    And the branches
      | NAME     | TYPE    | PARENT   | LOCATIONS     |
      | branch-1 | feature | main     | local, origin |
      | branch-2 | feature | branch-1 | local, origin |
      | branch-3 | feature | branch-2 | local, origin |
    And the current branch is "branch-2"
    And origin deletes the "branch-2" branch
    When I run "git-town swap"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH   | COMMAND                  |
      | branch-2 | git fetch --prune --tags |
    And Git Town prints the error:
      """
      please sync your branches before swapping
      """
    And the initial commits exist now
    And the initial lineage exists now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And the initial commits exist now
    And the initial lineage exists now
