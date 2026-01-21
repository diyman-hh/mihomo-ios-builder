# Mihomo iOS Build

This repository is tailored to build `Mihomo.xcframework` for iOS using GitHub Actions.

## How to Build

1.  Create a **New Public Repository** on GitHub (e.g., `my-mihomo`).
2.  Push all files in this folder to that repository.
3.  Go to the **Actions** tab in your GitHub repository.
4.  You should see "Build iOS XCFramework" running.
5.  Once completed, click on the workflow run and download the `Mihomo.xcframework` artifact.
6.  Unzip it and replace the existing framework in your ECMAIN project (`ECMAIN/Frameworks/Mihomo.xcframework` or root).

## Structure
- `mobile/mobile.go`: Bridge code exposing `Start` and `Stop` to iOS.
- `.github/workflows/ios.yml`: CI configuration to run `gomobile bind`.
