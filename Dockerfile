FROM golang:1.25.1-alpine

# Instala dependências necessárias
RUN apk add --no-cache ca-certificates curl unzip protobuf sudo bash git

ENV GOLANG_VERSION 1.25.1
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH"

WORKDIR $GOPATH

# Instala os plugins Go para Protocol Buffers
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Instala gqlgen
RUN go install github.com/99designs/gqlgen@latest

# Ajusta o diretório de trabalho para dentro do app
WORKDIR /go/src/app

# Default command: roda a aplicação
CMD ["go", "run", "./cmd/SistemaDeOrdem/"]
