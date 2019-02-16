package beer.store;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import beer.store.model.Beer;
import beer.store.service.BeerService;

@SpringBootTest
@RunWith(SpringRunner.class)
public class BeerStoreSuite {

  @Autowired
  private BeerService service;

  @Test
  public void contextLoads() {
    // proves that spring boot is up and running
  }

  @Test
  public void listBeers() {
    assertEquals(10, service.list("", 1, 10).size());
  }

  @Test
  public void findBeer() {
    assertEquals("Skol", service.find(9).getTitleBeer());
  }

  @Test
  public void insertPetra() {
    Beer petra = new Beer();
    petra.setTitleBeer("Petra");
    petra.setDescriptionBeer("Cerveja com cerveja deve ser!");
    service.insert(petra);
    assertNotNull(petra.getIdBeer());
  }

  @Test
  public void updateItaipava() {
    Beer itaipava = service.find(3);
    itaipava.setDescriptionBeer("A cerveja 100%");
    service.update(itaipava);
    itaipava = service.find(3);
    assertEquals("A cerveja 100%", itaipava.getDescriptionBeer());
  }

  @Test
  public void deleteHeinenken() {
    service.del(8);
    assertEquals(0, service.list("Heinenken", 1, 1).size());
  }
}