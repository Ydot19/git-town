@messyoutput
Feature: ask for missing parent information

  Scenario:
    Given a Git repo with origin
    And the branches
      | NAME | TYPE   | LOCATIONS |
      | old  | (none) | local     |
    And the current branch is "old"
    When I run "git-town prepend new" and enter into the dialog:
      | DIALOG                  | KEYS  |
      | parent branch for "old" | enter |
    Then Git Town runs the commands
      | BRANCH | COMMAND                  |
      | old    | git fetch --prune --tags |
      |        | git push -u origin old   |
      |        | git checkout -b new main |
    And this lineage exists now
      | BRANCH | PARENT |
      | new    | main   |
      | old    | new    |
