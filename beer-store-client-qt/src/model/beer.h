#ifndef BEER_H
#define BEER_H

#include <QDebug>
#include <QObject>
#include <QDateTime>
#include <QJsonObject>
#include <QJsonValue>

class Beer : public QObject {
  Q_OBJECT
public:
  explicit Beer(QObject *parent = nullptr);
  explicit Beer(const Beer &beer);
  static Beer *fromJson(QJsonValue val);
  Beer &operator=(Beer& other);

signals:

public slots:

private:
    int idbeer;
    QDateTime creationdatebeer;
    QString titlebeer;
    QString descriptionbeer;
    int idmedia;

};

#endif // BEER_H
