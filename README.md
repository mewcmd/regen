# regen

[![Build Status](https://travis-ci.org/mewcmd/regen.svg?branch=master)](https://travis-ci.org/mewcmd/regen)
[![Coverage Status](https://img.shields.io/coveralls/mewcmd/regen.svg)](https://coveralls.io/r/mewcmd/regen?branch=master)
[![GoDoc](https://godoc.org/github.com/mewcmd/regen?status.svg)](https://godoc.org/github.com/mewcmd/regen)

The `regen` tool generates all strings from the regular expression of a finite language.

## Installation

	go get github.com/mewcmd/regen

## Usage

	regen REGEX

## Examples

1. Generate the new 64-bit general purpose registers from a regular expression.

		regen "r(8|9|1[0-5])(b|w|d)?"
		// Output:
		// r8
		// r8b
		// r8d
		// r8w
		// r9
		// r9b
		// r9d
		// r9w
		// ...
		// r15
		// r15b
		// r15d
		// r15w

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
