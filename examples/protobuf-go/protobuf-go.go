package protobufgo

import (
	"fmt"

	userpb "github.com/salmanbao/practice/custom-protoc-plugin/gen/proto/user"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

func ProtobufGoExamples() {
	ProtoPkg()
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
}

func ProtojsonPkg() {

}
