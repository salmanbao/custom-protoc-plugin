package plugin

import (
	"github.com/golang/protobuf/descriptor"
	"github.com/salmanbao/practice/custom-protoc-plugin/gen/proto/user"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type IDescriptor interface {
	Messages() protoreflect.MessageDescriptors
	Extensions() protoreflect.ExtensionDescriptors
}

func Generate(gen *protogen.Plugin, dbtag string) error {

	// here we will iterate over the .proto files
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		iterateFileMessages(gen, f, dbtag)
	}
	//
	LearnProto(gen) // just for learning
	return nil
}

func iterateFileMessages(gen *protogen.Plugin, file *protogen.File, dbtag string) *protogen.GeneratedFile {
	// first we will give a name to our new go files which will be generated by our plugin
	filename := file.GeneratedFilenamePrefix + ".pb.go"
	// here we will create a new file
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	// g.P function will write the custom code into the new file
	g.P("// Code generated by protoc-gen-go-custom. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName) // write the package name defined in .proto file
	g.P()
	g.P("import \"fmt\"") // here we will write required packages

	// here we will iterate over the all messages in the .proto file
	for _, message := range file.Messages {
		// This function will generate struct from the message
		generateStructs(message, g, dbtag)

		var msg user.Person
		_, md := descriptor.MessageDescriptorProto(&msg)
		// Get all the available custom options on the given message type i.e Person
		options := md.GetOptions()

		// _, err := proto.GetExtension(md.Options, user.E_ValidateEmail)
		// if err == nil {
		// val, ok := ext.(*bool)
		// if ok {
		// fmt.Print("validateEmail:", *val)
		// }
		// }

		// iterate over the all custom options
		options.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {

			if !fd.IsExtension() {
				return true
			}
			// fmt.Print("print: (", fd.Name(), " : ", v.Bool(), ")")
			if v.Bool() {
				g.P("func (p ", message.Desc.Name(), ")", " ValidateEmail() error {")
				g.P("if len(p.Email) < 5 {")
				g.P("  return fmt.Errorf(\"email length should be greater than zero\")")
				g.P("}")
				g.P("return nil")
				g.P("}")

			}

			return true
		})
	}
	return nil
}

func generateStructs(message *protogen.Message, g *protogen.GeneratedFile, dbtag string) {
	g.P("type ", message.GoIdent, " struct {")
	for _, field := range message.Fields {
		if field.Desc.Kind() != protoreflect.EnumKind {
			g.P(field.GoName, " ", field.Desc.Kind(), " `json:\"", field.Desc.Name(), "\" ", dbtag, ":\"", field.Desc.Name(), "\"`")
		} else {
			g.P(field.GoName, " ", protoreflect.Int32Kind, " `json:\"", field.Desc.Name(), "\" ", dbtag, ":\"", field.Desc.Name(), "\"`")
			defer generateEnums(field, g)
		}
	}
	g.P("}")
}

func generateEnums(f *protogen.Field, g *protogen.GeneratedFile) {
	g.P("const (")
	g.P("")
	for i := 0; i < f.Desc.Enum().Values().Len(); i++ {
		e := f.Desc.Enum().Values().Get(i)
		g.P(e.Name(), " =  ", e.Index())
	}
	g.P(")")
}
