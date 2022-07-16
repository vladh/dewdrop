.PHONY: all

all:
	GOOS=linux GOARCH=arm tinygo build -o bin/dewdrop

upload:
	scp ./bin/dewdrop dew:.local/bin
