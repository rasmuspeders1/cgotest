package qrcode

/*
#cgo CFLAGS: -I../../include
#cgo LDFLAGS: -L../../build -lqrparser -lstdc++

#include <stdlib.h>
#include <sys/types.h>
#include "interface.h"

// Forward declaration of gateway function
extern void handler_cgo(QSetup_t *setup, void* context);

*/
import "C"
import (
	"errors"
	"log"
	"unsafe"
)

//var testQRCode = C.CString("MT:Y.K9042C00KA0648G00")

//export CallbackHandler
func CallbackHandler(setup *C.QSetup_t, context *C.void) {
	//
	log.Printf("Got callback, pin code: %d\n", setup.Passcode)
}

func Parse(code string) error {
	//
	var testQRCode = C.CString(code)
	defer C.free(unsafe.Pointer(testQRCode))

	log.Println("Start QRParser...")
	cb := (C.callback)(unsafe.Pointer(C.handler_cgo))
	C.RegisterCallback(cb, unsafe.Pointer(nil))

	setup := &C.QSetup_t{}
	err := C.QRParse(testQRCode, setup)
	if err != 0 {
		return errors.New("QR code parsing error")
	}

	//log.Printf("Setup %+v", setup)
	log.Println("--------------------------")
	log.Printf("Version: %d", setup.Version)
	log.Printf("Pin code: %d", setup.Passcode)
	log.Printf("Discriminator: %d", setup.Discriminator)
	log.Println("--------------------------")

	return nil
}
