Feature: handle merge conflicts between feature branch and main branch in a local repo

  Background:
    Given a local Git repo
    And the branches
      | NAME  | TYPE    | PARENT | LOCATIONS |
      | alpha | feature | main   | local     |
      | beta  | feature | main   | local     |
      | gamma | feature | main   | local     |
    And the commits
      | BRANCH | LOCATION | MESSAGE      | FILE NAME        | FILE CONTENT  |
      | main   | local    | main commit  | conflicting_file | main content  |
      | alpha  | local    | alpha commit | feature1_file    | alpha content |
      | beta   | local    | beta commit  | conflicting_file | beta content  |
      | gamma  | local    | gamma commit | feature3_file    | gamma content |
    And the current branch is "main"
    When I run "git-town sync --all"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                       |
      | main   | git checkout alpha            |
      | alpha  | git merge --no-edit --ff main |
      |        | git checkout beta             |
      | beta   | git merge --no-edit --ff main |
    And Git Town prints the error:
      """
      CONFLICT (add/add): Merge conflict in conflicting_file
      """
    And a merge is now in progress

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                   |
      | beta   | git merge --abort                         |
      |        | git checkout alpha                        |
      | alpha  | git reset --hard {{ sha 'alpha commit' }} |
      |        | git checkout main                         |
    And the initial commits exist now
    And no merge is in progress

  Scenario: skip
    When I run "git-town skip"
    Then Git Town runs the commands
      | BRANCH | COMMAND                       |
      | beta   | git merge --abort             |
      |        | git checkout gamma            |
      | gamma  | git merge --no-edit --ff main |
      |        | git checkout main             |
    And no merge is in progress
    And these commits exist now
      | BRANCH | LOCATION | MESSAGE                        |
      | main   | local    | main commit                    |
      | alpha  | local    | alpha commit                   |
      |        |          | Merge branch 'main' into alpha |
      | beta   | local    | beta commit                    |
      | gamma  | local    | gamma commit                   |
      |        |          | Merge branch 'main' into gamma |
    And these committed files exist now
      | BRANCH | NAME             | CONTENT       |
      | main   | conflicting_file | main content  |
      | alpha  | conflicting_file | main content  |
      |        | feature1_file    | alpha content |
      | beta   | conflicting_file | beta content  |
      | gamma  | conflicting_file | main content  |
      |        | feature3_file    | gamma content |

  Scenario: continue with unresolved conflict
    When I run "git-town continue"
    Then Git Town runs no commands
    And Git Town prints the error:
      """
      you must resolve the conflicts before continuing
      """
    And a merge is now in progress

  Scenario: resolve and continue
    When I resolve the conflict in "conflicting_file"
    And I run "git-town continue"
    Then Git Town runs the commands
      | BRANCH | COMMAND                       |
      | beta   | git commit --no-edit          |
      |        | git checkout gamma            |
      | gamma  | git merge --no-edit --ff main |
      |        | git checkout main             |
    And all branches are now synchronized
    And no merge is in progress
    And these committed files exist now
      | BRANCH | NAME             | CONTENT          |
      | main   | conflicting_file | main content     |
      | alpha  | conflicting_file | main content     |
      |        | feature1_file    | alpha content    |
      | beta   | conflicting_file | resolved content |
      | gamma  | conflicting_file | main content     |
      |        | feature3_file    | gamma content    |

  Scenario: resolve, commit, and continue
    When I resolve the conflict in "conflicting_file"
    And I run "git commit --no-edit"
    And I run "git-town continue"
    Then Git Town runs the commands
      | BRANCH | COMMAND                       |
      | beta   | git checkout gamma            |
      | gamma  | git merge --no-edit --ff main |
      |        | git checkout main             |
