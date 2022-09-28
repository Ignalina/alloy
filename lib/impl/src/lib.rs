extern crate libc;
use arrow2::ffi::ArrowSchema;

#[no_mangle]
pub extern "C" fn callwithtable(data: &mut ArrowSchema)  {
    println!("Hello from Rust, with ArrowSchema: {:?}", data);
}
