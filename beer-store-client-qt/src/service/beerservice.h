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

signals:

public slots:
    void teste(QString x);
    QString teste2();
};

#endif // BEERSERVICE_H
