const std = @import("std");

const face = @cImport(
    @cInclude("interface.h"),
);

pub fn main() void {
    //
    const argv = std.os.argv;
    if (argv.len == 1) {
        std.debug.print("Provide QR code as string argument\n", .{});
        return;
    }

    std.debug.print("Hello from Zig!\n", .{});

    var setup = face.QSetup_t{
        .VendorID = 0,
        .ProductID = 0,
        .Discriminator = 0,
        .Passcode = 0,
        .Version = 0,
        .Rendezvous = 0,
    };

    // Pass argument as input string (QR string)
    const ret = face.QRParse(argv[1], &setup);
    if (ret == 0) {
        face.PrintSetup(&setup);
    } else {
        std.debug.print("Parsing failed", .{});
    }
}
