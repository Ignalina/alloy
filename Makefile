
build-all: build-rust build-go


.PHONY: build-rust
build-rust:
	@cd lib/impl && cargo build --release
	@cp lib/impl/target/release/libimpl.a lib/

.PHONY: build-go
build-go:
	go build alloy


