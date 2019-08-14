#ifndef BEERSERVICE_H
#define BEERSERVICE_H

#include <QObject>
#include <QList>
#include <QUrl>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QNetworkRequest>

#include "../model/beer.h"

class BeerService : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString search READ getSearch WRITE setSearch NOTIFY searchChanged)
    Q_PROPERTY(int page READ getPage WRITE setPage NOTIFY pageChanged)
    Q_PROPERTY(int pageSize READ getPageSize WRITE setPageSize NOTIFY pageSizeChanged)
    Q_PROPERTY(QList<Beer> beers READ getBeers WRITE setBeers NOTIFY beersChanged)

public:
    explicit BeerService(QObject *parent = nullptr);

signals:
    void searchChanged();
    void pageChanged();
    void pageSizeChanged();
    void beersChanged();

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
    QList<Beer>getBeers();
    void setBeers(QList<Beer>beers);

private:
    QString search = "a";
    int page=1,pageSize=10;
    QList<Beer>beers;

    QNetworkAccessManager qnam;
};

#endif // BEERSERVICE_H
