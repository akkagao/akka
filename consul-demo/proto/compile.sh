#!/bin/bash
protoc --go_out=plugins=grpc:. helloworld.proto 
protoc --doc_out=./doc --doc_opt=html,helloworld.html helloworld.proto 
protoc --doc_out=./doc --doc_opt=markdown,helloworld.md helloworld.proto 