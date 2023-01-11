load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "endpoints.go",
        "server.go",
    ],
    importpath = "github.com/neticdk/go-bitbucket/mock",
    visibility = ["//visibility:public"],
    deps = [
        "//bitbucket:go_default_library",
        "@com_github_julienschmidt_httprouter//:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["server_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//bitbucket:go_default_library",
        "@com_github_stretchr_testify//assert:go_default_library",
    ],
)
