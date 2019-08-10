import QtQuick 2.12
import QtQuick.Window 2.12
import QtQuick.Controls 2.12

import "views"
import "components"
import beer.store 1.0

ApplicationWindow {
    id: rootWindow
    visible: true
    width: 640
    height: 480
    property BeerService beerService
    title: qsTr("Beer Store - Listing")

    onVisibleChanged:function(){

    }

    property string currentView: "beerList"

    function navTo(view) {
        rootWindow.currentView=view
    }

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
