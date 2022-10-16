

use arrow2::{array::Array, datatypes::DataType, error::Result,ffi};
use libc::c_int;

#[no_mangle]
pub unsafe extern "C" fn from_chunks_ffi(
    arrptr:  *const ffi::ArrowArray,
    schptr: *const ffi::ArrowSchema,
    l: usize
    ) -> c_int {



    let mut cs: Vec<Box<dyn Array>> = Vec::with_capacity(l);

    for index in 0..l {
            let field = ffi::import_field_from_c(&schptr.add(index).read()).unwrap();
            let array = ffi::import_array_from_c(arrptr.add(index).read(),field.data_type);
//          cs.push(array.unwrap());
            let a =array.unwrap();
            println!("array {:?}",a);
    }

    l as c_int
}




