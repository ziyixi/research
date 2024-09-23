load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "cloudmailin",
    srcs = ["cloudmailin.go"],
    importpath = "github.com/ziyixi/monorepo/self_host/packages/todofy/cloudmailin",
    visibility = ["//self_host/packages/todofy:__subpackages__"],
    deps = [
        "@com_github_johanneskaufmann_html_to_markdown//:html-to-markdown",
        "@com_github_tidwall_gjson//:gjson",
    ],
)
