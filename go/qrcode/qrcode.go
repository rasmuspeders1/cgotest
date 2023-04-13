package qrcode

/*
#cgo CFLAGS: -I../../include
#cgo CXXFLAGS: -I../../include
#cgo LDFLAGS: -L../../build -lqrparser -lstdc++

#include <stdlib.h>
#include <sys/types.h>
#include "interface.h"
*/
import "C"
import (
	"errors"
	"log"
	"unsafe"
)

//var testQRCode = C.CString("MT:Y.K9042C00KA0648G00")

func Parse(code string) error {
	//
	var testQRCode = C.CString(code)

	log.Println("Start QRParser...")
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

	C.free(unsafe.Pointer(testQRCode))
	return nil
}
