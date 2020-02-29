# atcoder-cli

Unofficial CLI for AtCoder users.

# Demo

![demo](./images/demo.gif)

## Features

* Template (built-in)
* Run test on Docker or Native
* Submit code to AtCoder
* Support major languages (If you want to add another languages, please send PR.)

# Install

## Linux

Download binary from [Release page](https://github.com/sachaos/atcoder/releases)

```shell
$ curl -L -o ./atcoder.tar.gz https://github.com/sachaos/atcoder/releases/download/v0.2.9/atcoder_0.2.9_Linux_x86_64.tar.gz && mkdir ./atcoder-bin && tar xvzf ./atcoder.tar.gz -C ./atcoder-bin && sudo mv ./atcoder-bin/atcoder /usr/local/bin/atcoder && sudo chmod +x /usr/local/bin/atcoder
```

## Mac OS X

```shell
$ brew install sachaos/atcoder/atcoder
```

## Build it yourself

You need go 1.13.

```shell
$ git clone https://github.com/sachaos/atcoder.git
$ cd atcoder
$ make install
```

# Setup

## Configure authentication information

```shell
$ atcoder config
```

**WARNING**: This software store raw authentication information on `~/.atcoder.toml`. This is not secure.
Please understand this behavior, and use carefully. Please contribute if you interest to fix this behavior.

## Install test runnder

```shell
$ atcoder install python3
```

# Usage

## Prepare to solve problems

```shell
$ atcoder prepare abc153
```

### Specify language to solve

```shell
$ atcoder prepare --language python3 abc153
```

## Browse problem (on web browser)

```shell
$ atcoder abc153/abc153_a browse
```

## Edit code

```shell
$ atcoder abc153/abc153_a edit
```

You can customize the editor by `$EDITOR` environment variable.

## Test

```shell
$ atcoder abc153/abc153_a test
```

## Submit if test passed

```shell
$ atcoder abc153/abc153_a submit
```

### Submit without test

```shell
$ atcoder abc153/abc153_a submit --skip-test
```

## Supporting Language

[AtCoder Languages and Compiler options](https://language-test-201603.contest.atcoder.jp/)

- C++14 (GCC 5.4.1)
- C++ (GCC 5.4.1)
- Go (1.6)
- Python2 (2.7.6)
    - You cannot use numpy, scipy, scikits now.
- Python3 (3.4.3)
- Ruby (2.3.3)
- C (GCC 5.4.1)
- Java7 (OpenJDK 1.7.0)
- Java8 (OpenJDK 1.8.0)
