# © 2022 Vlad-Stefan Harbuz <vlad@vladh.net>
# SPDX-License-Identifier: MIT

.PHONY: all upload

all:
	GOOS=linux GOARCH=arm tinygo build -o bin/dewdrop

upload: all
	scp ./bin/dewdrop dew-a:.local/bin/
	scp ./bin/dewdrop dew-b:.local/bin/
