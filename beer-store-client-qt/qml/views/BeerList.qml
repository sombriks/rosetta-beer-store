import QtQuick 2.12
import QtQuick.Controls 2.12

import "../components"
Item {
    id: element

    ListView {
        id: listView
        anchors.bottomMargin: 50
        anchors.rightMargin: 12
        anchors.leftMargin: 13
        anchors.topMargin: 68
        anchors.fill: parent
        model: service.beers // QVariantList
        delegate: BeerItem {
            x: 5
            width: parent.width
            height: 80
            beer: model.modelData
            doTap: function (e) {
                service.selected = model.modelData;
                rootWindow.navTo("beerDetail")
            }
        }
    }

    TextField {
        id: textField
        y: 15
        height: 40
        text: service.search
        focus: true
        anchors.right: parent.right
        anchors.rightMargin: 330
        anchors.left: parent.left
        anchors.leftMargin: 13
        onAccepted: doBusca()
    }


    Button {
        id: bPrev
        x: 316
        y: 15
        text: qsTr("Anterior")
        anchors.right: parent.right
        anchors.rightMargin: 224
        onClicked: doPrev()
    }


    Button {
        id: bNext
        x: 422
        y: 15
        text: qsTr("Próximo")
        anchors.right: parent.right
        anchors.rightMargin: 118
        onClicked: doNext()
    }


    Button {
        id: bSearch
        x: 528
        y: 15
        text: qsTr("Buscar")
        anchors.right: parent.right
        anchors.rightMargin: 12
        onClicked: doBusca()
    }

    function doBusca(){
        service.setSearch(textField.text)
        service.list()
    }

    function doPrev(){
        service.prev()
        doBusca()
    }

    function doNext(){
        service.next()
        doBusca()
    }


    Text {
        id: pagina
        x: 488
        y: 457
        text: "Página "+service.page
        styleColor: "#ffffff"
        verticalAlignment: Text.AlignVCenter
        anchors.right: parent.right
        anchors.rightMargin: 12
        anchors.bottom: parent.bottom
        anchors.bottomMargin: 8
        font.pixelSize: 12
    }
}





/*##^## Designer {
    D{i:0;autoSize:true;height:480;width:640}
}
 ##^##*/
