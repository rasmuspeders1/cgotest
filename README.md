[#](#) 

### BUILD C static library

```
cd build
./configure
make

```

It will create static lib: `libqrparser.a` in `build` directory

### BUILD an RUN golang client

```
cd go
make
./build/qrparser code justadummystring

```

### Build and Run Rust client

You have to have rust tools already installed. You can use use [rustup](https://rustup.rs/) to install rust.
Rustup is the officially recommended way to install and manage your rust toolchain(s) and associated tools like cargo.

```
cd crusttest
cargo run
```

