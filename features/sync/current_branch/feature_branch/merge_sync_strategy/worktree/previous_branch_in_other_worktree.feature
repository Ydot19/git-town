Feature: sync a branch when the previous branch is active in another worktree

  Background:
    Given a Git repo with origin
    And the branches
      | NAME     | TYPE    | PARENT | LOCATIONS     |
      | current  | feature | main   | local, origin |
      | previous | feature | main   | local, origin |
    And the current branch is "current" and the previous branch is "previous"
    And branch "previous" is active in another worktree
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                 |
      | current | git fetch --prune --tags                |
      |         | git merge --no-edit --ff main           |
      |         | git merge --no-edit --ff origin/current |
    And no commits exist now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And no commits exist now
