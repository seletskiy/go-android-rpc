#!/usr/bin/env bash

# Builds and runs the app on an android device.

set -e

./make.bash

adb install -r bin/groid-debug.apk

adb shell am start -a android.intent.action.MAIN \
    -n com.example.groid/com.example.groid.MainActivity
