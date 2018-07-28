GO_BIN=GOPATH=`pwd` go
MODEL_SRC=`pwd`/cmd

model_xor:
	$(GO_BIN) run $(MODEL_SRC)/model_xor.go
