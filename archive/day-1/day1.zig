const std = @import("std");
const data_test = @embedFile("example.txt");
const data_prod = @embedFile("data.txt");

pub fn part1(data: *const [:0]u8) void {
    var splits = std.mem.split(u8, data, "\n");
    while (splits.next()) |line| {
        const cols = std.mem.split(u8, line, " ");
        std.debug.print("%s %s\n", .{ cols[0], cols[1] });
    }
}

pub fn main() !void {
    part1(data_test.*[0..]);
    part1(data_prod.*[0..]);
}
