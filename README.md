## Go gRPC Implementation

gRPC is an implementation of the Remote Protocol Buffer (RPC) made by Google
that utilizes the power of **protocol buffers** which serializes structured data 
in a language-neutral and extensible manner.

This project implements gRPC in Go in the form of a Calculator Service to understand
how to design, build, and deploy gRPC services in Go and the cloud.


1. `docker build -t calc-grpc:latest .`
2. `docker run -d -p 8080:8080 calc-grpc:latest`
