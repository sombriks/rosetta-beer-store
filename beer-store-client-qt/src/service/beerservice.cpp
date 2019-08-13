#include <QDebug>
#include "beerservice.h"

BeerService::BeerService(QObject *parent) : QObject(parent)
{

}

void BeerService::teste(QString x)
{
    qDebug() << x;
}

