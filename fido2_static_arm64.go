package libfido2

/*
#cgo darwin,!libfido2_system LDFLAGS: -framework CoreFoundation -framework IOKit
#cgo darwin,!libfido2_system CFLAGS: 
#cgo darwin,libfido2_system LDFLAGS: -framework CoreFoundation -framework IOKit /usr/local/lib/libfido2.a /usr/local/opt/openssl@3/lib/libcrypto.a ${SRCDIR}/darwin/amd64/lib/libcbor.a
#cgo darwin,libfido2_system CFLAGS: -I/usr/local/opt/libfido2/include -I/usr/local/opt/openssl@3/include
*/
import "C"
