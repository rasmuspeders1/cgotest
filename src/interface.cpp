#include "interface.h"
#include <cstring>
#include <iostream>

// Initialize.
Event_t event = Event_t{
    .cb = 0,
    .context = 0,
};

void PrintSetup(const QSetup_t* payload) {

    if (payload == 0) {
        std::cerr << "Invalid/not intiliazed QSetup\n";
        return;
    }

    std::cout << "--------------\n";
    std::cout << "Version:  " << payload->Version << "\n";
    std::cout << "VendorID: " << payload->VendorID << "\n";
    std::cout << "ProductID: " << payload->ProductID << "\n";
    std::cout << "--------------\n";
}

int QRParse(const char* in, QSetup_t* out) {

    if (out == 0) {
        std::cerr << "Invalid/not intiliazed QSetup\n";
        return -1;
    }

    if (in == 0 || strlen(in) == 0) {
        std::cerr << "Invalid input \n";
        return -1;
    }

    std::string codeString = std::string{in};
    std::cout << "Representation: '" << codeString << "'\n";

    // Dummy parser ...
    // Dummy parser. i.e. after parsing set values to struct.

    out->ProductID = 1234;
    out->VendorID = 4567;
    out->Version = 1;
    out->Discriminator = 3840;
    out->Passcode = 20202021;

    // Callback     
    if (event.cb != 0) {
        event.cb(out, event.context);
    }

    return 0;
}

void RegisterCallback(callback cb, void* context) {
    
    // assign
    event.cb = cb;
    event.context = context; 
}
