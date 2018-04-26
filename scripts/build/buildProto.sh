#!/bin/bash


# This script builds all of the .proto files in the repo into their proper locations
# It can be run stand alone or by running make protos 

set -e

GOSRC=$GOPATH/src

protoc  -I$GOSRC --go_out=plugins=grpc:$GOSRC \
	$GOSRC/github.com/VolSec/ReverseTraceroutePublic/datamodel/*.proto 

protoc  -I$GOSRC --go_out=plugins=grpc:$GOSRC \
	$GOSRC/github.com/VolSec/ReverseTraceroutePublic/controller/pb/*.proto 
    

protoc  -I$GOSRC --go_out=plugins=grpc:$GOSRC \
	$GOSRC/github.com/VolSec/ReverseTraceroutePublic/plcontroller/pb/*.proto 
    

protoc  -I$GOSRC --go_out=plugins=grpc:$GOSRC \
	$GOSRC/github.com/VolSec/ReverseTraceroutePublic/vpservice/pb/*.proto 
    

protoc  -I$GOSRC --go_out=plugins=grpc:$GOSRC \
	$GOSRC/github.com/VolSec/ReverseTraceroutePublic/atlas/pb/*.proto 
    
PATH_REPLACE=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration

protoc -I/usr/local/include \
       -I.  \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --go_out=$PATH_REPLACE,plugins=grpc:$GOSRC \
       $GOSRC/github.com/VolSec/ReverseTraceroutePublic/revtr/pb/*.proto 


protoc -I/usr/local/include \
       -I.  \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --grpc-gateway_out=logtostderr=true:$GOSRC \
       $GOSRC/github.com/VolSec/ReverseTraceroutePublic/revtr/pb/*.proto 
