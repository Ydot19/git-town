Feature: on the main branch with uncommitted changes

  Background:
    Given a Git repo with origin
    And the commits
      | BRANCH | LOCATION | MESSAGE     |
      | main   | origin   | main commit |
    And the current branch is "main"
    And an uncommitted file
    When I run "git-town hack new"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                     |
      | main   | git add -A                  |
      |        | git stash -m "Git Town WIP" |
      |        | git checkout -b new         |
      | new    | git stash pop               |
      |        | git restore --staged .      |
    And the uncommitted file still exists
    And the initial commits exist now
    And this lineage exists now
      | BRANCH | PARENT |
      | new    | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                     |
      | new    | git add -A                  |
      |        | git stash -m "Git Town WIP" |
      |        | git checkout main           |
      | main   | git branch -D new           |
      |        | git stash pop               |
      |        | git restore --staged .      |
    And the uncommitted file still exists
    And the initial commits exist now
    And the initial branches and lineage exist now
