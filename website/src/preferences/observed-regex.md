# Observed regex

Branches matching this regular expression are treated as
[observed branches](../branch-types.md#observed-branches).

## configure in config file

Setting the observed regex in the [config file](../configuration-file.md) is
only useful when the matching branches should be considered observed by all team
members. This is typically the case for branches generated by external services,
like Renovate or Dependabot.

```toml
[branches]
observed-regex = "^renovate/"
```

## configure in Git metadata

To manually set the feature regex, run this command:

```wrap
git config [--global] git-town.observed-regex '^renovate/'
```

The optional `--global` flag applies this setting to all Git repositories on
your local machine. When not present, the setting applies to the current repo.
