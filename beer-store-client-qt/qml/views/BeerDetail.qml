import QtQuick 2.12
import QtQuick.Controls 2.12

Item {
    Text {
        id: xpto
        x: 263
        y: 169
        text: qsTr("Detalhe")
        font.pixelSize: 12
    }

    TapHandler {
       onTapped: function (e) {
           rootWindow.navTo("beerList")
        }
    }
}
