build-all:
	cd ./ffi/rust && cargo build --release
	cp ./ffi/rust/target/release/librust_impl.a ./ffi/rust/
	go build

clean:
	rm -f ./ffi/rust/librust_impl.a
	rm -rf ./ffi/rust/target
	rm -f ./ffi/rust/Cargo.lock
	rm -f alloy
	go clean

