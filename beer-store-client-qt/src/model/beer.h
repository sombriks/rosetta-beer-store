#ifndef BEER_H
#define BEER_H

#include <QDebug>
#include <QObject>
#include <QDateTime>
#include <QJsonObject>
#include <QJsonValue>

class Beer : public QObject {
  Q_OBJECT
  Q_PROPERTY(int idbeer READ getIdbeer WRITE setIdbeer NOTIFY idbeerChanged)
  Q_PROPERTY(QDateTime creationdatebeer READ getCreationdatebeer WRITE setCreationdatebeer NOTIFY creationdatebeerChanged)
  Q_PROPERTY(QString titlebeer READ getTitlebeer WRITE setTitlebeer NOTIFY titlebeerChanged)
  Q_PROPERTY(QString descriptionbeer READ getDescriptionbeer WRITE setDescriptionBeer NOTIFY descriptionbeerChanged)
  Q_PROPERTY(int idmedia READ getIdmedia WRITE setIdmedia NOTIFY idmediaChanged)

public:
  explicit Beer(QObject *parent = nullptr);
  explicit Beer(const Beer &beer);
  static Beer *fromJson(QJsonValue val);
  Beer &operator=(Beer& other);

signals:
    void idbeerChanged();
    void creationdatebeerChanged();
    void titlebeerChanged();
    void descriptionbeerChanged();
    void idmediaChanged();

public slots:
    int getIdbeer();
    void setIdbeer(int idbeer);
    QDateTime getCreationdatebeer();
    void setCreationdatebeer(QDateTime creationdatebeer);
    QString getTitlebeer();
    void setTitlebeer(QString titlebeer);
    QString getDescriptionbeer();
    void setDescriptionBeer(QString descriptionbeer);
    int getIdmedia();
    void setIdmedia(int idmedia);

private:
    int idbeer;
    QDateTime creationdatebeer;
    QString titlebeer;
    QString descriptionbeer;
    int idmedia;

};

#endif // BEER_H
