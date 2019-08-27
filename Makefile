GO ?= go
CASK ?= cask

all: go-emacs-go-base64.so

go-emacs-go-base64.so: main.go
	$(GO) build -buildmode=c-shared -o $@

clean:
	$(RM) go-emacs-go-base64.so

test: go-emacs-go-base64.so
	$(CASK) exec ert-runner
