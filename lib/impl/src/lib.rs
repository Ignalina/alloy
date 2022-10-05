use libc::c_int;
use arrow::ffi::{FFI_ArrowSchema, FFI_ArrowArray};
use arrow::array::{ArrayRef, make_array_from_raw};
use arrow::error::{ArrowError};

#[no_mangle]
pub extern "C" fn call_with_ffi(
    ffi_array: *const FFI_ArrowArray,
    ffi_schema: *const FFI_ArrowSchema
    ) -> c_int {

    unsafe {
        let array_ref = match make_array_from_raw(ffi_array, ffi_schema) {
            Ok(a) => a,
            Err(ArrowError) => panic!("Could not make_array_from_raw, {:?}", ArrowError),
        };

        println!("Hello from Rust, with Arrow ArrayRef.");
    }

    1 as c_int
}

