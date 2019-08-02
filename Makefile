MAIN_VERSION = $(shell cat ./VERSION)
BRANCH = $(shell git branch | grep \* | cut -d ' ' -f2)
COMMIT_HASH = $(shell git describe --always --long)
FULL_VERSION = ${MAIN_VERSION}-${BRANCH}-${COMMIT_HASH}

version:
	@echo ${FULL_VERSION}

.PHONY: version

