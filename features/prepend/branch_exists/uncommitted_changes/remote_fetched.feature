Feature: already existing remote branch

  Background:
    Given a Git repo with origin
    And the branches
      | NAME     | TYPE    | PARENT | LOCATIONS     |
      | old      | feature | main   | local, origin |
      | existing | feature | main   | origin        |
    And the current branch is "old"
    And an uncommitted file
    And I run "git fetch"
    When I run "git-town prepend existing"

  Scenario: result
    Then Git Town runs no commands
    And Git Town prints the error:
      """
      there is already a branch "existing"
      """
    And the uncommitted file still exists
    And the initial commits exist now
    And the initial branches and lineage exist now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And the uncommitted file still exists
    And the initial commits exist now
    And the initial lineage exists now
