
# Custom Protoc Plugin

This plugin is just for educational purpose which shows how we can create our own custom protoc plugin to generate custom code based on .proto files.

In this demo we are trying to generate code based on the custom protobuf options.
We are specifying a custom MessagOption name `validateEmail` and if it true we are generating a function name `ValidateEmail` in our custom coded file.

``Internal`` folder contain the newly generated code

In `examples` directory you will find example usage of different packages of [protobuf-go](https://github.com/protocolbuffers/protobuf-go) so you can get the idea of its usage.

## Installation

Install Prequisites

install [go](https://go.dev/doc/install) , [buf cli](https://buf.build/docs/installation) , [protoc](https://grpc.io/docs/protoc-installation/)

    
## Run Locally

Clone the project

```bash
  git clone git@github.com:salmanbao/custom-protoc-plugin.git
```

Go to the project directory

```bash
  cd custom-protoc-plugin
```

Install dependencies

```bash
  go mod tidy
```

Compile and run

```bash
  chmod +x ./compile

  ./compile
```
in our compile file we are executing our command where we are passing `dbtag` as a custom tag for newly generated structs

```bash



```
## Authors

- [@salmanbao](https://www.github.com/salmanbao)

