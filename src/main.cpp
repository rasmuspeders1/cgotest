#include "interface.h"
#include <string>
#include <vector>
#include <iostream>

int main()
{
    const char* mCode = "MT:Y.K9042C00KA0648G00";
    
    QSetup_t payload;

    int err = QRParse(mCode, &payload);
    if (err == 0) {
        PrintSetup(&payload);
    } else {
        std::cout << "QR code parsing error: " << err << "\n";
    }

    return 0;
}


