load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "utils",
    srcs = [
        "gin_utils.go",
        "utils.go",
    ],
    importpath = "github.com/ziyixi/monorepo/self_host/packages/todofy/utils",
    visibility = ["//self_host/packages/todofy:__subpackages__"],
    deps = [
        "@com_github_gin_gonic_gin//:gin",
        "@com_github_go_resty_resty_v2//:resty",
    ],
)
