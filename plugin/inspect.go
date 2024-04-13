package plugin

import (
	"fmt"
	"log"

	"github.com/salmanbao/practice/custom-protoc-plugin/gen/proto/user"
	protogen "google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

/*
// msg.Location.Path() => will return the location of message in .proto file with index
// msg.Location.Soucefile() => will return the path of proto file
// msg.GoIdent.GoName() => will return the Name of the generated Go type i.e Childre, Person in this case
// msg.GoIdent.GoImportPath() => will return the import path of .proto file <module/path-to-the-.proto file>
// msg.Desc.Name() => will return the Message Name i.e Person
// msg.Desc.FullName() => will return the full name the of message <package>.MessageName i.e user.Person in this case
*/
func InspectMessages(msgs []*protogen.Message) {
	for _, msg := range msgs {
		InspectFields(msg)
	}
}

/*
// f.Desc.FullName() => full name of the field inside the proto file i.e user.Person.id
// f.Desc.TextName() => just field name i.e id
// f.Desc.Syntax() => will return the protobuf file syntax version i.e proto3
// f.Desc.Number() => will return number assigned to field in .proto file
// f.Desc.Kind() => will return type of field i.e string
// f.Desc.JSONName() => will return json name of the field
// f.Desc.HasOptionalKeyword() => will return true/false either you defined optional keyword with the field or not
// f.Desc.Options() => with return options specified in the field .i.e deprecated:true
*/
func InspectFields(m *protogen.Message) {
	for _, f := range m.Fields {
		fmt.Print("(", f.GoName, ")", "=>")
	}
}

func LearnProto(gen *protogen.Plugin) {
	var msg user.Person
	msg.Id = 1
	msg.Email = "john@gmail.com"
	err := proto.CheckInitialized(&msg)
	if err != nil {
		log.Print(err)
	}

	// protobufexample.ProtobufExamples()
	// protobufgo.ProtobufGoExamples()

}
