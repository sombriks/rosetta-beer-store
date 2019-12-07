# beer-store-client-nativescript

Nativescript "hybrid" app to consume the beer service

## requirements

- Node latest lts
- Java 8+
- Android platform SDK
- Android studio
- At least one virtual device or the real thing

Also do not forget to set the `$ANDROID_HOME` variable.

## how to run this

First do the very normal android development setup. Install android studio,
download sdk's, create an emulator and so on.

Second, install nativescript command line tool

```bash
npm --unsafe-perm install -g nativescript
```

Then enter this folder and type `tns run android`.


The nativescript client will perform lots of checks and tell you what must be
made in order to rock the boat. Be patient and set variables, download sdks,
create virtual devices and so on.

If it keeps complaining about not being able to find your emulator, try to call
it first and the try the command:

```bash
$ANDROID_HOME/emulator/emulator -avd pixel2
```

Where pixel2 is a nice and short name i gave to my emulator.
