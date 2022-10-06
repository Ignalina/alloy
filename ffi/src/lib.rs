use libc::c_int;
use arrow2::datatypes::Field;
use arrow2::array::Array;
use arrow2::ffi::
{
    ArrowArray,
    ArrowSchema,
    import_array_from_c,
    import_field_from_c
};

// https://github.com/jorgecarleitao/arrow2/blob/main/examples/ffi.rs
#[no_mangle]
pub extern "C" fn call_with_ffi(array: ArrowArray, schema: &ArrowSchema) -> c_int {
    let field: Field = match unsafe
    { import_field_from_c(schema) } {
        Ok(f) => f,
        Err(e) => panic!("Could not import_field_from_c, {:?}", e),
    };

    println!("Imported field from schema: {:?}", field);
    
    let _arr: Box<dyn Array> = match unsafe
    { import_array_from_c(array, field.data_type) } {
        Ok(a) => a,
        Err(e) => panic!("Could not import_array_from_c, {:?}", e),
    };

    println!("Hello from Rust, with Arrow Array");

    1 as c_int
}

