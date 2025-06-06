Feature: shipped the head branch of a synced stack with dependent changes while main also received independent updates

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
    And I add commit "additional commit" to the "main" branch
    And the current branch is "beta"
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                                           |
      | beta   | git fetch --prune --tags                          |
      |        | git checkout main                                 |
      | main   | git -c rebase.updateRefs=false rebase origin/main |
      |        | git push                                          |
      |        | git branch -D alpha                               |
      |        | git checkout beta                                 |
      | beta   | git merge --no-edit --ff main                     |
      |        | git checkout --ours file                          |
      |        | git add file                                      |
      |        | git commit --no-edit                              |
      |        | git merge --no-edit --ff origin/beta              |
      |        | git push                                          |
    And all branches are now synchronized
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE                       | FILE NAME | FILE CONTENT  |
      | main   | local, origin | alpha commit                  | file      | alpha content |
      |        |               | additional commit             | new_file  |               |
      | beta   | local, origin | alpha commit                  | file      | alpha content |
      |        |               | beta commit                   | file      | beta content  |
      |        |               | Merge branch 'main' into beta |           |               |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                              |
      | beta   | git reset --hard {{ sha-before-run 'beta commit' }}  |
      |        | git push --force-with-lease --force-if-includes      |
      |        | git branch alpha {{ sha-before-run 'alpha commit' }} |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE           | FILE NAME | FILE CONTENT  |
      | main   | local, origin | alpha commit      | file      | alpha content |
      |        |               | additional commit | new_file  |               |
      | alpha  | local         | alpha commit      | file      | alpha content |
      | beta   | local, origin | beta commit       | file      | beta content  |
      |        | origin        | alpha commit      | file      | alpha content |
    And the initial branches and lineage exist now
