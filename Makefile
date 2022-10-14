build-all:
	cd ./ffi/ && cargo build --release
	cp ./ffi/target/release/librust_impl.a ./ffi/
	go build ./...

clean:
	rm -f ./ffi/librust_impl.a
	rm -rf ./ffi/target
	rm -f ./ffi/Cargo.lock
	rm -f alloy
	go clean

