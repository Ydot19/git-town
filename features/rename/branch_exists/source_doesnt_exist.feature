Feature: branch does not exist

  Scenario:
    Given a Git repo with origin
    And the commits
      | BRANCH | LOCATION      | MESSAGE     |
      | main   | local, origin | main commit |
    And the current branch is "main"
    When I run "git-town rename non-existing new"
    Then Git Town runs the commands
      | BRANCH | COMMAND                  |
      | main   | git fetch --prune --tags |
    And Git Town prints the error:
      """
      there is no branch "non-existing"
      """
