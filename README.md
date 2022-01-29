# liblokinet-go
## Experimental go binding for liblokinet

This is not stable, this is a test, DO NOT USE IN PRODUCTION !

## Use
```sh
go get -u github.com/monok-o/liblokinet-go
```

## Build Lokinet

For build dependencies, look at [lokinet's readme](https://github.com/oxen-io/lokinet)

### Clone the repo with its submodules
```
git clone --recursive https://github.com/monok-o/liblokinet-go
cd liblokinet-go
```

### Build lokinet
```sh
cd external/lokinet
mkdir build
cd build
# here we just set DBUILD_SHARED_LIBS=ON so it will build the lib
cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED_LIBS=ON
make -j$(nproc)
```

then move the generated `liblokinet.so` file from external/build/llarp to lib/