Feature: sync the current branch that is observed via regex

  Background:
    Given a Git repo with origin
    And the branches
      | NAME       | TYPE   | LOCATIONS     |
      | renovate/1 | (none) | local, origin |
    And the commits
      | BRANCH     | LOCATION      | MESSAGE       | FILE NAME   |
      | main       | local, origin | main commit   | main_file   |
      | renovate/1 | local         | local commit  | local_file  |
      |            | origin        | origin commit | origin_file |
    And the current branch is "renovate/1"
    And Git setting "git-town.observed-regex" is "^renovate"
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH     | COMMAND                                                 |
      | renovate/1 | git fetch --prune --tags                                |
      |            | git -c rebase.updateRefs=false rebase origin/renovate/1 |
    And these commits exist now
      | BRANCH     | LOCATION      | MESSAGE       |
      | main       | local, origin | main commit   |
      | renovate/1 | local, origin | origin commit |
      |            | local         | local commit  |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH     | COMMAND                                              |
      | renovate/1 | git reset --hard {{ sha-before-run 'local commit' }} |
    And the initial commits exist now
    And the initial branches and lineage exist now
