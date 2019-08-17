# beer-store-client-qt

This client uses the Qt platform to consume the beer service.

## How do i run this

In order to make this project work, you need to install
[The Qt framework and Qt Creator](https://www.qt.io/download). The open source
edition is good enough

Make sure that you have a few native tools (make,g++,clang and so on) properly
installed.

For fedora do:

```bash
# if unsure about the package name use sudo dnf grouplist
sudo dnf groupinstall "C Development Tools and Libraries"
sudo dnf install mesa-libGL-devel
sudo dnf install clang
```

Once you got everything ok, run qtcreator (it will appear on your menu, but
it's inside the Qt installation) and open the beer-store-client-qt.pro file.

This project was tested against the QT LTS 5.12.4 version.
