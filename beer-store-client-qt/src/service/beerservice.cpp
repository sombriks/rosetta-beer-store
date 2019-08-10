#include "beerservice.h"

BeerService::BeerService(QObject *parent) : QObject(parent)
{

}

QList<Beer> BeerService::list(QString search, int page, int size)
{

}

Beer BeerService::find(int idbeer)
{

}
