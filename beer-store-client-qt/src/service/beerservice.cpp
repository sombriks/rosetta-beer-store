#include <QDebug>


#include "beerservice.h"

BeerService::BeerService(QObject *parent) : QObject(parent) {}

void BeerService::list() {
  QString req = "http://localhost:3000/beer/list?page=%1&pageSize=%2&search=%3";
  QUrl url(QString(req.arg(page).arg(pageSize).arg(search)));
  QNetworkReply *reply = qnam.get(QNetworkRequest(url));

  BeerService *self = this;

  connect(reply, &QNetworkReply::finished, this, [reply, url, self] {
    // @see https://doc.qt.io/qt-5/qnetworkreply.html#finished
    QByteArray bytes = reply->readAll();
    QJsonDocument doc = QJsonDocument::fromJson(bytes);
    QJsonArray arr = doc.array();
    for(int i=0; i < self->beers.size();i++){
        Beer *b = self->beers.at(i);
        delete b;
    }
    self->beers.clear();
    for(int i=0; i< arr.size(); i++){
        Beer *b = Beer::fromJson(arr.at(i));
        self->beers.push_back(b);
    }
    emit self->beersChanged();

    qDebug() << "url: " << url << " reply: " << arr.size();
  });
}

void BeerService::next() {
  // XXX só tem next se a lista for não-vazia
  if (beers.size() == pageSize) {
    page += 1;
    emit pageChanged();
  }
}

void BeerService::prev() {
  if (page > 1) {
    page -= 1;
    emit pageChanged();
  }
}

QString BeerService::getSearch() { return search; }

int BeerService::getPage() { return page; }

int BeerService::getPageSize() { return pageSize; }

void BeerService::setSearch(QString search) {
  this->search = search;
  emit searchChanged();
}

void BeerService::setPage(int page) {
  this->page = page;
  emit pageChanged();
}

void BeerService::setPageSize(int pageSize) {
  this->pageSize = pageSize;
  emit pageSizeChanged();
}

QList<Beer*> BeerService::getBeers() { return beers; }

void BeerService::setBeers(QList<Beer*> beers) {
  this->beers = beers;
  emit beersChanged();
}
