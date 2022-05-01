TARGET := Panopeia

GO      = ${GOROOT}/bin/go

#--------------------------------------------------------------------------------
LINUX:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${TARGET}

tool:
	$(GO) fmt ./
	$(GO) vet ./

BUILD:
	@echo "begin build" && export GO111MODULE=on
	$(GO) test -gcflags "all=-N -l" ./... | grep "FAIL" | wc -l
	$(GO) build $(GO_BUILD_FLAG) -o ${TARGET}

test:
	@echo "begin go test..."
	$(GO) test -gcflags "all=-N -l" ./... -covermode=count

clean:
	@if [ -f ${TARGET} ]; then rm ${TARGET}; fi

run:BUILD
	./${TARGET}
