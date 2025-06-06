Feature: delete another than the current branch

  Background:
    Given a Git repo with origin
    And the branches
      | NAME | TYPE    | PARENT | LOCATIONS     |
      | good | feature | main   | local, origin |
      | dead | feature | main   | local, origin |
    And the commits
      | BRANCH | LOCATION      | MESSAGE            | FILE NAME        | FILE CONTENT |
      | main   | local, origin | conflicting commit | conflicting_file | main content |
      | dead   | local, origin | dead-end commit    | file             | dead content |
      | good   | local, origin | good commit        | file             | good content |
    And the current branch is "good"
    And an uncommitted file with name "conflicting_file" and content "conflicting content"
    When I run "git-town delete dead"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                     |
      | good   | git fetch --prune --tags    |
      |        | git add -A                  |
      |        | git stash -m "Git Town WIP" |
      |        | git push origin :dead       |
      |        | git branch -D dead          |
      |        | git stash pop               |
      |        | git restore --staged .      |
    And the uncommitted file has content:
      """
      conflicting content
      """
    And the branches are now
      | REPOSITORY    | BRANCHES   |
      | local, origin | main, good |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE            |
      | main   | local, origin | conflicting commit |
      | good   | local, origin | good commit        |
    And this lineage exists now
      | BRANCH | PARENT |
      | good   | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                     |
      | good   | git add -A                                  |
      |        | git stash -m "Git Town WIP"                 |
      |        | git branch dead {{ sha 'dead-end commit' }} |
      |        | git push -u origin dead                     |
      |        | git stash pop                               |
      |        | git restore --staged .                      |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE            |
      | main   | local, origin | conflicting commit |
      | dead   | local, origin | dead-end commit    |
      | good   | local, origin | good commit        |
    And the initial branches and lineage exist now
    And the uncommitted file has content:
      """
      conflicting content
      """
