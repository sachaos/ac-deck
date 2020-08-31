我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

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

- C++ (GCC 9.2.1)
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
