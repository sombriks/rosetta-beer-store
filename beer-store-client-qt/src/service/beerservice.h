#ifndef BEERSERVICE_H
#define BEERSERVICE_H

#include <QJsonArray>
#include <QJsonDocument>
#include <QNetworkAccessManager>
#include <QNetworkRequest>
#include <QNetworkReply>
#include <QVariantList>
#include <QVariant>
#include <QObject>
#include <QUrl>

#include "../model/beer.h"

class BeerService : public QObject {
  Q_OBJECT
  Q_PROPERTY(QString search READ getSearch WRITE setSearch NOTIFY searchChanged)
  Q_PROPERTY(int page READ getPage WRITE setPage NOTIFY pageChanged)
  Q_PROPERTY(
      int pageSize READ getPageSize WRITE setPageSize NOTIFY pageSizeChanged)
  Q_PROPERTY(
      QVariantList beers READ getBeers WRITE setBeers NOTIFY beersChanged)
  Q_PROPERTY(Beer* selected READ getSelected WRITE setSelected NOTIFY selectedChanged)

public:
  explicit BeerService(QObject *parent = nullptr);

signals:
  void searchChanged();
  void pageChanged();
  void pageSizeChanged();
  void beersChanged();
  void teste1Changed();
  void selectedChanged();

public slots:
  void list();
  void prev();
  void next();
  QString getSearch();
  void setSearch(QString search);
  int getPage();
  void setPage(int page);
  int getPageSize();
  void setPageSize(int pageSize);
  QVariantList getBeers();
  void setBeers(QVariantList beers);
  Beer *getSelected();
  void setSelected(Beer *selected);

private:
  QString search = "";
  int page = 1, pageSize = 10;
  QVariantList beers;
  // XXX figure out why does it needs to be class member instead local variable
  // inside function
  QNetworkAccessManager qnam;
  Beer *selected;
};

#endif // BEERSERVICE_H
