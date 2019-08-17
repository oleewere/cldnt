# cldnt

[![GoDoc Widget](https://godoc.org/github.com/oleewere/cldnt?status.svg)](https://godoc.org/github.com/oleewere/cldnt/client)
[![Build Status](https://travis-ci.org/oleewere/cldnt.svg?branch=master)](https://travis-ci.org/oleewere/cldnt)
[![Go Report Card](https://goreportcard.com/badge/github.com/oleewere/cldnt)](https://goreportcard.com/report/github.com/oleewere/cldnt)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## Requirements

- Internet connection
- Go 1.12.x (for development)

## Installation 

### Installation on Mac OSX
```bash
brew tap oleewere/repo
brew install cldnt
```

### Installation on Linux

Using wget:
```bash
CLDNT_VERSION=0.5.0
wget -qO- "https://github.com/oleewere/cldnt/releases/download/v${CLDNT_VERSION}/cldnt_${CLDNT_VERSION}_linux_64-bit.tar.gz" | tar -C /usr/bin -zxv cldntl
```

Using curl:
```bash
CLDNT_VERSION=0.5.0
curl -L -s "https://github.com/oleewere/cldnt/releases/download/v${CLDNT_VERSION}/cldnt_${CLDNT_VERSION}_linux_64-bit.tar.gz" | tar -C /usr/bin -xzv cldnt
```

## Build

Build locally: 

```bash
make build
```

Or build in docker image:

```bash
docker build -t oleewere/cldnt:latest .
```

## Usage

```bash
# if it is installed
cldnt --help
# if it is built locally
./cldnt --help
```

Run in docker container (if it has already built)

``bash
docker run --rm oleewere/cldnt:latest --help
```

Currently only airports command is supported:

```bash
# see available parameters
cldnt help airports
# using it:
cldnt airports
```

