Feature: syncing a branch whose parent was shipped

  Background:
    Given a Git repo with origin
    And the branches
      | NAME   | TYPE    | PARENT | LOCATIONS     |
      | parent | feature | main   | local, origin |
      | child  | feature | parent | local, origin |
    And the commits
      | BRANCH | LOCATION      | MESSAGE       |
      | parent | local, origin | parent commit |
      | child  | local, origin | child commit  |
    And origin ships the "parent" branch using the "squash-merge" ship-strategy
    And the current branch is "child"
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                                           |
      | child  | git fetch --prune --tags                          |
      |        | git checkout main                                 |
      | main   | git -c rebase.updateRefs=false rebase origin/main |
      |        | git branch -D parent                              |
      |        | git checkout child                                |
      | child  | git merge --no-edit --ff main                     |
      |        | git merge --no-edit --ff origin/child             |
      |        | git push                                          |
    And Git Town prints:
      """
      deleted branch "parent"
      """
    And the branches are now
      | REPOSITORY    | BRANCHES    |
      | local, origin | main, child |
    And this lineage exists now
      | BRANCH | PARENT |
      | child  | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                                |
      | child  | git reset --hard {{ sha 'child commit' }}              |
      |        | git push --force-with-lease --force-if-includes        |
      |        | git checkout main                                      |
      | main   | git reset --hard {{ sha 'initial commit' }}            |
      |        | git branch parent {{ sha-before-run 'parent commit' }} |
      |        | git checkout child                                     |
    And the initial branches and lineage exist now
