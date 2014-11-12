regen
=====

The `regen` tool generates all strings from the regular expression of a finite language.


Installation
------------

	go get github.com/mewcmd/regen

Usage
-----

	regen REGEX

Examples
--------

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

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
