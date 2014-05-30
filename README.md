Isana : a software for the game of Go
=====================================

## Overview

Isana is an open source project to develop a software for the game of Go.
It is currently uses the UCT engine with Monte-Carlo aproach

*still weak!*

### GTP Support

* protocol_version
* name
* version
* list_commands
* komi
* play
* genmove
* quit
* known_command
* boardsize
* clear_board
* showboard

## Build / Install

    go get github.com/foozea/isana
    go install github.com/foozea/isana

## Requirements

* The [Go](http://golang.org) Programming Language.

## License

Isana is free, and distributed under the **GNU General Public License**
(GPL). Essentially, this means that you are free to do almost exactly
what you want with the program, including distributing it among your
friends, making it available for download from your web site, selling
it (either by itself or as part of some bigger software package), or
using it as the starting point for a software project of your own.

The only real limitation is that whenever you distribute Stockfish in
some way, you must always include the full source code, or a pointer
to where the source code can be found. If you make any changes to the
source code, these changes must also be made available under the GPL.

For full details, read the copy of the GPL found in the file named
*LICENSE*
