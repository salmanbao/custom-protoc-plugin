# Protobuf-go packages details

## github.com/protocolbuffers/protobuf-go

### Packages
1. `proto`: Package proto provides functions operating on protobuf messages such as cloning, merging, and checking equality, as well as binary serialization.
2. `encoding/protojson`: Package `protojson` serializes protobuf messages as JSON.
3. `encoding/prototext`: Package `prototext` serializes protobuf messages as the text format
4. `encoding/protowire`: Package `protowire` parses and formats the low-level raw wire encoding. Most users should use package `proto` to serialize messages in the wire format.
5. `reflect/protoreflect`: Package `protoreflect` provides interfaces to dynamically manipulate protobuf messages.
6. `reflect/protoregistry`: Package `protoregistry` provides data structures to register and lookup protobuf descriptor types.
7. `reflect/protodesc`: Package `protodesc` provides functionality for converting `descriptorpb.FileDescriptorProto` messages to/from the reflective `protoreflect.        FileDescriptor`
8. `reflect/protopath`: Package `protopath` provides a representation of a sequence of protobuf reflection operations on a message.
9. `reflect/protorange`: Package `protorange` provides functionality to traverse a protobuf message.
10. `testing/protocmp`: Package `protocmp` provides protobuf specific options for the `cmp` package.
11. `testing/protopack`: Package `protopack` aids manual encoding and decoding of the wire format.
12. `testing/prototest`: Package `prototest` exercises the protobuf reflection implementation for concrete message types.
13. `types/dynamicpb`: Package `dynamicpb` creates protobuf messages at runtime from protobuf descriptors.
14. `types/known/anypb`: Package `anypb` is the generated package for `google/protobuf/any.proto`.
15. `types/known/timestamppb`: Package `timestamppb` is the generated package for `google/protobuf/timestamp.proto`
16. `types/known/durationpb`: Package `durationpb` is the generated package for `google/protobuf/duration.proto`
17. `types/known/wrapperspb`: Package `wrapperspb` is the generated package for `google/protobuf/wrappers.proto`
18. `types/known/structpb`: Package `structpb` is the generated package for `google/protobuf/struct.proto`
19. `types/known/fieldmaskpb`: Package `fieldmaskpb` is the generated package for `google/protobuf/field_mask.proto`
20. `types/known/apipb`: Package `apipb` is the generated package for `google/protobuf/api.proto`
21. `types/known/typepb`: Package `typepb` is the generated package for `google/protobuf/type.proto`
22. `types/known/sourcecontextpb`: Package `sourcecontextpb` is the generated package for `google/protobuf/source_context.proto`
23. `types/known/emptypb`: Package `emptypb` is the generated package for `google/protobuf/empty.proto`
24. `types/descriptorpb`: Package `descriptorpb` is the generated package for `google/protobuf/descriptor.proto`
25. `types/pluginpb`: Package `pluginpb` is the generated package for `google/protobuf/compiler/plugin.proto`
26. `compiler/protogen`: Package `protogen` provides support for writing protoc plugins.
27. `cmd/protoc-gen-go`: The `protoc-gen-go` binary is a protoc plugin to generate a Go protocol buffer package.
