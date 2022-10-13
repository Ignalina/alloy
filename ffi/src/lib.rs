//use libc::c_int;
//use arrow::ffi::{FFI_ArrowArray, FFI_ArrowSchema};
//use arrow::array::{ArrayRef, make_array_from_raw};

use arrow2::{array::Array, datatypes::DataType, error::Result,ffi};
use libc::c_int;


#[no_mangle]
pub unsafe extern "C" fn call_with_ffi(
    array_ptr:  *const ffi::ArrowArray,
    schema_ptr: *const ffi::ArrowSchema
    ) -> c_int {

    println!("Hello2 from Rust!");


    let mut cs: Vec<Box<dyn Array>> = Vec::with_capacity(1);
    cs.push(ffi::import_array_from_c(array_ptr.add(0).read(), DataType::Int32).unwrap());

    1 as c_int
}

/// # Safety
/// `ArrowArray` and `ArrowSchema` must be valid
unsafe fn import(array: ffi::ArrowArray, schema: &ffi::ArrowSchema) -> Result<Box<dyn Array>> {
    let field = ffi::import_field_from_c(schema)?;
    ffi::import_array_from_c(array, field.data_type)
}

pub unsafe extern "C" fn from_chunks(arrsptr: *const ffi::ArrowArray, l: usize) {
    let mut cs: Vec<Box<dyn Array>> = Vec::with_capacity(l);

    for index in 0..l {
        cs.push(ffi::import_array_from_c(arrsptr.add(index).read(), DataType::Int64).unwrap())
    }
}



