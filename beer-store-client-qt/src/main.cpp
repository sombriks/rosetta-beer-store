#include <QGuiApplication>
#include <QQmlApplicationEngine>

#include "service/beerservice.h"
#include "model/beer.h"

int main(int argc, char *argv[])
{
    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);

    QGuiApplication app(argc, argv);

    QQmlApplicationEngine engine;

    qmlRegisterType<BeerService>("beer.store",1,0,"BeerService");
    qmlRegisterType<Beer>("beer.store",1,0,"Beer");

    const QUrl url(QStringLiteral("qrc:/qml/main.qml"));

    QObject::connect(&engine, &QQmlApplicationEngine::objectCreated,
                     &app, [url](QObject *obj, const QUrl &objUrl) {
        if (!obj && url == objUrl)
            QCoreApplication::exit(-1);
    }, Qt::QueuedConnection); 

    engine.load(url);

    return app.exec();
}
