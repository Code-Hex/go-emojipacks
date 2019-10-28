PROJECT = github.com/Code-Hex/go-emojipacks

.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/emojipacks -trimpath -ldflags "-w -s" \
		$(PROJECT)/cmd/emojipacks