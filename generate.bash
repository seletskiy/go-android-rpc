#!/bin/bash

mkdir -p src/go/groid
gobind -lang=go github.com/seletskiy/go-android-rpc/android > android/go_android/rpc.go
gobind -lang=java github.com/seletskiy/go-android-rpc/android > src/go/groid/Groid.java
