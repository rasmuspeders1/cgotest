package qrcode

// Calling Go function from gateway functions.

/*
#cgo CFLAGS: -I../../include
#include "interface.h"

extern void CallbackHandler(QSetup_t* setup, void* context);

void handler_cgo(QSetup_t* setup, void *context) {
	CallbackHandler(setup, context);
}
*/
import "C"
