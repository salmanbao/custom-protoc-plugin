package protobufgo

import (
	"fmt"
	"io"
	"os"
	"time"

	newspb "github.com/salmanbao/practice/custom-protoc-plugin/gen/proto/news"
	userpb "github.com/salmanbao/practice/custom-protoc-plugin/gen/proto/user"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protorange"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"

	"github.com/jhump/protoreflect/desc/protoparse"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProtobufGoExamples() {
	// ProtoPkg()
	// ProtoJsonPkg()
	// ProtoTextPkg()
	// ReflectProtoReflectPkg()
	// ReflectProtoRegistryPkg()
	// ReflectProtoPathAndProtoRangePkg()
	ReflectProtoDescPkg()
}

func ProtoPkg() {
	person := &userpb.Person{
		Id:    1,
		Email: "example@gmail.com",
	}

	extType := userpb.E_ValidateEmail

	options := person.ProtoReflect().Descriptor().Options()
	// Here we are just manipulating the existing extension
	proto.SetExtension(options, extType, true)
	fmt.Print("[ProtoPkg:HasExtension: ", proto.HasExtension(options, extType), "]")
	ext := proto.GetExtension(options, extType)
	fmt.Print("[ProtoPkg:GetExtension: ", ext, "]")

	// Set Message Option dynamically
	// Here we are dynamically creating an extension and setting it
	customExt := &protoimpl.ExtensionInfo{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         101010100,
		Name:          "type",
		Tag:           "bytes,101010100,opt,name=type",
	}

	proto.SetExtension(options, customExt, "CustomExtension")
	fmt.Print("[ProtoPkg:HasExtension:CustomExt ", proto.HasExtension(options, customExt), "]")
	cust_ext := proto.GetExtension(options, customExt)
	fmt.Print("[ProtoPkg:GetExtension: ", cust_ext, "]")

	bz, err := proto.Marshal(person)
	if err != nil {
		fmt.Print("ERROR:ProtoPkg")
	}
	fmt.Print("ProtoPkg:", bz)
}

func ProtoJsonPkg() {
	person := &userpb.Person{
		Id:    1,
		Email: "example@gmail.com",
	}

	fmt.Print("ProtojsonPkg:", protojson.Format(person))
	bz, err := protojson.Marshal(person)
	if err != nil {
		fmt.Print("ProtojsonPkg:ERROR")
	}
	fmt.Print("ProtojsonPkg:Marshal:", bz)

	per := &userpb.Person{}

	protojson.Unmarshal(bz, per)
	fmt.Print("ProtojsonPkg:Unmarshal", per)
}

func ProtoTextPkg() {
	person := &userpb.Person{
		Id:    1,
		Email: "example@gmail.com",
	}

	fmt.Print("ProtoTextPkg:", prototext.Format(person))
	bz, err := prototext.Marshal(person)
	if err != nil {
		fmt.Print("ProtoTextPkg:ERROR")
	}
	fmt.Print("ProtoTextPkg:Marshal:", bz)

	per := &userpb.Person{}

	prototext.Unmarshal(bz, per)
	fmt.Print("ProtoTextPkg:Unmarshal", per)
}

func ProtoWireFormatPkg() {
	// use proto.Marshal and proto.Unmarshal instead of this package
}

// protoreflect provides interfaces to dynamically manipulate messages.
// This package has All Descriptors interfaces
//
//	╔═════════════════════╗
//	║ Descriptors         ║
//	╠═════════════════════╣
//	║ FileDescriptor      ║
//	║ MessageDescriptor   ║
//	║ FieldDescriptor     ║
//	║ OneofDescriptor     ║
//	║ EnumDescriptor      ║
//	║ EnumValueDescriptor ║
//	║ ServiceDescriptor   ║
//	║ MethodDescriptor    ║
//	╚═════════════════════╝
func ReflectProtoReflectPkg() {

	person := &userpb.Person{
		Id:    1,
		Email: "example@gmail.com",
	}

	// 	┌──────────────── New() ─────────────────┐
	// 	│                                        │
	// 	│         ┌─── Descriptor() ─────┐       │   ┌── Interface() ───┐
	// 	│         │                      V       V   │                  V
	//   ╔═════════════╗  ╔═══════════════════╗  ╔═════════╗  ╔══════════════╗
	//   ║ MessageType ║  ║ MessageDescriptor ║  ║ Message ║  ║ ProtoMessage ║
	//   ╚═════════════╝  ╚═══════════════════╝  ╚═════════╝  ╚══════════════╝
	// 		 Λ           Λ                      │ │  Λ                  │
	// 		 │           └──── Descriptor() ────┘ │  └─ ProtoReflect() ─┘
	// 		 │                                    │
	// 		 └─────────────────── Type() ─────────┘

	message := person.ProtoReflect() // returns Message
	_ = message.Descriptor()         // returns MessageDescriptor
	messageType := message.Type()    // returns MessageType
	_ = messageType.Descriptor()     // returns MessageDescriptor

	// iterate over the message fields
	message.Range(func(f protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		fmt.Printf("[Field: %s , Value: %s]", f.Name(), v.Interface())
		return true
	})

	// Create new Message
	newMessage := messageType.New()
	// Get the field descriptor
	id_fd := newMessage.Descriptor().Fields().ByName(protoreflect.Name("id"))
	id_fd.IsList() // Return true if field is List
	id_fd.IsMap()  // Return true if field is Map
	// Create a value
	idValue := protoreflect.ValueOfInt32(2) // This Package also provide other utility functions for creating Value type instances
	// Set the field and value in the Message
	newMessage.Set(id_fd, idValue)

	newMessage.Range(func(f protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		fmt.Printf("[Field: %s , Value: %s]", f.Name(), v.Interface())
		return true
	})
}

func ReflectProtoRegistryPkg() {
	person := &userpb.Person{
		Id:    1,
		Email: "example@gmail.com",
	}
	fd := person.ProtoReflect().Descriptor().ParentFile() // return the messgae FileDescriptor
	fmt.Printf("[ReflectProtoRegistryPkg => %s]", fd.Name())

	// protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
	// 	fmt.Printf("[ReflectProtoRegistryPkg:RangeFiles => %s]", fd.Path())
	// 	return true
	// })

	mt, err := protoregistry.GlobalTypes.FindMessageByName("user.Person")
	if err != nil {
		fmt.Print(" [message not found] ")
	}
	fmt.Printf("[MessageName => %s]", mt.Descriptor().Name())

	files := protoregistry.GlobalFiles

	user_fd, err := files.FindFileByPath("proto/user/User.proto")
	if err != nil {
		fmt.Print(" [file not found] ")
	} else {
		fmt.Printf("[File name => %s]", user_fd.Name())
	}

	// External/New .proto file which we have to register with global registry
	filename := "proto/product/Product.proto"
	// This is an external package used to parse .proto files
	parser := protoparse.Parser{Accessor: func(name string) (io.ReadCloser, error) {
		return os.Open(name)
	}}

	descriptors, err := parser.ParseFiles(filename)
	if err != nil {
		fmt.Print(" [unable to parse .proto files] ")
	}

	// Try to find Product.proto file before registering
	product_fd, err := files.FindFileByPath(filename)

	if err != nil {
		fmt.Print(" [Product.proto file not found] ")

		// Create a FileDescriptor from the given FileDescriptorProto
		product_fd, err = protodesc.NewFile(descriptors[0].AsFileDescriptorProto(), protoregistry.GlobalFiles)
		if err != nil {
			fmt.Print("unable to create FileDescriptorProto")
		}

		messageName := product_fd.Messages().Get(0).FullName() // Get the message name for reference

		// Register FileDescriptor in the Global Registry using the above FileDescriptor
		err = protoregistry.GlobalFiles.RegisterFile(product_fd)
		if err != nil {
			fmt.Print("Product.proto FileDescriptor is unable to register")
		} else {
			fmt.Printf("[ Product.proto is registered in gloabal registry => %s]", product_fd.Name())

			fmt.Print(" [Message we are looking for is: ", messageName, "] ")

		}

		for i := 0; i < product_fd.Messages().Len(); i++ {

			// instance of MessageType
			mt := dynamicpb.NewMessageType(product_fd.Messages().Get(i)) // Its also apply on enum and services

			//register MessageType in the global registry
			if err = protoregistry.GlobalTypes.RegisterMessage(mt); err != nil {
				fmt.Print("[ Message Type not found after registrations ]")
			}

			_, err := protoregistry.GlobalTypes.FindMessageByName(messageName)
			if err != nil {
				fmt.Print(" [Product message not found: ", err, "] ")
			} else {
				fmt.Print("[ Message Type found after registrations ]")
			}
		}

	}
}

func mustMarshal(m proto.Message) []byte {
	if m != nil {

		b, err := proto.Marshal(m)
		if err != nil {
			panic(err)
		}
		return b
	}
	return []byte{}
}

func ReflectProtoPathAndProtoRangePkg() {
	m := &newspb.Article{
		Author:  "Russ Cox",
		Date:    timestamppb.New(time.Date(2019, time.November, 8, 0, 0, 0, 0, time.UTC)),
		Title:   "Go Turns 10",
		Content: "Happy birthday, Go! This weekend we celebrate the 10th anniversary of the Go release...",
		Status:  newspb.Article_PUBLISHED,
		Tags:    []string{"community", "birthday"},
		Attachments: []*anypb.Any{{
			TypeUrl: "google.golang.org.BinaryAttachment",
			Value: mustMarshal(&newspb.BinaryAttachment{
				Name: "gopher-birthday.png",
				Data: []byte("<binary data>"),
			}),
		}},
	}

	// Traverse over all reachable values and print the path.
	protorange.Range(m.ProtoReflect(), func(p protopath.Values) error {
		fmt.Println(" [", p.Path[1:], "] ")
		return nil
	})
}

func ReflectProtoDescPkg() {
	// protodesc.NewFile() is already used above for converting FileDescriptorProto to FileDescriptor

	// External/New .proto file which we have to register with global registry
	filename := "proto/product/Product.proto"

	// Step-1 Parse .proto file
	// This is an external package used to parse .proto files
	parser := protoparse.Parser{Accessor: func(name string) (io.ReadCloser, error) {
		return os.Open(name)
	}}

	descriptors, err := parser.ParseFiles(filename)
	if err != nil {
		fmt.Print(" [unable to parse .proto files] ")
	}

	// Step-2 Convert FileDescriptor to FileDescriptorProto
	fileDescriptorProto := descriptors[0].AsFileDescriptorProto()

	// Step-3 Create a FileDescriptorSet for managing multiple .proto file info
	fileDescriptorSet := &descriptorpb.FileDescriptorSet{}
	fileDescriptorSet.File = append(fileDescriptorSet.File, fileDescriptorProto)

	// This will return the FileRegistry
	fileRegistry, err := protodesc.NewFiles(fileDescriptorSet)
	if err != nil {
		fmt.Print("Error: ", err)
	}
	fileRegistry.NumFiles()

	fileDescriptor, err := fileRegistry.FindFileByPath(filename)
	if err != nil {
		fmt.Print("[ File not found :", err, "] ")
	}

	// There also Other methods for conversion
	//	╔═══════════════════════════════╗
	//	║       Methods       			║
	//	╠═══════════════════════════════╣
	//	║ ToDescriptorProto   			║
	//	║ ToFileDescriptorProto   		║
	//	║ ToFieldDescriptorProto     	║
	//	║ ToOneofDescriptorProto     	║
	//	║ ToEnumDescriptorProto      	║
	//	║ ToEnumValueDescriptorProto    ║
	//	║ ToServiceDescriptorProto   	║
	//	║ ToMethodDescriptorProto    	║
	//	╚═══════════════════════════════╝

	// fdp := protodesc.ToFileDescriptorProto(fileDescriptor)

	// Will return ServiceDescriptorProto
	// fdp.Service

	// Will return EnumDescriptorProto
	// fdp.EnumType

	// Will return DescriptorProto
	// fdp.MessageType

	// Will return FieldDescriptorProto
	// fdp.MessageType[0].GetField()
	fileDescriptor.FullName()
}
