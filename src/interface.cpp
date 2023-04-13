#include "interface.h"
#include <string>
#include <iostream>


void PrintSetup(const QSetup_t* payload) {
    
    std::cout << "--------------\n";
    std::cout << "Version:  " << payload->Version << "\n";
    std::cout << "VendorID: " << payload->VendorID << "\n";
    std::cout << "ProductID: " << payload->ProductID << "\n";
    std::cout << "--------------\n";
}

int QRParse(const char* in, QSetup_t* out)
{
    if (out == 0) {
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
    return 0;
}

