load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/kigster/thermomatic",
)

go_library(
    name = "library",
    srcs = ["main.go"],
    importpath = "github.com/kigster/thermomatic",
    visibility = ["//visibility:private"],
    deps = ["//internal/common:library"],
)

go_binary(
    name = "thermomatic",
    embed = [":library"],
    visibility = ["//visibility:public"],
)
