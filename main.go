package main

import (
	"errors"
	"flag"

	"github.com/salmanbao/practice/custom-protoc-plugin/plugin"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var dbtag *string

func main() {

	var flags flag.FlagSet

	// This flag is to generate custom tag for the structs
	dbtag = flags.String("dbtag", "", "db tag to use for generated structs")
	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	// options.Run function will run the pluging on all given .proto files
	options.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(*pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL.Enum())
		if *dbtag == "" {
			gen.Error(errors.New("dbtag must be set"))
		}
		return plugin.Generate(gen, *dbtag)
	})
}
