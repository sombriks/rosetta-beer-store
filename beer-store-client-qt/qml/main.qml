import QtQuick 2.12
import QtQuick.Window 2.12
import QtQuick.Controls 2.12

import "webclient.js" as Client

Window {
    id: w1
    visible: true
    width: 640
    height: 480
    title: qsTr("Beer Store - Listing")


    TextField {
        id: textField
        y: 15
        height: 40
        text: qsTr("")
        focus: true
        anchors.right: parent.right
        anchors.rightMargin: 330
        anchors.left: parent.left
        anchors.leftMargin: 13
    }


    Button {
        id: bPrev
        x: 316
        y: 15
        text: qsTr("Anterior")
        anchors.right: parent.right
        anchors.rightMargin: 224
    }


    Button {
        id: bNext
        x: 422
        y: 15
        text: qsTr("Próximo")
        anchors.right: parent.right
        anchors.rightMargin: 118
    }


    Button {
        id: bSearch
        x: 528
        y: 15
        text: qsTr("Buscar")
        anchors.right: parent.right
        anchors.rightMargin: 12
        onClicked: Client.func()
    }

    ListView {
        id: listView
        anchors.bottomMargin: 16
        anchors.rightMargin: 12
        anchors.leftMargin: 13
        anchors.topMargin: 68
        anchors.fill: parent
        delegate: Item {
            x: 5
            width: 80
            height: 40
            Row {
                id: row1
                Rectangle {
                    width: 40
                    height: 40
                    color: colorCode
                }

                Text {
                    text: name
                    anchors.verticalCenter: parent.verticalCenter
                    font.bold: true
                }
                spacing: 10
            }
        }
        model: ListModel {
            ListElement {
                name: "Grey"
                colorCode: "grey"
            }

            ListElement {
                name: "Red"
                colorCode: "red"
            }

            ListElement {
                name: "Blue"
                colorCode: "blue"
            }

            ListElement {
                name: "Green"
                colorCode: "green"
            }
        }
    }
}











/*##^## Designer {
    D{i:1;anchors_height:400;anchors_width:615;anchors_x:13;anchors_y:61}D{i:2;anchors_width:297;anchors_x:13}
D{i:5;anchors_height:399;anchors_width:615;anchors_x:13;anchors_y:66}
}
 ##^##*/
