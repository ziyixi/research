load("//tools:make_executable.bzl", "make_executable")

# Rule to make the script executable
make_executable(
    name = "update_executable",
    src = "update.sh",
)

# `sh_binary` depends on the executable produced by `make_executable`
sh_binary(
    name = "update",
    srcs = [":update_executable"],
    visibility = ["//visibility:public"],
)
