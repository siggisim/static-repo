load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@k8s_deploy//:defaults.bzl", "k8s_deploy")

# Use the filegroup rule to package up our static_files into a new target that
# we can include in our go_binary. To be safe, we only include html, css, and
# jpg files, rather than *everything*.
filegroup(
    name = "static_files",
    srcs = glob([
        "static/*.html",
        "static/*.css",
        "static/*.jpg",
    ]),
    visibility = ["//visibility:private"],
)

# Use the go_binary rule to create an executable from our main file. Depend on
# the static_files we created above so they are included.
go_binary(
    name = "main",
    srcs = ["main.go"],
    data = [
        ":static_files",
    ],
    deps = [
        "@io_bazel_rules_go//go/tools/bazel:go_default_library",
    ],
)

# Build a docker image similar to the go_binary above, but use the "go_image"
# rule from @io_bazel_rules_docker instead, which creates a docker image.
go_image(
    name = "server_image",
    srcs = ["main.go"],
    data = [
        ":static_files",
    ],
    deps = [
        "@io_bazel_rules_go//go/tools/bazel:go_default_library",
    ],
)

k8s_deploy(
    name = "deploy",
    cluster = "gke_$(GKE_PROJECT)_$(GKE_ZONE)_$(GKE_CLUSTER)",
    images = {
        "$(REGISTRY_URL)server:server_image": ":server_image",
    },
    substitutions = {
        "%{REGISTRY_URL}": "$(REGISTRY_URL)",
    },
    tags = ["manual"],
    template = ":deployment.yaml",
)


build changes