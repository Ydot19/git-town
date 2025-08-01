@messyoutput
Feature: make a feature branch perennial

  Background:
    Given a Git repo with origin
    And the branches
      | NAME   | TYPE    | PARENT | LOCATIONS     |
      | parent | feature | main   | local, origin |
      | child  | feature | parent | local, origin |
    And the current branch is "child"
    When I run "git-town set-parent" and enter into the dialog:
      | DIALOG                    | KEYS       |
      | parent branch for "child" | down enter |

  Scenario: result
    Then Git Town prints:
      """
      branch "child" is now perennial
      """
    And Git Town runs no commands
    And the perennial branches are now "child"
    And this lineage exists now
      | BRANCH | PARENT |
      | parent | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And the initial commits exist now
    And the initial branches and lineage exist now
