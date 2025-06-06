Feature: does not sync branches that exist only on remotes

  Background:
    Given a Git repo with origin
    And the branches
      | NAME  | TYPE    | PARENT | LOCATIONS     |
      | mine  | feature | main   | local, origin |
      | other | feature | main   | origin        |
    And the commits
      | BRANCH | LOCATION      | MESSAGE         |
      | main   | origin        | main commit     |
      | mine   | local, origin | my commit       |
      | other  | origin        | coworker commit |
    And the current branch is "main"
    When I run "git-town sync --all"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                                           |
      | main   | git fetch --prune --tags                          |
      |        | git -c rebase.updateRefs=false rebase origin/main |
      |        | git checkout mine                                 |
      | mine   | git merge --no-edit --ff main                     |
      |        | git merge --no-edit --ff origin/mine              |
      |        | git push                                          |
      |        | git checkout main                                 |
      | main   | git push --tags                                   |
    And all branches are now synchronized
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE                       |
      | main   | local, origin | main commit                   |
      | mine   | local, origin | my commit                     |
      |        |               | Merge branch 'main' into mine |
      | other  | origin        | coworker commit               |
