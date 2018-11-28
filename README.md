# json-util

[![GoDoc](https://godoc.org/github.com/scottgreenup/json-util?status.svg)](http://godoc.org/github.com/scottgreenup/json-util)
[![Build Status](https://travis-ci.org/scottgreenup/json-util.svg?branch=master)](https://travis-ci.org/scottgreenup/json-util)

## Example Usage

```
$ ju find -k ID example.json
.glossary.GlossDiv.GlossList.GlossEntry.ID
```

## Help

```
A utility to handle JSON.

Usage:
  ju [command]

Available Commands:
  find        Find the JSON path to bits of the JSON blob
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.ju.yaml)
  -h, --help            help for ju
  -t, --toggle          Help message for toggle

Use "ju [command] --help" for more information about a command.
```
