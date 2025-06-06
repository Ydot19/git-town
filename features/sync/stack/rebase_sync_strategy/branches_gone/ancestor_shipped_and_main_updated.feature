Feature: shipped the head branch of a synced stack with dependent changes while main also received independent updates

  Background:
    Given a Git repo with origin
    And the branches
      | NAME     | TYPE    | PARENT | LOCATIONS     |
      | branch-1 | feature | main   | local, origin |
    And the commits
      | BRANCH   | LOCATION      | MESSAGE  | FILE NAME | FILE CONTENT |
      | branch-1 | local, origin | commit 1 | file      | content 1    |
    And the branches
      | NAME     | TYPE    | PARENT   | LOCATIONS     |
      | branch-2 | feature | branch-1 | local, origin |
    And the commits
      | BRANCH   | LOCATION      | MESSAGE  | FILE NAME | FILE CONTENT |
      | branch-2 | local, origin | commit 2 | file      | content 2    |
    And the branches
      | NAME     | TYPE    | PARENT   | LOCATIONS     |
      | branch-3 | feature | branch-2 | local, origin |
    And the commits
      | BRANCH   | LOCATION      | MESSAGE  | FILE NAME | FILE CONTENT |
      | branch-3 | local, origin | commit 3 | file      | content 3    |
    And the branches
      | NAME     | TYPE    | PARENT   | LOCATIONS     |
      | branch-4 | feature | branch-3 | local, origin |
    And the commits
      | BRANCH   | LOCATION      | MESSAGE  | FILE NAME | FILE CONTENT |
      | branch-4 | local, origin | commit 4 | file      | content 4    |
    And Git setting "git-town.sync-feature-strategy" is "rebase"
    And origin ships the "branch-1" branch using the "squash-merge" ship-strategy
    And origin ships the "branch-2" branch using the "squash-merge" ship-strategy and resolves the merge conflict in "file" with "content 2" and commits as "commit 2"
    And I add commit "additional commit" to the "main" branch
    And the current branch is "branch-4"
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH   | COMMAND                                                    |
      | branch-4 | git fetch --prune --tags                                   |
      |          | git checkout main                                          |
      | main     | git -c rebase.updateRefs=false rebase origin/main          |
      |          | git push                                                   |
      |          | git checkout branch-2                                      |
      | branch-2 | git -c rebase.updateRefs=false rebase --onto main branch-1 |
      |          | git checkout branch-3                                      |
      | branch-3 | git pull                                                   |
      |          | git -c rebase.updateRefs=false rebase --onto main branch-2 |
      |          | git push --force-with-lease                                |
      |          | git checkout branch-4                                      |
      | branch-4 | git pull                                                   |
      |          | git -c rebase.updateRefs=false rebase --onto main branch-2 |
      |          | git push --force-with-lease                                |
      |          | git -c rebase.updateRefs=false rebase branch-3             |
      |          | git push --force-with-lease --force-if-includes            |
      |          | git -c rebase.updateRefs=false rebase branch-3             |
      |          | git branch -D branch-1                                     |
      |          | git branch -D branch-2                                     |
    And all branches are now synchronized
    And these commits exist now
      | BRANCH   | LOCATION      | MESSAGE           | FILE NAME | FILE CONTENT |
      | main     | local, origin | commit 1          | file      | content 1    |
      |          |               | commit 2          | file      | content 2    |
      |          |               | additional commit | new_file  |              |
      | branch-3 | local, origin | commit 3          | file      | content 3    |
      | branch-4 | local, origin | commit 4          | file      | content 4    |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH   | COMMAND                                             |
      | branch-4 | git checkout branch-3                               |
      | branch-3 | git reset --hard {{ sha-before-run 'commit 3' }}    |
      |          | git push --force-with-lease --force-if-includes     |
      |          | git checkout branch-4                               |
      | branch-4 | git reset --hard {{ sha-before-run 'commit 4' }}    |
      |          | git push --force-with-lease --force-if-includes     |
      |          | git branch branch-1 {{ sha-before-run 'commit 1' }} |
      |          | git branch branch-2 {{ sha-before-run 'commit 2' }} |
    And these commits exist now
      | BRANCH   | LOCATION      | MESSAGE           | FILE NAME | FILE CONTENT |
      | main     | local, origin | commit 1          | file      | content 1    |
      |          |               | commit 2          | file      | content 2    |
      |          |               | additional commit | new_file  |              |
      | branch-1 | local         | commit 1          | file      | content 1    |
      | branch-2 | local         | commit 2          | file      | content 2    |
      | branch-3 | local, origin | commit 3          | file      | content 3    |
      |          | origin        | commit 1          | file      | content 1    |
      |          |               | commit 2          | file      | content 2    |
      | branch-4 | local, origin | commit 4          | file      | content 4    |
    And the initial branches and lineage exist now
