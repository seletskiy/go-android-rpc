#!/usr/bin/env bash

# Copyright 2014 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e

if [ ! -f make.bash ]; then
	exit 1
fi

mkdir -p libs/armeabi-v7a src/go/rpc
ANDROID_APP=$PWD

ln -sf $GOPATH/src/golang.org/x/mobile/app/*.java $ANDROID_APP/src/go
ln -sf $GOPATH/src/golang.org/x/mobile/bind/java/Seq.java $ANDROID_APP/src/go

go generate
CGO_ENABLED=1 GOOS=android GOARCH=arm GOARM=7 \
	go build -ldflags="-shared" .
mv -f go-android-rpc libs/armeabi-v7a/libgojni.so
ant debug
