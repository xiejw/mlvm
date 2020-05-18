PACKAGES=github.com/xiejw/mlvm/go/...

compile:
	go build ${PACKAGES}

fmt:
	go fmt ${PACKAGES}
