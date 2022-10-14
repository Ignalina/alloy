use libc::c_int;
use arrow2::ffi::ArrowSchema;

#[no_mangle]
pub extern "C" fn call_with_ffi_schema(
    ffi_schema: &ArrowSchema
) -> c_int {

    println!("Hello from Rust! Here is the schema: {:?}", ffi_schema); 
    1 as c_int
}
