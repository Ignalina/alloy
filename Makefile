build-all:
	cd ./ffi/ && cargo build --release
	cp ./ffi/target/release/librust_impl.a ./ffi/
	go build ./...

clean:
	rm -f ffi/libimpl.a && rm -rf ./ffi/target
	rm -f alloy
	go clean

