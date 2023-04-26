# Rust Language

This is some notes on the Rust programming Language.
This is just my collection of notes on rust and some opinionated pros and cons.

# Table Of Contents

<!--toc:start-->

- [Rust Language](#rust-language)
- [Table Of Contents](#table-of-contents)
- [Documentation Resources](#documentation-resources)
- [Why Rust?](#why-rust)
  - [Performance](#performance)
  - [Reliability](#reliability)
  - [Productivity](#productivity)
- [Tooling](#tooling)
  - [Easy to install and manage with Rustup and Cargo](#easy-to-install-and-manage-with-rustup-and-cargo)
    - [Creating new Project with Cargo](#creating-new-project-with-cargo)
- [Hellow World](#hellow-world)
  - [Hello world Cargo.toml](#hello-world-cargotoml)
- [A Web Service Example with Axum](#a-web-service-example-with-axum)
  - [Axum Example Cargo.toml](#axum-example-cargotoml)
  - [Axum Example main.rs](#axum-example-mainrs)
  - [Example intput /output](#example-intput-output)
- [Threading](#threading)
  - [Basic thread example](#basic-thread-example)
  - [move Closures with Threads](#move-closures-with-threads)
- [Projects with more files (modules)](#projects-with-more-files-modules)
- [Cross Compilation with Rustup and Cargo](#cross-compilation-with-rustup-and-cargo)
- [Cross Compilation with Cross](#cross-compilation-with-cross)
- [Pros/Cons](#proscons)
  - [Pros](#pros)
  - [Cons](#cons)
  <!--toc:end-->

# Documentation Resources

Before really starting a new project in Rust it is highly recommended to at least skim [**the book**](https://doc.rust-lang.org/book) from the official rust documentation.

For hands on exercises the [rustlings](https://github.com/rust-lang/rustlings) is excellent and fun. It is also worth mentioning [Rust by example](https://doc.rust-lang.org/rust-by-example/).

# Why Rust?

Directly from [rust-lang.org](https://www.rust-lang.org/)

## Performance

Rust is blazingly fast and memory-efficient: with no runtime or garbage collector, it can power performance-critical services, run on embedded devices, and easily integrate with other languages.

## Reliability

Rust’s rich type system and ownership model guarantee memory-safety and thread-safety — enabling you to eliminate many classes of bugs at compile-time.

## Productivity

Rust has great documentation, a friendly compiler with useful error messages, and top-notch tooling — an integrated package manager and build tool, smart multi-editor support with auto-completion and type inspections, an auto-formatter, and more.

# Tooling

## Easy to install and manage with Rustup and Cargo

Install rustup

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

Update rust and tools

```bash
rustup self update # update rustup
rustup update # update rust and associated tools
```

### Creating new Project with Cargo

Creating a new project is super simple.

```bash
cargo new projectname # Create new folder with project structure
cd projectname
cargo run # compile and run in debug mode
```

# Hello World

The `cargo new` command normally creates a `main.rs` source file with the following content.

```rust
fn main() {
    println!("Hello, world!");
}
```

We can add some user input like this:

```rust
fn main() {
    let mut line = String::new();
    println!("Please enter your name:");
    std::io::stdin().read_line(&mut line).unwrap();
    println!("Hello, {}!",line.trim());
}
```

The resulting output would look something like this:

```
Please enter your name:
Din Djarin
Hello, Din Djarin!
```

## Hello world Cargo.toml

`cargo new` creates a new project folder structure including the `Cargo.toml` file which is the primary project configuration file.

Cargo.toml:

```toml
[package]
name = "myproject"
version = "0.1.0"
edition = "2021"
```

You can add a new dependency to the project with cargo without having to manually edit `Cargo.toml`

```bash
cargo add axum
```

Then `Cargo.toml` would look like the below and the axum crate will automatically be downloaded and compiled when you run `cargo run/build`

```toml
[package]
name = "myproject"
version = "0.1.0"
edition = "2021"

[dependencies]
axum = "0.6.17"
```

# A Web Service Example with Axum

To see something a bit more real lets look at the "hello world" example from the popular web framework "Axum".
This example is a webservice that prints hello work on the "/" root path and accepts some JSON input to create a user on the "/user/" path.

## Axum Example Cargo.toml

```toml
[package]
name = "hello_axum"
version = "0.1.0"
edition = "2021"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
axum = "0.6.16"
serde = { version = "1.0.160", features = ["derive"] }
serde_json = "1.0.96"
tokio = { version = "1.27.0", features = ["full"] }
tracing = "0.1.37"
tracing-subscriber = { version = "0.3.16", features = ["env-filter"] }
```

## Axum Example main.rs

```rust
use axum::{
    http::StatusCode,
    response::IntoResponse,
    routing::{get, post},
    Json, Router,
};
use serde::{Deserialize, Serialize};
use std::net::SocketAddr;

#[tokio::main]
async fn main() {
    // initialize tracing
    tracing_subscriber::fmt::init();

    // build our application with a route
    let app = Router::new()
        // `GET /` goes to `root`
        .route("/", get(root))
        // `POST /users` goes to `create_user`
        .route("/users", post(create_user));

    // run our app with hyper
    // `axum::Server` is a re-export of `hyper::Server`
    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    tracing::debug!("listening on {}", addr);
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

// basic handler that responds with a static string
async fn root() -> &'static str {
    "Hello, World!"
}

async fn create_user(
    // this argument tells axum to parse the request body
    // as JSON into a `CreateUser` type
    Json(payload): Json<CreateUser>,
) -> impl IntoResponse {
    // insert your application logic here
    let user = User {
        id: 1337,
        username: payload.username,
    };
    tracing::debug!("Created new user {:?}", user);

    // this will be converted into a JSON response
    // with a status code of `201 Created`
    (StatusCode::CREATED, Json(user))
}

// the input to our `create_user` handler
#[derive(Deserialize, Debug)]
struct CreateUser {
    username: String,
}

// the output to our `create_user` handler
#[derive(Serialize, Debug)]
struct User {
    id: u64,
    username: String,
}
```

## Example intput /output

For "/":

```
curl localhost:3000
Hello, World!
```

Creating user:

```
curl localhost:3000/users -H 'Content-Type: application/json' -d '{"username": "Din Djarin"}'
{"id":1337,"username":"Din Djarin"}
```

# Threading

Rust has what has been nick named fearless concurrency.
It turned out during development of Rust that the mechanisms used for memory safety together with the type system is very useful for safe concurrency.

## Basic thread example

```rust
use std::thread;
use std::time::Duration;

fn main() {
    let handle = thread::spawn(|| {
        for i in 1..10 {
            println!("Thread: {}", i);
            thread::sleep(Duration::from_millis(5));
        }
    });
    for i in 1..10 {
        println!("Main: {}", i);
        thread::sleep(Duration::from_millis(3));
    }

    handle.join().unwrap();
}
```

## move Closures with Threads

danger!

```rust
use std::thread;

fn main() {
    let v = vec![1, 2, 3];

    let handle = thread::spawn(|| {
        println!("Here's a vector: {:?}", v);
    });
    //drop(v);
    handle.join().unwrap();
}
```

`println` needs a reference to `v` so only a reference to `v` is captured in the closure.
Rust can't tell how long the thread will live so it is not certain if the reference to `v` will be valid when `println` is called.

See the nice error description from rustc:

```
❯ cargo run
   Compiling myproject v0.1.0 (/home/rpe/scratchpad/myproject)
error[E0373]: closure may outlive the current function, but it borrows `v`, which is owned by the current function
 --> src/main.rs:6:32
  |
6 |     let handle = thread::spawn(|| {
  |                                ^^ may outlive borrowed value `v`
7 |         println!("Here's a vector: {:?}", v);
  |                                           - `v` is borrowed here
  |
note: function requires argument type to outlive `'static`
 --> src/main.rs:6:18
  |
6 |       let handle = thread::spawn(|| {
  |  __________________^
7 | |         println!("Here's a vector: {:?}", v);
8 | |     });
  | |______^
help: to force the closure to take ownership of `v` (and any other referenced variables), use the `move` keyword
  |
6 |     let handle = thread::spawn(move || {
  |                                ++++

For more information about this error, try `rustc --explain E0373`.
error: could not compile `myproject` due to previous error
```

The solution is to use the `move` keyword to force the closure to take ownership of `v`.

```rust
use std::thread;

fn main() {
    let v = vec![1, 2, 3];

    let handle = thread::spawn(move || {
        println!("Here's a vector: {:?}", v);
    });

    handle.join().unwrap();
}
```

# Projects with more files (modules)

You can create modules for your code to separate your code into multiple files. For the full details on how a more advanced project structure looks like and how rust modules work please refer to [modules](https://doc.rust-lang.org/book/ch07-02-defining-modules-to-control-scope-and-privacy.html) section of "the book".

A simplified example of a project with some functionality split into a module called `mymod` with some features further split into a separate `mymodstuff` module.

```
src
├── main.rs
├── mymod
│  └── mymodstuff.rs
└── mymod.rs
```

```rust
//main.rs
pub mod mymod;
use crate::mymod::mymodstuff::Stuff;
fn main() {
    let stuff = crate::mymod::mymodstuff::Stuff {
        something: String::from("stuff1"),
    };
    let stuff2 = Stuff {
        something: String::from("stuff2"),
    };
    println!("stuff:\n{:#?}\nstuff2:\n{:#?}", stuff, stuff2)
}
```

```rust
//mymod.rs
pub mod mymodstuff;
```

```rust
//mymodstuff.rs
#[derive(Debug)]
pub struct Stuff {
    pub something: String
}
```

# Cross Compilation with Rustup and Cargo

Install new target toolchain for different architecture

```bash
rustup target add x86_64-pc-windows-gnu
sudo apt install mingw64
```

Create new hello world project and compile for your host and then cross compile

```bash
cargo new myproject
cd myproject
cargo build
cargo build --target x86_64-pc-windows-gnu
```

# Cross Compilation with Cross

For a more "batteries include" solution you can use "cross" project.
Cross relies docker or podman to get necessary dependencies so you don't need to provide e.g. mingw your self.

```bash
cargo install cross
```

```bash
cross build --target aarch64-unknown-linux-gnu
cross build --target x86_64-pc-windows-gnu
```

# Pros/Cons

This is an opinionated list of pros and cons in no particular order

## Pros

- Memory Safety
- Fealess Concurrency
- Fast
  - "Zero Cost Abstractions" (Similar to C++)
  - Usually a number of Rust web frameworks are in the top 10 of fast web (backend) frameworks. [benchmarks](https://www.techempower.com/benchmarks/#section=data-r21)
  - Also fast frontend/full stack frameworks. Similar performance to fast JS/TS frameworks [benchmarks](https://krausest.github.io/js-framework-benchmark/current.html#eyJmcmFtZXdvcmtzIjpbImtleWVkL2Rpb3h1cyIsImtleWVkL2xlcHRvcyIsImtleWVkL3JlYWN0LWhvb2tzIiwia2V5ZWQvc29saWQiLCJrZXllZC9zeWNhbW9yZSIsImtleWVkL3ZhbmlsbGFqcyIsImtleWVkL3Z1ZSJdLCJiZW5jaG1hcmtzIjpbIjAxX3J1bjFrIiwiMDJfcmVwbGFjZTFrIiwiMDNfdXBkYXRlMTB0aDFrX3gxNiIsIjA0X3NlbGVjdDFrIiwiMDVfc3dhcDFrIiwiMDZfcmVtb3ZlLW9uZS0xayIsIjA3X2NyZWF0ZTEwayIsIjA4X2NyZWF0ZTFrLWFmdGVyMWtfeDIiLCIwOV9jbGVhcjFrX3g4IiwiMjFfcmVhZHktbWVtb3J5IiwiMjJfcnVuLW1lbW9yeSIsIjIzX3VwZGF0ZTUtbWVtb3J5IiwiMjVfcnVuLWNsZWFyLW1lbW9yeSIsIjI2X3J1bi0xMGstbWVtb3J5IiwiMzFfc3RhcnR1cC1jaSIsIjM0X3N0YXJ0dXAtdG90YWxieXRlcyJdLCJkaXNwbGF5TW9kZSI6MSwiY2F0ZWdvcmllcyI6WzEsMiwzLDQsNV19)
- Modern tooling
  - Rustup for installing and managing toolchains and tools.
  - Cargo for building, testing etc. (really really good!)
- Large and very friendly community
  - Good library support for multiple different domains
- "General Purpose" Can run on small MCUs as well as on large "super computers"
  - Larger number of supported platforms
  - WebAssembly and Rust plays well together (write both your web frontend and backend in rust today)
- Pretty good adoption in industry. [2022 Review | The adoption of Rust in Business ](https://rustmagazine.org/issue-1/2022-review-the-adoption-of-rust-in-business/)
- No OOP

## Cons

- Steep(er) learning curve
- Hard(er) to write "unsafe" Rust code. It may be easier to write safe C or C++ than unsafe Rust. Zig might be more interesting as a modern language for writing unsafe code.
- Longer compile times
- No stable ABI for dynamic libraries(ABI is dependent on compiler version and flags and more?)
  - You probably have to use C ABI for shared system libraries (No worse than C++?).
- No formal specification (is rustc basically the specification?)
- Less governance on crates (packages) in official crate repository than some other languages.
  - Minor problems with name squatting
  - Security concerned users should audit all dependencies (anyway?).
- No OOP
