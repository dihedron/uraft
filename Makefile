
.PHONY: binary
binary:
	@go build -o uraft

.PHONY: test
test: binary
	cd tests/raft && goreman -f bootstrap.procfile start && cd -
#	./uraft --state=tests/raft/store node1
#	./uraft --peer=@tests/raft/node2.json --peer=@tests/raft/node3.json --state=tests/raft/store node1