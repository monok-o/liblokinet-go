package lokinet

/*
#cgo CFLAGS: -I ../external/lokinet/include
#cgo LDFLAGS: -L ../external/lokinet/build/llarp -llokinet -Wl,-rpath=../external/lokinet/build/llarp
#include <lokinet.h>
*/
import "C"

func T() {
	ctx := C.lokinet_context_new()
	C.lokinet_context_start(ctx)
}
