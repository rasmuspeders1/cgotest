#![allow(non_upper_case_globals)]
#![allow(non_camel_case_types)]
#![allow(non_snake_case)]

include!(concat!(env!("OUT_DIR"), "/bindings.rs"));
use std::ffi::CString;
use std::env;
fn parse_qr_code(code: &str) {
    let mut setup: QSetup = QSetup::default();

    //Create raw pointer to mutable setup var. Create from mutable reference (&mut)
    let setup_ptr: *mut QSetup = &mut setup;

    let c_str_code = CString::new(code).unwrap();

    // The C functions are unsafe so we have to call them in an unsafe block
    unsafe {
        let res = QRParse(c_str_code.as_ptr(), setup_ptr);
        if res != 0 {
            println!("Error: {}", res);
        }
    }
    println!("QSetup: {:#?}", setup);
}

fn main() {
    println!("Hello, from rust");
    let args: Vec<String> = env::args().collect();
    if args.len() != 2{
        println!("You must provide exactly 1 QR Code string as first argument");
        return;
    }
    parse_qr_code(args.get(1).unwrap());
}
