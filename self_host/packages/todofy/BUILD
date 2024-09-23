load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "todofy_lib",
    srcs = [
        "router.go",
        "todofy.go",
    ],
    importpath = "github.com/ziyixi/monorepo/self_host/packages/todofy",
    visibility = ["//visibility:private"],
    deps = [
        "//self_host/packages/todofy/cloudmailin",
        "//self_host/packages/todofy/gemini",
        "//self_host/packages/todofy/mailjet",
        "//self_host/packages/todofy/utils",
        "@com_github_gin_gonic_gin//:gin",
    ],
)

go_binary(
    name = "todofy",
    embed = [":todofy_lib"],
    pure = "on",
    static = "on",
    visibility = ["//visibility:public"],
)
