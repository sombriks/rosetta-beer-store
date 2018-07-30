package beer.store;

import org.glassfish.embeddable.BootstrapProperties;
import org.glassfish.embeddable.GlassFish;
import org.glassfish.embeddable.GlassFishProperties;
import org.glassfish.embeddable.GlassFishRuntime;

public class Main {

	public static void main(String[] args) throws Exception {
		BootstrapProperties bootstrap = new BootstrapProperties();
    GlassFishRuntime runtime = GlassFishRuntime.bootstrap(bootstrap);
    GlassFishProperties glassfishProperties = new GlassFishProperties();
    glassfishProperties.setPort("http-listener", 3000);
    GlassFish glassfish = runtime.newGlassFish(glassfishProperties);
    glassfish.start();
	}

}
