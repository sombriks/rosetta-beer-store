import QtQuick 2.0
import beer.store 1.0

Item {
    id: element3
    height: 60

    Text {
        id: element
        y: 23
        text: model.modelData
        anchors.left: parent.left
        anchors.leftMargin: 86
        font.pixelSize: 12
    }

    Text {
        id: element1
        y: 28
        text: model.modelData.descriptionbeer
        anchors.left: parent.left
        anchors.leftMargin: 185
        font.pixelSize: 12
    }

    MouseArea {
        id: mouseArea
        anchors.fill: parent
        onClicked: console.log(model.modelData)
    }

    Component.onCompleted: console.log(model.modelData)
}

/*##^## Designer {
    D{i:1;anchors_x:86}D{i:2;anchors_x:152}D{i:3;anchors_x:240}
}
 ##^##*/
