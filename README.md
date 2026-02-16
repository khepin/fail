# fail

A simple CLI tool that prints a message and exits with code 1.

## Usage

```sh
fail <message>
```

## Example: use in scripts

```sh
#!/bin/bash

curl -s https://example.com/data > output.json || fail "failed to fetch data"
jq '.results' output.json > results.json || fail "failed to parse JSON"
upload results.json || fail "failed to upload results"
```

Each step either succeeds or halts the script with a clear error message and a non-zero exit code.

## Install

```sh
go install github.com/khepin/fail@latest
```
