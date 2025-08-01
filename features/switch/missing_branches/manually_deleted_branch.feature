@messyoutput
Feature: switch branches while a manually deleted branch is still listed in the lineage

  Scenario: repo contains a manually deleted branch
    Given a Git repo with origin
    And the branches
      | NAME  | TYPE    | PARENT | LOCATIONS |
      | alpha | feature | main   | local     |
      | beta  | feature | main   | local     |
      | gamma | feature | main   | local     |
    And the current branch is "alpha"
    And I run "git branch -D beta"
    When I run "git-town switch" and enter into the dialogs:
      | DIALOG        | KEYS       |
      | switch-branch | down enter |
    Then Git Town runs the commands
      | BRANCH | COMMAND            |
      | alpha  | git checkout gamma |
