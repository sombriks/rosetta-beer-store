import QtQuick 2.12
import QtQuick.Window 2.12
import QtQuick.Controls 2.12

import "views"
import beer.store 1.0


ApplicationWindow {
    id: rootWindow
    visible: true
    width: 640
    height: 480
    title: qsTr("Beer Store - Listing")

    property string currentView: "beerList"

    BeerService {
        id:service
    }

    function navTo(view) {
        rootWindow.currentView=view
    }

    Component.onCompleted: service.list()

    BeerList {
        id: beerList
        anchors.fill: parent
        visible: rootWindow.currentView === "beerList"
    }

    BeerDetail {
        id: beerDetail
        anchors.fill: parent
        visible: rootWindow.currentView === "beerDetail"
    }
}























/*##^## Designer {
    D{i:1;anchors_x:627;anchors_y:11}
}
 ##^##*/
