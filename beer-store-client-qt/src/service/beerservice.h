#ifndef BEERSERVICE_H
#define BEERSERVICE_H

#include <QObject>
#include <QList>

#include "../model/beer.h"

class BeerService : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString search READ getSearch WRITE setSearch NOTIFY searchChanged)
    Q_PROPERTY(int page READ getPage WRITE setPage NOTIFY pageChanged)
    Q_PROPERTY(int pageSize READ getPageSize WRITE setPageSize NOTIFY pageSizeChanged)
public:
    explicit BeerService(QObject *parent = nullptr);

signals:
    void searchChanged();
    void pageChanged();
    void pageSizeChanged();

public slots:
    void list();
    void prev();
    void next();

    QString getSearch();
    int getPage();
    int getPageSize();
    void setSearch(QString search);
    void setPage(int page);
    void setPageSize(int pageSize);

private:
    QString search = "";
    int page=1,pageSize=10;

};

#endif // BEERSERVICE_H
