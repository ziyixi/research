load("@gazelle//:def.bzl", "gazelle")
load("@rules_go//go:def.bzl", "go_binary", "go_library")

gazelle(name = "gazelle")

go_library(
    name = "mail2todo_lib",
    srcs = ["main.go"],
    importpath = "github.com/ziyixi/monorepo/self_host/packages/mail2todo",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "mail2todo",
    embed = [":mail2todo_lib"],
    visibility = ["//visibility:public"],
)
