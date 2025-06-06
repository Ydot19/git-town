Feature: shipped the head branch of a synced stack with dependent changes that create a file while main also creates the same file

  Background:
    Given a Git repo with origin
    And the branches
      | NAME  | TYPE    | PARENT | LOCATIONS     |
      | alpha | feature | main   | local, origin |
    And the commits
      | BRANCH | LOCATION      | MESSAGE      | FILE NAME | FILE CONTENT  |
      | alpha  | local, origin | alpha commit | file      | alpha content |
    And the branches
      | NAME | TYPE    | PARENT | LOCATIONS     |
      | beta | feature | alpha  | local, origin |
    And the commits
      | BRANCH | LOCATION      | MESSAGE     | FILE NAME | FILE CONTENT |
      | beta   | local, origin | beta commit | file      | beta content |
    And Git setting "git-town.sync-feature-strategy" is "merge"
    And origin ships the "alpha" branch using the "squash-merge" ship-strategy
    And I add this commit to the "main" branch
      | MESSAGE                    | FILE NAME | FILE CONTENT   |
      | independent commit on main | file      | main content 1 |
    And the current branch is "beta"
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                                           |
      | beta   | git fetch --prune --tags                          |
      |        | git checkout main                                 |
      | main   | git -c rebase.updateRefs=false rebase origin/main |
    And Git Town prints the error:
      """
      CONFLICT (add/add): Merge conflict in file
      """
    And a rebase is now in progress

  Scenario: resolve and continue
    When I resolve the conflict in "file" with "resolved main content"
    And I run "git-town continue" and close the editor
    Then Git Town runs the commands
      | BRANCH | COMMAND                                   |
      | main   | git -c core.editor=true rebase --continue |
      |        | git push                                  |
      |        | git branch -D alpha                       |
      |        | git checkout beta                         |
      | beta   | git merge --no-edit --ff main             |
    And Git Town prints the error:
      """
      CONFLICT (add/add): Merge conflict in file
      """
    And a merge is now in progress
    When I resolve the conflict in "file" with "resolved beta content"
    And I run "git-town continue" and close the editor
    Then Git Town runs the commands
      | BRANCH | COMMAND                              |
      | beta   | git commit --no-edit                 |
      |        | git merge --no-edit --ff origin/beta |
      |        | git push                             |
    And all branches are now synchronized
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE                       | FILE NAME | FILE CONTENT          |
      | main   | local, origin | alpha commit                  | file      | alpha content         |
      |        |               | independent commit on main    | file      | resolved main content |
      | beta   | local, origin | alpha commit                  | file      | alpha content         |
      |        |               | beta commit                   | file      | beta content          |
      |        |               | Merge branch 'main' into beta | file      | resolved beta content |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND            |
      | main   | git rebase --abort |
      |        | git checkout beta  |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE                    | FILE NAME | FILE CONTENT   |
      | main   | local         | independent commit on main | file      | main content 1 |
      |        | origin        | alpha commit               | file      | alpha content  |
      | alpha  | local         | alpha commit               | file      | alpha content  |
      | beta   | local, origin | beta commit                | file      | beta content   |
      |        | origin        | alpha commit               | file      | alpha content  |
    And the initial branches and lineage exist now
