# liblokinet-go
## Experimental go binding for liblokinet

This is not stable, this is a test, DO NOT USE.

## Build

```
git clone --recursive https://github.com/monok-o/liblokinet-go
cd liblokinet-go
```

### 1. Lokinet

For build dependencies, look at [lokinet' readme](https://github.com/oxen-io/lokinet)

```sh
cd external/lokinet
mkdir build
cd build
# here we just set DBUILD_SHARED_LIBS=ON so it will build the lib
cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED_LIBS=ON
make -j$(nproc)
```

### 2. Run the test
```sh
cd ../../..
export LD_LIBRARY_PATH=${PWD}/external/lokinet/build/llarp
go run main.go
```
