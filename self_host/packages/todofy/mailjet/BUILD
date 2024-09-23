load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "mailjet",
    srcs = ["mailjet.go"],
    embedsrcs = ["todoDescription.tmpl"],
    importpath = "github.com/ziyixi/monorepo/self_host/packages/todofy/mailjet",
    visibility = ["//self_host/packages/todofy:__subpackages__"],
    deps = [
        "//self_host/packages/todofy/cloudmailin",
        "//self_host/packages/todofy/utils",
        "@com_github_gin_gonic_gin//:gin",
        "@com_github_mailjet_mailjet_apiv3_go_v4//:mailjet-apiv3-go",
    ],
)
