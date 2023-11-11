/*
* MIT License
* 
* Copyright (c) 2023 Wilhelm Ågren & Rickard Ernst Björn Lundin
* 
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
* 
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
* 
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
* 
* File created: 2023-11-11
* Last updated: 2023-11-11
*/

use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn hello(
    message: *const libc::c_char,
) {
    let cstr = unsafe { CStr::from_ptr(message) };
    println!("From Rust code: {}", cstr.to_str().unwrap());
}


#[cfg(test)]
pub mod tests {

    use std::ffi::CString;
    use super::*;

    #[test]
    fn test_hello() {
        hello(CString::new("cool code").unwrap().into_raw());
    }
}
