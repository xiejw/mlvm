PACKAGES=github.com/xiejw/mlvm/go/...

compile:
	go build ${PACKAGES}

fmt:
	go fmt ${PACKAGES}

test:
	go test ${PACKAGES}

bench:
	go test -bench=. ${PACKAGES}
