package beer.store.rs;

import java.util.List;

import javax.ejb.Stateless;
import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;

import beer.store.model.Beer;

@Path("/beer")
@Stateless
public class Beers {

	@PersistenceContext(unitName = "jdbc/beer-pu")
	private EntityManager em;

	@GET
	@Path("/list")
	@Produces("application/json")
	public List<Beer> list() {
		return em.createQuery("select b from Beer b", Beer.class).getResultList();
	}
}
