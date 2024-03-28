.POSIX:
.SUFFIXES:

all: test lint

test:
	go test -race -shuffle=on -cover ./...

lint:
	golangci-lint run

tidy:
	go mod tidy

readme:
	python3 scripts/generate_readme.py

# run `make pre-commit` once to install the hook.
pre-commit: .git/hooks/pre-commit test lint tidy readme
	git diff --exit-code

.git/hooks/pre-commit:
	echo "make pre-commit" > .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
