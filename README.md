# atcoder-cli

Unofficial CLI for AtCoder users.

# Install

## Download binary

### Linux

```shell
$ sudo curl -L -o /usr/local/bin/atcoder https://github.com/sachaos/atcoder/releases/download/v0.1.0/atcoder_linux_amd64 && sudo chmod +x /usr/local/bin/atcoder
```

### Mac OS X

```shell
$ sudo curl -L -o /usr/local/bin/atcoder https://github.com/sachaos/atcoder/releases/download/v0.1.0/atcoder_darwin_amd64 && sudo chmod +x /usr/local/bin/atcoder
```

## Build it yourself

You need go 1.13.

```shell
$ git clone https://github.com/sachaos/atcoder.git
$ cd atcoder
$ make install
```

# Setup

```shell
$ atcoder config
```

**WARNING**: This software store raw authentication information on `~/.atcoder.toml`. This is not secure.
Please understand this behavior, and use carefully. Please contribute if you interest to fix this behavior.

# Usage

## Prepare to solve problems

```shell
$ atcoder prepare abc153
```

### Specify language to solve

```shell
$ atcoder prepare --language python3 abc153
```

## Test

```shell
$ atcoder test abc153/abc153_a
```

## Submit if test passed

```shell
$ atcoder submit abc153/abc153_a
```

### Submit without test

```shell
$ atcoder submit abc153/abc153_a --skip-test
```
