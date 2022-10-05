build-all:
	cd lib/impl && cargo build --release
	cp lib/impl/target/release/libimpl.a lib/
	go build ./...

clean:
	rm lib/libimpl.a || rm -rf lib/impl/target
	go clean

