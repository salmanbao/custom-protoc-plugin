package protobufexample

import (
	"bytes"
	"fmt"
	"time"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/salmanbao/practice/custom-protoc-plugin/gen/proto/user"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
)

func ProtobufExamples() {
	GetMessageValues()
	MessageToJson()
	GetMessageDescriptor()
	PTypes()
}
func GetMessageValues() {
	person := user.Person{
		Id:    1,
		Email: "example@gmail.com",
	}
	msg := person.ProtoReflect()
	// Get message descriptor
	descriptor := msg.Descriptor()

	for i := 0; i < descriptor.Fields().Len(); i++ {
		// Get the field descriptor
		fieldDesc := descriptor.Fields().Get(i)
		fieldName := fieldDesc.Name()
		// Get the value of field
		fieldValue := msg.Get(fieldDesc)

		switch fieldDesc.Kind() {
		case protoreflect.StringKind:
			fmt.Printf("Field: %s, Value: %s\n", fieldName, fieldValue.String())
		case protoreflect.Int32Kind:
			fmt.Printf("Field: %s, Value: %d\n", fieldName, fieldValue.Int())
		}
	}
}

/*
// [descriptor] examples
*/
func GetMessageDescriptor() {
	person := &user.Person{
		Id:    1,
		Email: "example@gmail.com",
	}

	// File Descriptor return all the info related to the .proto file
	fd, mdp := descriptor.MessageDescriptorProto(person)      // return File Descriptor , Message Descriptor Proto
	fmt.Print("[GetFileDescriptor:", *fd.Name, "]")           // output: proto/user/User.proto
	fmt.Print("[GetFileDescriptor:", (*fd).GetPackage(), "]") // output: user
	fmt.Print("[GetFileDescriptor:", (*fd).Dependency, "]")   // output: google/protobuf/descriptor.proto
	fmt.Print("[GetFileDescriptor:", (*fd).GetOptions(), "]") // output: return all options

	fmt.Print("[GetMessageDescriptor:", *mdp.Name, "]")           // output: Person
	fmt.Print("[GetMessageDescriptor:", (*mdp).GetOptions(), "]") // output: [user.validateEmail]:true
}

/*
// [jsonpb] examples
*/
func MessageToJson() {

	person := &user.Person{
		Id:    1,
		Email: "example@gmail.com",
	}

	// convert you message struct into json string
	marshaler := &jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(person)
	if err != nil {
		fmt.Print("ERR_MessageToJson")
	}
	fmt.Print("JsonString:", str)

	// convert json string into message struct
	unmarshaler := &jsonpb.Unmarshaler{}
	unmarshaler.Unmarshal(bytes.NewReader([]byte(str)), person)
	fmt.Print("JsonObject:", person)
}

/*
// [ptypes] examples
*/
func PTypes() {
	t := time.Second * 100
	d := durationpb.New(t)
	duration, err := ptypes.Duration(d)
	if err != nil {
		fmt.Print("PTypes:Error")
	}
	fmt.Print("PTypes:", duration)
}

/*
// [proto] examples
// buffer.go file have buffer related functions
// deprecated.go file have all deprecate functions
// discard.go has function to discard unknown fields from the message
// extensions.go has function for extension management
*/
func GetExtensions() {
	person := &user.Person{
		Id:    1,
		Email: "example@gmail.com",
	}

	proto.MessageV1(person) // convert message to MessageV1
	proto.MessageV2(person) // convert message to MessageV2
	proto.MessageReflect(person)

}
