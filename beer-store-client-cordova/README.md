# beer-store-cliente-cordova

This sample repo shows how to build the beer store cliente using
[Apache Cordova](https://cordova.apache.org/). It also features a vue project
devel mode and a few tricks to properly pacakge the cordova mobile app.

We're using vanilla cordova and vu cli 3 was used to the bundle step.

## How to run this

Open a terminal in this folder and:

```bash
npm install
npm run dev
```

Please note that the way that cordova works demands you to provide a valid
android or ios development environment.

In order to generate the mobile app (.apk, .ipa), you need to use the
[cordova command line tool](https://cordova.apache.org/#getstarted):

```bash
sudo npm -g install cordova
cordova platform add android
# if you're working with a mac to run into ios:
cordova platform add ios
cordova build
cordova run android
```

### Building for android

The easiest way to proper configure the android environment is installing
[android studio](https://developer.android.com/studio/?hl=pt-br) and download
the desired android sdk version.

Newer cordova versions can detect android tools if there is an android studio
installation.

### Building for iOS

The simplest way to locally build for ios is installing xcode and keeping it
up to date.

Once you run **cordova build ios**, open xcode with the following command:

```bash
open platforms/ios/hellocordova.xcworkspace
```

Then set up a device, an emulator, or just package it as generic ios device
in order to upload it to the [App Store](https://appstoreconnect.apple.com/login)
