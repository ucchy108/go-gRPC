FROM golang:1.18

WORKDIR /workspaces
COPY ./ /workspaces/

RUN apt-get update && apt-get install -y curl unzip protobuf-compiler


# RUN go get google.golang.org/grpc \
#     go get google.golang.org/grpc/cmd/protoc-gen-go-grpc