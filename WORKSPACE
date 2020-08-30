load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Golang/Buildifier Lab Start (if we do one)
# Skylib is an implicit dep,
http_archive(
    name = "bazel_skylib",
    sha256 = "97e70364e9249702246c0e9444bccdc4b847bed1eb03c5a3ece4f83dfe6abc44",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.0.2/bazel-skylib-1.0.2.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.0.2/bazel-skylib-1.0.2.tar.gz",
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8663604808d2738dc615a2c3eb70eba54a9a982089dd09f6ffe5d0e75771bc4f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.6/rules_go-v0.23.6.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.6/rules_go-v0.23.6.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    sha256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

http_archive(
    name = "com_google_protobuf",
    sha256 = "d8dfa7697e1b7c8f4d36a6838f476052f0ecf7cb49dbb13391623c05b7e1aa7b",
    strip_prefix = "protobuf-3.12.4",
    urls = ["https://github.com/protocolbuffers/protobuf/releases/download/v3.12.4/protobuf-all-3.12.4.zip"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()
# Golang/Buildifier Lab End

#Install proto dependencies
http_archive(
    name = "rules_proto",
    sha256 = "602e7161d9195e50246177e7c55b2f39950a9cf7366f74ed5f22fd45750cd208",
    strip_prefix = "rules_proto-97d8af4dc474595af3900dd85cb3a29ad28cc313",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

# Golang/Buildifier Lab End

# github.com/onsi/ginkgo
go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    remote = "https://github.com/onsi/ginkgo",
    tag = "v1.14.0",
    vcs = "git",
)

# github.com/onsi/ginkgo
go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    remote = "https://github.com/onsi/gomega",
    tag = "v1.10.1",
    vcs = "git",
)

go_repository(
    name = "com_github_nxadm_tail",
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "github.com/nxadm/tail",
    sum = "h1:DQuhQpB1tVlglWS2hLQ5OV6B5r8aGxSrPc5Qo6uTN78=",
    version = "v1.4.4",
)

go_repository(
    name = "org_golang_x_text",
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/text",
    sum = "h1:cokOdA+Jmi5PJGXLlLllQSgYigAEfHXJAERHVMaCc2k=",
    version = "v0.3.3",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "org_golang_x_net",
    commit = "5ccada7d0a7ba9aeb5d3aca8d3501b4c2a509fec",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "eb3733d160e74a9c7e442f435eb3bea458e1d19f",
    importpath = "gopkg.in/yaml.v2",
)

#Install proto dependencies
http_archive(
    name = "rules_proto_grpc",
    sha256 = "5f0f2fc0199810c65a2de148a52ba0aff14d631d4e8202f41aff6a9d590a471b",
    strip_prefix = "rules_proto_grpc-1.0.2",
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/1.0.2.tar.gz"],
)

load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_repos", "rules_proto_grpc_toolchains")

rules_proto_grpc_toolchains()

rules_proto_grpc_repos()

# end proto

http_archive(
    name = "rules_pkg",
    sha256 = "aeca78988341a2ee1ba097641056d168320ecc51372ef7ff8e64b139516a4937",
    urls = [
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.2.6-1/rules_pkg-0.2.6.tar.gz",
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.2.6/rules_pkg-0.2.6.tar.gz",
    ],
)

load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

rules_pkg_dependencies()
# Python etls lab end

# Docker rules begin https://github.com/bazelbuild/rules_docker#setup
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "4521794f0fba2e20f3bf15846ab5e01d5332e587e9ce81629c7f96c793bb7036",
    strip_prefix = "rules_docker-0.14.4",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.14.4/rules_docker-v0.14.4.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

container_pull(
    name = "debian10",
    # tag = "10.5-slim",
    digest = "sha256:e0a33348ac8cace6b4294885e6e0bb57ecdfe4b6e415f1a7f4c5da5fe3116e02",
    registry = "index.docker.io",
    repository = "library/debian",
)

# Docker rules end
