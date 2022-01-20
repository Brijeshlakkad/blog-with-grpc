#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc math/mathpb/math.proto --go_out=plugins=grpc:.
protoc blog/blogpb/blog.proto --go-grpc_out=plugins=grpc:.