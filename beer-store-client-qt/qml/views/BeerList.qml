import QtQuick 2.12
import QtQuick.Controls 2.12

import "../components"
import beer.store 1.0

Item {

    property BeerService beerService: BeerService {
        id:service
    }

    onVisibleChanged:function(){
        if(visible) service.teste("olá")
    }

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
        onClicked: rootWindow.navTo("beerDetail")
    }

    ListView {
        id: listView
        anchors.bottomMargin: 16
        anchors.rightMargin: 12
        anchors.leftMargin: 13
        anchors.topMargin: 68
        anchors.fill: parent
        delegate: BeerItem {
            x: 5
            width: parent.width
            height: 40
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
