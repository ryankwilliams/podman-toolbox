# Podman toolbox

Collection of command line utilities used when working with podman. One
of the primary reasons behind this project was to learn how to use podman
directly within go programs using the (podman service bindings package)[podman-golang-bindings].

## Install

You can install/run `podman-toolbox` by any of the following methods:

1. Go install

```
go install github.com/ryankwilliams/podman-toolbox@main
```

2. Build binary

```
make build
./out/podman-toolbox
```

## Usage

In order to use this utility, you will need to have the podman service running.
You can do this by starting it from your terminal either by issuing:

```
podman system service -t 0 &
```

Or use the make target:

```
make start-podman-service
```

Refer to the `--help` option to view the sub-commands provided by `podman-toolbox`:

```
podman-toolbox --help
```

[podman-golang-bindings]: https://github.com/containers/podman/blob/main/pkg/bindings/README.md