load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")

buildifier(
    name = "buildifier-fix",
    lint_mode = "fix",
)

buildifier(
    name = "buildifier-check",
    lint_mode = "warn",
)

gazelle(
    name = "gazelle",
    prefix = "github.com/kigster/thermomatic",
)

go_library(
    name = "library",
    srcs = ["main.go"],
    importpath = "github.com/kigster/thermomatic",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "thermomatic",
    embed = [":library"],
    visibility = ["//visibility:public"],
)
