#ifndef BEER_H
#define BEER_H

#include <QObject>

class Beer : public QObject
{
    Q_OBJECT
public:
    explicit Beer(QObject *parent = nullptr);

signals:

public slots:
};

#endif // BEER_H
