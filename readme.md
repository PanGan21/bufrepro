# bufrepro

**bufrepro** is a blockchain scaffolded with Ignite cli v28.0.0. It's purpose is to showcase a bug in protobuf generation.

## Details

When importing `import "bufrepro/moda/base.proto";` in `modb/extension.proto` then the `extension.pulsar.go` file prefixes the local import with `cosmossdk.io/api` which should not happen for local modules.

**Note**: The repository is not building intentionally
