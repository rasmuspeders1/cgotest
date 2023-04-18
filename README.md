[#](#) 

### Build C static library


```
cd build
./configure
make

```

It will create static lib: `libqrparser.a` in `build` directory

### Build an Run golang client

Follow this [link](https://go.dev/doc/install) to install `golang` compiler

```
cd bindings/go
make
./build/qrparser code justadummystring

```

### Build and Run Rust client

You have to have rust tools already installed. You can use use [rustup](https://rustup.rs/) to install rust.
Rustup is the officially recommended way to install and manage your rust toolchain(s) and associated tools like cargo.

```
cd binding/rust 
cargo run justadummystring
```

