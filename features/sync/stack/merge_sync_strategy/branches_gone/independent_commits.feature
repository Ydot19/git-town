Feature: stacked changes

  Background:
    Given a Git repo with origin
    And the branches
      | NAME   | TYPE    | PARENT | LOCATIONS     |
      | parent | feature | main   | local, origin |
      | child  | feature | parent | local, origin |
    And the commits
      | BRANCH | LOCATION | MESSAGE              |
      | main   | local    | local main commit    |
      |        | origin   | origin main commit   |
      | parent | local    | local parent commit  |
      |        | origin   | origin parent commit |
      | child  | local    | local child commit   |
      |        | origin   | origin child commit  |
    And the current branch is "child"
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                                           |
      | child  | git fetch --prune --tags                          |
      |        | git checkout main                                 |
      | main   | git -c rebase.updateRefs=false rebase origin/main |
      |        | git push                                          |
      |        | git checkout parent                               |
      | parent | git merge --no-edit --ff main                     |
      |        | git merge --no-edit --ff origin/parent            |
      |        | git push                                          |
      |        | git checkout child                                |
      | child  | git merge --no-edit --ff parent                   |
      |        | git merge --no-edit --ff origin/child             |
      |        | git push                                          |
    And all branches are now synchronized
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE                                                  |
      | main   | local, origin | origin main commit                                       |
      |        |               | local main commit                                        |
      | child  | local, origin | local child commit                                       |
      |        |               | Merge branch 'parent' into child                         |
      |        |               | origin child commit                                      |
      |        |               | Merge remote-tracking branch 'origin/child' into child   |
      | parent | local, origin | local parent commit                                      |
      |        |               | Merge branch 'main' into parent                          |
      |        |               | origin parent commit                                     |
      |        |               | Merge remote-tracking branch 'origin/parent' into parent |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                                                                         |
      | child  | git reset --hard {{ sha-before-run 'local child commit' }}                                      |
      |        | git push --force-with-lease origin {{ sha-in-origin-before-run 'origin child commit' }}:child   |
      |        | git checkout parent                                                                             |
      | parent | git reset --hard {{ sha-before-run 'local parent commit' }}                                     |
      |        | git push --force-with-lease origin {{ sha-in-origin-before-run 'origin parent commit' }}:parent |
      |        | git checkout child                                                                              |
