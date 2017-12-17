workspace(name = "fn_tenor")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.8.0/rules_go-0.8.0.tar.gz",
    sha256 = "8eaf2e62811169d9cf511209153effcb132826cea708b2f75d4dd5f9e57ea2aa",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()


git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    tag = "v0.3.0",
)

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
    container_repositories = "repositories",
)
container_repositories()

container_pull(
  name = "fn_go",
  registry = "registry.hub.docker.com",
  repository = "fnproject/go",
  tag = "1.9.2"
)