load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "gemini",
    srcs = ["gemini.go"],
    embedsrcs = ["prompt.tmpl"],
    importpath = "github.com/ziyixi/monorepo/self_host/packages/todofy/gemini",
    visibility = ["//self_host/packages/todofy:__subpackages__"],
    deps = [
        "//self_host/packages/todofy/utils",
        "@com_github_gin_gonic_gin//:gin",
        "@com_github_google_generative_ai_go//genai",
        "@org_golang_google_api//option",
    ],
)
