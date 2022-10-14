use libc::c_int;
use arrow2::datatypes::{Field, DataType};
use arrow2::ffi::{ArrowSchema, export_field_to_c};

#[no_mangle]
pub extern "C" fn call_with_ffi_schema(
    ffi_schema: &ArrowSchema
) -> ArrowSchema {

    println!("[Rust]\tHello! Here is the schema from Go side: {:?}", ffi_schema);
    
    let field: Field = Field::new("F1-i32", DataType::Int32, true); 
    let schema: ArrowSchema = export_field_to_c(&field);
    
    println!("[Rust]\tcreated Field: {:?}", field);
    println!("[Rust]\tffi_schema created on rust side from rust-arrow2 datatype: {:?}", schema);

    println!("[Rust]\tsending the Schema to Go now...");

    schema
}
