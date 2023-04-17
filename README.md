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

