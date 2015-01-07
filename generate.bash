#!/bin/bash

gobind -lang=go github.com/seletskiy/java-go-rpc/rpc > rpc/go_rpc/rpc.go
gobind -lang=java github.com/seletskiy/java-go-rpc/rpc > src/go/Rpc.java
