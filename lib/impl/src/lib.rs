use arrow::ffi::{FFI_ArrowSchema, FFI_ArrowArray};
use arrow::array::make_array_from_raw;
use libc::c_int;

#[no_mangle]
pub extern "C" fn call_with_ffi(
    ffi_array: *const FFI_ArrowArray,
    ffi_schema: *const FFI_ArrowSchema
    ) -> c_int {

    unsafe {
        let _array_ref = match make_array_from_raw(ffi_array, ffi_schema) {
            Ok(a) => a,
            Err(arrow_error) => panic!("Could not make_array_from_raw, {:?}", arrow_error),
        };

        println!("Hello from Rust, with Arrow ArrayRef.");
    }

    1 as c_int
}

