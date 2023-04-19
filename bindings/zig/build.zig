const std = @import("std");

pub fn build(b: *std.Build) void {
    const optimize = b.standardOptimizeOption(.{});
    const exe = b.addExecutable(.{
        .name = "parser_cli",
        .root_source_file = .{ .path = "./src/main.zig" },
        .optimize = optimize,
    });
    b.default_step.dependOn(&exe.step);

    exe.addIncludePath("../../include");
    exe.addCSourceFile("../../src/interface.cpp", &[_][]const u8{});
    //exe.addObjectFile("../../build/libqrparser.a");

    exe.linkLibC();
    exe.linkLibCpp();
    exe.install();
}
