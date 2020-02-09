# atcoder-cli

Unofficial CLI for AtCoder users.

# Setup

```shell
$ atcoder config
```

**WARNING**: This software store raw authentication information on `~/.atcoder.conf`. This is not secure.
Please understand this behavior, and use carefully. Please contribute if you interest to fix this behavior.

# Usage

## Prepare to solve problems

```shell
$ atcoder prepare abc153
```

### Specify language to solve

```shell
$ atcoder prepare --language python abc153
```

## Test

```shell
$ atcoder test abc153/abc153_a
```

### Submit if example passed (not implemented yet)

```shell
$ atcoder test --submit abc153/abc153_a
```

## Submit (not implemented yet)

```shell
$ atcoder submit abc153/abc153_a
```
