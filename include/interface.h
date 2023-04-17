#pragma once

typedef void (*callback)(const char *, void* context);

typedef struct Event {
    callback cb;
    void *context;
} Event_t;

typedef enum EnumRendezvous {
    SOFTAP = 1,
    BLE = 2,
    ONNETWORK = 3,
} EnumRendezvous_t;

typedef struct QSetup 
{
    unsigned long Version;
    unsigned long VendorID;
    unsigned long ProductID;
    EnumRendezvous_t Rendezvous;
    unsigned long Discriminator;
    unsigned long Passcode;
} QSetup_t;


#ifdef __cplusplus
    extern "C" {
#endif

    int QRParse(const char* in, QSetup_t* out);
    void PrintSetup(const QSetup_t* payload);

    void RegisterCallback(callback cb, void *context);

#ifdef __cplusplus
    }
#endif
