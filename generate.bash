#!/bin/bash

gobind -lang=go golang.org/x/mobile/example/java-go-rpc/rpc > rpc/go_rpc/rpc.go
gobind -lang=java golang.org/x/mobile/example/java-go-rpc/rpc > src/go/Rpc.java
