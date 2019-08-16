import QtQuick 2.0
import beer.store 1.0

Item {
    id: beerItem
    height: 80
    property Beer beer: model.modelData

    Rectangle {
        id: rectangle
        x: 10
        y: 10
        color: "#ffffff"
        radius: 6
        anchors.right: parent.right
        anchors.rightMargin: 10
        anchors.left: parent.left
        anchors.leftMargin: 10
        anchors.bottom: parent.bottom
        anchors.bottomMargin: 10
        anchors.top: parent.top
        anchors.topMargin: 10
        border.width: 2

        Text {
            id: element
            x: 99
            y: 8
            text: beer.titlebeer
            anchors.left: parent.left
            anchors.leftMargin: 128
            font.pixelSize: 12
        }

        Text {
            id: element2
            x: 13
            y: 22
            text: "#" + beer.idbeer
            font.pixelSize: 12
        }

        Text {
            id: element1
            x: 175
            y: 29
            text: beer.descriptionbeer
            anchors.left: parent.left
            anchors.leftMargin: 128
            font.pixelSize: 12
        }

        MouseArea {
            id: mouseArea
            anchors.fill: parent
            onClicked: function(){
                service.selected = beer;
                rootWindow.navTo("beerDetail")
            }

        }
    }
}




/*##^## Designer {
    D{i:2;anchors_x:86}D{i:4;anchors_x:152}D{i:5;anchors_height:60;anchors_width:640;anchors_x:240;anchors_y:0}
D{i:1;anchors_height:33;anchors_width:608;anchors_x:10;anchors_y:10}
}
 ##^##*/
