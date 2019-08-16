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
  b->creationdatebeer =
      QDateTime::fromString(o.value("creationdatebeer").toString());
  b->titlebeer = o.value("titlebeer").toString();
  b->descriptionbeer = o.value("descriptionbeer").toString();
  b->idmedia = o.value("idmedia").toInt();
  return b;
}

Beer &Beer::operator=(Beer &other) {
  Beer *b = new Beer(other);
  return *b;
}

int Beer::getIdbeer() { return idbeer; }

void Beer::setIdbeer(int idbeer) { this->idbeer = idbeer; emit idbeerChanged(); }

QDateTime Beer::getCreationdatebeer() { return creationdatebeer; }

void Beer::setCreationdatebeer(QDateTime creationdatebeer) {
  this->creationdatebeer = creationdatebeer;
  emit creationdatebeerChanged();
}

QString Beer::getTitlebeer() { return titlebeer; }

void Beer::setTitlebeer(QString titlebeer) {
  this->titlebeer = titlebeer;
  emit titlebeerChanged();
}

QString Beer::getDescriptionbeer() { return descriptionbeer; }

void Beer::setDescriptionBeer(QString descriptionbeer) {
  this->descriptionbeer = descriptionbeer;
  emit descriptionbeerChanged();
}

int Beer::getIdmedia() { return idmedia; }

void Beer::setIdmedia(int idmedia) {
  this->idmedia = idmedia;
  emit idmediaChanged();
}
