Feature: compress a branch when the previous branch is active in another worktree

  Background:
    Given a Git repo with origin
    And the branches
      | NAME     | TYPE    | PARENT | LOCATIONS     |
      | current  | feature | main   | local, origin |
      | previous | feature | main   | local, origin |
    And the commits
      | BRANCH  | LOCATION      | MESSAGE  |
      | current | local, origin | commit 1 |
      |         |               | commit 2 |
    And the current branch is "current" and the previous branch is "previous"
    And branch "previous" is active in another worktree
    When I run "git-town compress"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                         |
      | current | git fetch --prune --tags                        |
      |         | git reset --soft main                           |
      |         | git commit -m "commit 1"                        |
      |         | git push --force-with-lease --force-if-includes |
    And the previous Git branch is now "current"
    And the initial branches and lineage exist now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                         |
      | current | git reset --hard {{ sha 'commit 2' }}           |
      |         | git push --force-with-lease --force-if-includes |
    And the previous Git branch is now "current"
    And the initial branches and lineage exist now
