use arrow2::array::Array;
use arrow2::datatypes::Field;
use arrow2::ffi;
use libc::c_uint;
use chrono::offset::Local;

macro_rules! info {
    ($($arg:tt)*) => {
        println!(
            "{} [INFO] [Rust]\t{}", 
            Local::now().format("[%Y-%m-%d %H:%M:%S]"),
            format!($($arg)*)
        );
    }
}

#[no_mangle]
pub unsafe extern "C" fn from_chunks_ffi(
    arrptr: *const ffi::ArrowArray,
    schptr: *const ffi::ArrowSchema,
    l: usize
    ) -> c_uint {

    info!("Hello! Reading the ffi pointers now.");
    let mut arrays: Vec<Box<dyn Array>> = Vec::with_capacity(l);

    // Lets actually do proper Rust error handling. We pattern match on the result and
    // handle the error approprietly instead of unwrap. Also, lets return the the length
    // of the built Vec, instead of just returning the provided arg l.
    for index in 0..l {
            let field: Field = match ffi::import_field_from_c(&schptr.add(index).read()) {
                Ok(f) => f,
                Err(e) => panic!("Could not import_field_from_c: {:?}", e),
            };

            let array: Box<dyn Array> = match ffi::import_array_from_c(arrptr.add(index).read(), field.data_type) {
                Ok(a) => a,
                Err(e) => panic!("Could not import_array_from_c: {:?}", e),
            };

            arrays.push(array);
    }

    for (i, array) in arrays.iter().enumerate() {
        info!("array{}: {:?}", i + 1, array);
    }

    arrays.len() as c_uint
}

