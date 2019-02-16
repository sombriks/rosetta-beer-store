package beer.store.service;

import java.util.List;

import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;
import javax.transaction.Transactional;

import org.springframework.stereotype.Service;

import beer.store.model.Beer;

@Service
public class BeerService {
  @PersistenceContext
  private EntityManager em;

  public List<Beer> list(String search, int page, int pageSize) {
    String q = "select b from Beer b where b.titleBeer like :title or b.descriptionBeer like :description";
    return em.createQuery(q, Beer.class)//
        .setParameter("title", "%" + search + "%")//
        .setParameter("description", "%" + search + "%")//
        .setFirstResult(0 * page)//
        .setMaxResults(pageSize)//
        .getResultList();
  }

  public Beer find(int idBeer) {
    return em.find(Beer.class, idBeer);
  }

  @Transactional
  public Beer insert(Beer beer) {
    em.persist(beer);
    em.flush();
    return beer;
  }

  @Transactional
  public Beer update(Beer beer) {
    em.merge(beer);
    em.flush();
    return beer;
  }

  @Transactional
  public void del(int idBeer) {
    Beer beer = find(idBeer);
    em.remove(beer);
    em.flush();
  }

}