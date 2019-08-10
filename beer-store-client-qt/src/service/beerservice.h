#ifndef BEERSERVICE_H
#define BEERSERVICE_H

#include <QObject>
#include <QList>

#include "../model/beer.h"

class BeerService : public QObject
{
    Q_OBJECT
public:
    explicit BeerService(QObject *parent = nullptr);
    QList<Beer>list(QString search,int page,int size);
    Beer find(int idbeer);

signals:

public slots:
};

#endif // BEERSERVICE_H
