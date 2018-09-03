# vim: foldenable foldmethod=marker foldlevel=1
#
# {{{1 Configurations
PWD=`pwd`
GO_PATH=GOPATH=$(PWD)
GO=$(GO_PATH) go
GO_DOC=$(GO_PATH) godoc
SRC=$(PWD)/src
BIN=$(PWD)/bin
TOOLS=$(PWD)/tools
TOOLS_BIN=$(TOOLS)/bin

MODEL_SRC=$(PWD)/cmd

# {{{1 Actions
# {{{2 Default
default: clean fmt generate model_xor

generate:
	go build -o $(TOOLS_BIN)/immutable_slice_gen $(TOOLS)/immutable_slice_gen.go
	$(TOOLS_BIN)/immutable_slice_gen > $(SRC)/mlvm/modules/layers/inputs_generated.go

# {{{2 model_xor
model_xor:
	$(GO) build -o $(BIN)/model_xor $(MODEL_SRC)/model_xor.go

run_model_xor: model_xor
	$(BIN)/model_xor && dot -Tpng /tmp/123.dot > /tmp/123.png

# {{{1 Maintenance.
# {{{2 fmt
fmt:
	gofmt -w -l $(SRC)
	gofmt -w -l $(TOOLS)

# {{{2 clean
clean:
	rm -f $(BIN)/*

# {{{2 doc
.PHONY: doc
doc:
	@echo "**** Open http://localhost:6060/pkg/mlvm/ ****\n"
	$(GO_DOC) -v --http=:6060

# {{{2 dependency tree
dependency_tree:
	${GO_PATH} goimportdot -pkg=mlvm > /tmp/pkg.dot && \
		dot -Tsvg /tmp/pkg.dot > /tmp/pkg.svg && \
		echo "Open /tmp/pkg.svg with your viewer."

