LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

.PHONY: build
run: $(LOCALBIN)
	go build -o bin/tykstation && bin/tykstation
