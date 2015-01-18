#!/bin/bash

set -e

mkdir -p src/go/android
rm -rf bin/ gen/

gobind -lang=go github.com/seletskiy/go-android-rpc/android/rpc > android/go_android/rpc.go
gobind -lang=java github.com/seletskiy/go-android-rpc/android/rpc > src/go/android/Rpc.java
