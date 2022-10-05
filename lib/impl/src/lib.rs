use arrow::ffi::{FFI_ArrowArray, ArrowArray};
use arrow::array::{Int32Array, Array, make_array_from_raw};
use libc::c_int;

#[no_mangle]
pub extern "C" fn call_with_ffi(ffi_array: &mut ArrowArray) -> c_int {
    
    let (arr_ptr, sch_ptr) = ArrowArray::into_raw(unsafe { ArrowArray::empty() });
    let array = unsafe { make_array_from_raw(arr_ptr, sch_ptr) };

    println!("Hello from Rust, with Arrow Array: {:?}", array);

    1 as c_int
}

