import QtQuick 2.0

Item {
    property var beer
    id: element3
    height: 60
    Text {
        id: element
        y: 23
        text: beer.name
        anchors.left: parent.left
        anchors.leftMargin: 86
        font.pixelSize: 12
    }

    Text {
        id: element1
        y: 28
        text: beer.detail
        anchors.left: parent.left
        anchors.leftMargin: 185
        font.pixelSize: 12
    }
}

/*##^## Designer {
    D{i:1;anchors_x:86}D{i:2;anchors_x:152}D{i:3;anchors_x:240}
}
 ##^##*/
