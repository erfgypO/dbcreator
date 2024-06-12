#!/bin/zsh

go build -o dbcreator cmd/main.go
cp dbcreator ~/.scripts/
rm dbcreator