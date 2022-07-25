# © 2022 Vlad-Stefan Harbuz <vlad@vladh.net>
# SPDX-License-Identifier: MIT

.PHONY: all

all:
	GOOS=linux GOARCH=arm tinygo build -o bin/dewdrop

upload:
	scp ./bin/dewdrop dew:.local/bin
