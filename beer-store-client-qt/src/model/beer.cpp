#include "beer.h"

Beer::Beer(QObject *parent) : QObject(parent)
{

}

Beer::Beer(const Beer &beer) : QObject (beer.parent())
{

}

