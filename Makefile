.PHONY: all

all:
	GOOS=linux GOARCH=arm tinygo build -o bin/dewdrop


