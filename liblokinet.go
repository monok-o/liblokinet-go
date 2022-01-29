package liblokinet

/*
#cgo CFLAGS: -I${SRCDIR}/external/lokinet/include
#cgo LDFLAGS: -L${SRCDIR}/lib -llokinet -Wl,-rpath=${SRCDIR}/lib
#include <lokinet.h>
*/
import "C"

var Ctx = C.lokinet_context_new()

func Start() bool {
	C.lokinet_context_start(Ctx)

	for {
		isReady := int(C.lokinet_wait_for_ready(100, Ctx))
		if isReady == 0 {
			return true
		}
	}
}

func Stop() {
	C.lokinet_context_stop(Ctx)
}

func Addr() string {
	return C.GoString(C.lokinet_address(Ctx))
}

func Connect(address string) {
	C.lokinet_outbound_stream(&C.struct_lokinet_stream_result{}, C.CString(address), nil, Ctx)
}
