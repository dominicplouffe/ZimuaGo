# ZimuaGo
GOLang Chess Engine

A simple Chess engine written in Golang with some UCI support.  Tested with xBoard on Mac.

## Build
go build -gcflags "-s -w" -o engine *.go

## Execute
./engine -uci
