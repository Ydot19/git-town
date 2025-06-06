Feature: handle merge conflicts between feature branch and main branch

  Background:
    Given a Git repo with origin
    And the branches
      | NAME  | TYPE    | PARENT | LOCATIONS     |
      | alpha | feature | main   | local, origin |
      | beta  | feature | main   | local, origin |
      | gamma | feature | main   | local, origin |
    And the commits
      | BRANCH | LOCATION      | MESSAGE            | FILE NAME            | FILE CONTENT        |
      | main   | origin        | main commit        | conflicting_file     | main content        |
      | alpha  | local, origin | alpha commit       | feature1_file        | alpha content       |
      | beta   | local         | local beta commit  | conflicting_file     | local beta content  |
      |        | origin        | origin beta commit | feature2_origin_file | origin beta content |
      | gamma  | origin        | gamma commit       | feature3_file        | gamma content       |
    And the current branch is "main"
    When I run "git-town sync --all"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                                           |
      | main   | git fetch --prune --tags                          |
      |        | git -c rebase.updateRefs=false rebase origin/main |
      |        | git checkout alpha                                |
      | alpha  | git merge --no-edit --ff main                     |
      |        | git merge --no-edit --ff origin/alpha             |
      |        | git push                                          |
      |        | git checkout beta                                 |
      | beta   | git merge --no-edit --ff main                     |
    And Git Town prints the error:
      """
      CONFLICT (add/add): Merge conflict in conflicting_file
      """
    And a merge is now in progress

  Scenario: abort
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                         |
      | beta   | git merge --abort                               |
      |        | git checkout alpha                              |
      | alpha  | git reset --hard {{ sha 'alpha commit' }}       |
      |        | git push --force-with-lease --force-if-includes |
      |        | git checkout main                               |
      | main   | git reset --hard {{ sha 'initial commit' }}     |
    And no merge is in progress
    And the initial commits exist now
    And the initial branches and lineage exist now

  Scenario: skip
    When I run "git-town skip"
    Then Git Town runs the commands
      | BRANCH | COMMAND                               |
      | beta   | git merge --abort                     |
      |        | git checkout gamma                    |
      | gamma  | git merge --no-edit --ff main         |
      |        | git merge --no-edit --ff origin/gamma |
      |        | git push                              |
      |        | git checkout main                     |
      | main   | git push --tags                       |
    And no merge is in progress
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE                                                |
      | main   | local, origin | main commit                                            |
      | alpha  | local, origin | alpha commit                                           |
      |        |               | Merge branch 'main' into alpha                         |
      | beta   | local         | local beta commit                                      |
      |        | origin        | origin beta commit                                     |
      | gamma  | local, origin | gamma commit                                           |
      |        |               | Merge remote-tracking branch 'origin/gamma' into gamma |
    And these committed files exist now
      | BRANCH | NAME             | CONTENT            |
      | main   | conflicting_file | main content       |
      | alpha  | conflicting_file | main content       |
      |        | feature1_file    | alpha content      |
      | beta   | conflicting_file | local beta content |
      | gamma  | conflicting_file | main content       |
      |        | feature3_file    | gamma content      |

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
      | BRANCH | COMMAND                               |
      | beta   | git commit --no-edit                  |
      |        | git merge --no-edit --ff origin/beta  |
      |        | git push                              |
      |        | git checkout gamma                    |
      | gamma  | git merge --no-edit --ff main         |
      |        | git merge --no-edit --ff origin/gamma |
      |        | git push                              |
      |        | git checkout main                     |
      | main   | git push --tags                       |
    And all branches are now synchronized
    And no merge is in progress
    And these committed files exist now
      | BRANCH | NAME                 | CONTENT             |
      | main   | conflicting_file     | main content        |
      | alpha  | conflicting_file     | main content        |
      |        | feature1_file        | alpha content       |
      | beta   | conflicting_file     | resolved content    |
      |        | feature2_origin_file | origin beta content |
      | gamma  | conflicting_file     | main content        |
      |        | feature3_file        | gamma content       |

  Scenario: resolve, commit, and continue
    When I resolve the conflict in "conflicting_file"
    And I run "git commit --no-edit"
    And I run "git-town continue"
    Then Git Town runs the commands
      | BRANCH | COMMAND                               |
      | beta   | git merge --no-edit --ff origin/beta  |
      |        | git push                              |
      |        | git checkout gamma                    |
      | gamma  | git merge --no-edit --ff main         |
      |        | git merge --no-edit --ff origin/gamma |
      |        | git push                              |
      |        | git checkout main                     |
      | main   | git push --tags                       |
