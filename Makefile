# vim: foldenable foldmethod=marker foldlevel=1
#
# {{{1 Configurations
PWD=`pwd`
GO_PATH=GOPATH=$(PWD)
GO=$(GO_PATH) go
SRC=$(PWD)/src
BIN=$(PWD)/bin

MODEL_SRC=$(PWD)/cmd

# {{{1 Actions
# {{{2 Default
default: clean fmt model_xor

# {{{2 model_xor
model_xor:
	$(GO) build -o $(BIN)/model_xor $(MODEL_SRC)/model_xor.go

# {{{1 Maintenance.
# {{{2 fmt
fmt:
	gofmt -w -l $(SRC)

# {{{2 clean
clean:
	rm -f $(BIN)/*

