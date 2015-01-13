#!/bin/bash

gobind -lang=go github.com/seletskiy/go-android-rpc/rpc > rpc/go_rpc/rpc.go
gobind -lang=java github.com/seletskiy/go-android-rpc/rpc > src/go/rpc/Rpc.java
