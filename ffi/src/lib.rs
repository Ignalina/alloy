use libc::c_int;
use arrow::ffi::{FFI_ArrowArray, FFI_ArrowSchema};
use arrow::array::{ArrayRef, make_array_from_raw};


#[no_mangle]
pub extern "C" fn call_with_ffi(
    ffi_array: *const FFI_ArrowArray,
    ffi_schema: *const FFI_ArrowSchema
    ) -> c_int {

    println!("Hello2 from Rust!");

    let arr = unsafe { match make_array_from_raw(ffi_array, ffi_schema) {
        Ok(a) => a,
        Err(e) => panic!("Could not make_array_from_raw, {:?}", e),
    } };

    println!("Hello3 from Rust, again!\nArrow::Array = {:?}", arr);

    1 as c_int
}

