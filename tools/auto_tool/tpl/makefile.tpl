
all:proto

.PHONY: proto, boot

proto:
	@./build_proto.sh

boot:
	@./build_boot.sh