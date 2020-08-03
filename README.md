# AC Deck

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

Download binary from [Release page](https://github.com/sachaos/ac-deck/releases)

```shell
$ curl -L -o ./ac-deck.tar.gz https://github.com/sachaos/ac-deck/releases/download/v0.3.4/ac-deck_0.3.4_Linux_x86_64.tar.gz && mkdir ./ac-deck-bin && tar xvzf ./ac-deck.tar.gz -C ./ac-deck-bin && sudo mv ./ac-deck-bin/acd /usr/local/bin/acd && sudo chmod +x /usr/local/bin/acd
```

## Mac OS X

```shell
$ brew install sachaos/tap/ac-deck
```

## Build it yourself

You need go 1.13.

```shell
$ git clone https://github.com/sachaos/ac-deck.git
$ cd ac-deck
$ make install
```

# Setup

## Configure authentication information

```shell
$ acd config
```

**WARNING**: This software store raw authentication information on `~/.ac-deck.toml`. This is not secure.
Please understand this behavior, and use carefully. Please contribute if you interest to fix this behavior.

## Install test runnder

```shell
$ acd install python3
```

# Usage

## Prepare to solve problems

```shell
$ acd prepare abc153
```

### Specify language to solve

```shell
$ acd prepare --language python3 abc153
```

## Browse problem (on web browser)

```shell
$ acd abc153/abc153_a browse
```

## Edit code

```shell
$ acd abc153/abc153_a edit
```

You can customize the editor by `$EDITOR` environment variable.

## Test

```shell
$ acd abc153/abc153_a test
```

## Submit if test passed

```shell
$ acd abc153/abc153_a submit
```

### Submit without test

```shell
$ acd abc153/abc153_a submit --skip-test
```

## Supporting Language

[AtCoder Languages and Compiler options](https://atcoder.jp/contests/language-test-202001)

- Python3 (3.8.2)
- Go (1.14.1)
    - You cannot use gonum, gods now.

### Old versions

[AtCoder Languages and Compiler options (Old)](https://language-test-201603.contest.atcoder.jp/)

- C++14 (GCC 5.4.1)
- C++ (GCC 5.4.1)
- C# (Mono 4.6.2.0)
- Go (1.6)
- Python2 (2.7.6)
    - You cannot use numpy, scipy, scikits now.
- Python3 (3.4.3)
- Ruby (2.3.3)
- C (GCC 5.4.1)
- Java7 (OpenJDK 1.7.0)
- Java8 (OpenJDK 1.8.0)
