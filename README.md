# qsctl

[![Build Status](https://travis-ci.com/qingstor/qsctl.svg?branch=master)](https://travis-ci.com/qingstor/qsctl)
[![GoDoc](https://godoc.org/github.com/qingstor/qsctl?status.svg)](https://godoc.org/github.com/qingstor/qsctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/qingstor/qsctl)](https://goreportcard.com/report/github.com/qingstor/qsctl)
[![codecov](https://codecov.io/gh/qingstor/qsctl/branch/master/graph/badge.svg)](https://codecov.io/gh/qingstor/qsctl)
[![License](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/qingstor/qsctl/blob/master/LICENSE)

[中文文档](./docs/README-zh_CN.md)

qsctl is intended to be an advanced command line tool for QingStor, it provides
powerful unix-like commands to let you manage QingStor resources just like files
on local machine.

## Installation

### Binary

Visit <https://github.com/qingstor/qsctl/releases> to get latest releases.

## Getting Started

### Configure via initialization wizard

### Configure manually

To use qsctl, there must be a configuration file , for example

```yaml
access_key_id: 'ACCESS_KEY_ID_EXAMPLE'
secret_access_key: 'SECRET_ACCESS_KEY_EXAMPLE'
```

The configuration file is `~/.qingstor/config.yaml` by default, it also
can be specified by the option `-c /path/to/config`.

You can also config other option like `host` , `port` and so on, just
add lines below into configuration file, for example

```yaml
host: 'qingstor.com'
port: 443
protocol: 'https'
connection_retries: 3
# Valid levels are 'debug', 'info', 'warn', 'error', and 'fatal'.
log_level: 'debug'
zone: 'zone_name'
```

## Available Commands

Commands supported by qsctl are listed below:

- `cat`: Cat a remote object into stdout.
- `cp`: Copy local file(s) to QingStor or QingStor key(s) to local.
- `ls`: List buckets, or objects with given prefix.
- `mb`: Make a new bucket.
- `presign`: Get the pre-signed URL by given object key.
- `rb`: Delete a bucket.
- `rm`: Remove remote object(s).
- `stat`: Stat a remote object.
- `sync`: Sync between local directory and QS-Directory.
- `tee`: Tee from stdin to a remote object.

See the detailed usage and examples with `qsctl help` or `qsctl <command> help`.
