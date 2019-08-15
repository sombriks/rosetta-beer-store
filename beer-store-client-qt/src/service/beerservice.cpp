#include <QDebug>

#include "beerservice.h"

BeerService::BeerService(QObject *parent) : QObject(parent)
{

}

void BeerService::list()
{
//    qDebug() << "call list()";
    QString req = "http://localhost:3000/beer/list?page=%1&pageSize=%2&search=%3";
    QUrl url(QString(req.arg(page).arg(pageSize).arg(search)));
    QNetworkReply *reply = qnam.get(QNetworkRequest(url));
    connect(reply, &QNetworkReply::finished, this, [reply,url]
    {
        QByteArray bytes = reply->readAll(); // @see  https://doc.qt.io/qt-5/qnetworkreply.html#finished
        QJsonDocument doc = QJsonDocument::fromBinaryData(bytes);
        QJsonArray arr = doc.array();

        qDebug() << "url: " << url << " reply: " << arr.size();
    });

}

void BeerService::next()
{
    // XXX só tem next se a lista for não-vazia
    if(beers.size()==pageSize) {
        page+=1;
        emit pageChanged();
    }
}

void BeerService::prev()
{
    if(page>1) {
        page-=1;
        emit pageChanged();
    }
}

QString BeerService::getSearch()
{
    return search;
}

int BeerService::getPage()
{
    return page;
}

int BeerService::getPageSize()
{
    return pageSize;
}

void BeerService::setSearch(QString search)
{
    this->search=search;
    emit searchChanged();
}

void BeerService::setPage(int page)
{
    this->page=page;
    emit pageChanged();
}

void BeerService::setPageSize(int pageSize)
{
    this->pageSize=pageSize;
    emit pageSizeChanged();
}

QList<Beer> BeerService::getBeers()
{
    return beers;
}

void BeerService::setBeers(QList<Beer> beers)
{
    this->beers = beers;
    emit beersChanged();
}


