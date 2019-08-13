#include <QDebug>
#include "beerservice.h"

BeerService::BeerService(QObject *parent) : QObject(parent)
{

}

void BeerService::list()
{
    qDebug() << "listing from service ";
    qDebug() << "search: " << search << " page: " << page << " pageSize: " << pageSize << endl;
    // TODO fazer a busca no serviço
}

void BeerService::next()
{
    // XXX só tem next se a lista for não-vazia
}

void BeerService::prev()
{
    if(page>1) {
        page-=1;
        emit pageChanged();
    }
}

QString BeerService::getSearch()
{
    return search;
}

int BeerService::getPage()
{
    return page;
}

int BeerService::getPageSize()
{
    return pageSize;
}

void BeerService::setSearch(QString search)
{
    this->search=search;
    emit searchChanged();
}

void BeerService::setPage(int page)
{
    this->page=page;
    emit pageChanged();
}

void BeerService::setPageSize(int pageSize)
{
    this->pageSize=pageSize;
    emit pageSizeChanged();
}


