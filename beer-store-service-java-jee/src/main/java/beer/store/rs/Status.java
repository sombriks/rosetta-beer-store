package beer.store.rs;

import javax.ws.rs.GET;
import javax.ws.rs.Path;

@Path("/status")
public class Status {

	@GET
	public String status() {return "online";}
}
