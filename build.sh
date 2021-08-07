#!/bin/bash
# Basic build & install script for Mac/Linux
set -e
BINDIR="/usr/local/bin" # Path to install binaries within
# go mod init
go get github.com/howeyc/gopass # Lib for password input

cd ./wavhide
go build
mv wavhide $BINDIR

cd ../wavreveal
go build
mv wavreveal $BINDIR

echo "Packages built and installed in $BINDIR ($(which wavhide), $(which wavreveal))" || echo "Error building/installing packages :-("
