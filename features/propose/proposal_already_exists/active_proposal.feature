@skipWindows
Feature: open the page of an already existing proposal

  Background:
    Given a Git repo with origin
    And the branches
      | NAME    | TYPE    | PARENT | LOCATIONS     |
      | feature | feature | main   | local, origin |
    And tool "open" is installed
    And the current branch is "feature"
    And the origin is "git@github.com:git-town/git-town.git"
    And a proposal for this branch exists at "https://github.com/git-town/git-town/pull/123"

  Scenario: a PR for this branch exists already
    When I run "git-town propose"
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                            |
      | feature | git fetch --prune --tags                           |
      | (none)  | Looking for proposal online ... ok                 |
      | feature | git merge --no-edit --ff main                      |
      |         | git merge --no-edit --ff origin/feature            |
      | (none)  | open https://github.com/git-town/git-town/pull/123 |
    And the initial branches and lineage exist now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And the initial commits exist now
