/*
MIT License

Copyright (c) 2022 Wilhelm Ågren & Rickard Ernst Björn Lundin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

use arrow2::array::Array;
use arrow2::datatypes::Field;
use arrow2::ffi;
use libc::c_uint;
use chrono::offset::Local;
use arrowalloy::api::OneShot;

macro_rules! info {
    ($($arg:tt)*) => {
        println!(
            "{} [INFO] [Rust]\t{}", 
            Local::now().format("[%Y-%m-%d %H:%M:%S]"),
            format!($($arg)*)
        );
    }
}
struct DummyBackend {
    ddlname: String,
}

impl OneShot for DummyBackend {
    fn set_lib(&self) -> usize {
        1

    }

    fn from_chunks(&self, arrays: Vec<Box<dyn Array>>) -> usize {

        info!("Hello! Once you Go Rust you never Go Back !");

        for (i, array) in arrays.iter().enumerate() {
            info!("array{}: {:?}", i + 1, array);
        }

        arrays.len()
    }
}


#[no_mangle]
pub unsafe extern "C" fn from_chunks_ffi(
    arrptr: *const ffi::ArrowArray,
    schptr: *const ffi::ArrowSchema,
    l: usize
    ) -> c_uint {

    info!("Hello! Reading the rust pointers now.");
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

    // @Wilhelm help me make a load specified .so file and create instance of OneShot instead of DummyBackend
    let  db: DummyBackend = DummyBackend{ ddlname: String::from("dummy") };
    db.from_chunks(arrays) as c_uint


}

