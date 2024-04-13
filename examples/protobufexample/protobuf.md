# Protobuf packages details

## github.com/golang/protobuf

### Packages
1. `proto` Package proto provides functionality for handling protocol buffer messages.
2. `jsonpb` Package jsonpb provides functionality to marshal and unmarshal between a protocol buffer message and JSON.
3. `descriptor:` Package descriptor provides functions for obtaining the protocol buffer descriptors of generated Go types.
4. `protoc-gen-go` protoc-gen-go is a plugin for the Google protocol buffer compiler to generate Go code.
    1. `descriptor` Package `descriptor` is the generated package for `google/protobuf/descriptor.proto`.
    2. `plugin` Package `plugin` is the generated package for `google/protobuf/compiler/plugin.proto` 
5. `ptypes` Package ptypes provides functionality for interacting with well-known types.
    1. `any` Package `any` is the generated package for `google/protobuf/any.proto`. 
    2. `empty` Package empty is the generated package for `google/protobuf/empty.proto`.
    3. `timestamp` Package `timestamp` is the generated package for `google/protobuf/timestamp.proto`
    4. `duration` Package `duration` is the generated package for `google/protobuf/duration.proto`.
    5. `wrappers` Package `wrappers` is the generated package for `google/protobuf/wrappers.proto`.
    6. `struct` Package `structpb` is the generated package for `google/protobuf/struct.proto`.