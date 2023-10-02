build:
	BRANCH=`git rev-parse --abbrev-ref HEAD`; \
	HASH=`git rev-parse HEAD`; \
	VERSION=$(version); \
	echo "Current Branch: $$BRANCH"; \
	echo "Last Commit Hash: $$HASH"; \
	echo "Version: $$VERSION"; \
	go build -a --race -gcflags=all="-l -B -C" -ldflags="-w -s -X 'syncommit/cmd.VERSION=$$VERSION' -X 'syncommit/cmd.BRANCH_NAME=$$BRANCH' -X 'syncommit/cmd.HASH_NAME=$$HASH'"
