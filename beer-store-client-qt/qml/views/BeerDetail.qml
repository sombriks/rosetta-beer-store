import QtQuick 2.12
import QtQuick.Controls 2.12

import "../components"

Item {
    id: element
    BeerItem {
        anchors.right: parent.right
        anchors.rightMargin: 109
        anchors.left: parent.left
        anchors.leftMargin: 109
        anchors.bottom: parent.bottom
        anchors.bottomMargin: 200
        anchors.top: parent.top
        anchors.topMargin: 200
        beer: service.selected
        doTap: function (e) {
            rootWindow.navTo("beerList")
        }
    }
}

/*##^## Designer {
    D{i:0;autoSize:true;height:480;width:640}D{i:1;anchors_height:80;anchors_width:422;anchors_x:109;anchors_y:200}
}
 ##^##*/
