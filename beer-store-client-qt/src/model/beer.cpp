#include "beer.h"

Beer::Beer(QObject *parent) : QObject(parent) {}

Beer::Beer(const Beer &beer) : QObject(beer.parent()) {
  this->idbeer = beer.idbeer;
  this->creationdatebeer = beer.creationdatebeer;
  this->titlebeer = beer.titlebeer;
  this->descriptionbeer = beer.descriptionbeer;
  this->idmedia = beer.idmedia;
}

Beer *Beer::fromJson(QJsonValue val) {
  Beer *b = new Beer();
  QJsonObject o = val.toObject();
  b->idbeer = o.value("idbeer").toInt();
  qDebug() << o;
  return b;
}

Beer &Beer::operator=(Beer &other)
{
    Beer *b = new Beer(other);
    return *b;
}
