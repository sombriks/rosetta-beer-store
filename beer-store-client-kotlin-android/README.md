# beer-store-client-kotlin-android

This project features a native android mobile solution to consume our
beer service.

## how do i run this

You need [Android Studio](https://developer.android.com/studio), 
~~but Intellij Ultimate might works too~~ the 3.5 version or superior.

Open the project, set up your emulator or device and rock the boat.

Also make sure to correctly point to the service address, since the emulator and
your development machine aren't the same thing.

## Noteworthy

Kotlin coroutines does a nice job handling async calls to the REST API with
Retrofit.