# beer-store-service-java-jee

This sample repo shows how to build the beer store service.

The Java Enterprise way.

[Payara Server](https://www.payara.fish/) is the big boy here.
It's a JEE8 compliant application server therefore it provides a full-featured
enterprise platform. It means:

- [JAX-RS](https://docs.oracle.com/javaee/7/tutorial/jaxrs.htm#GIEPU) for REST API
- [JPA](https://docs.oracle.com/javaee/7/tutorial/persistence-intro.htm#BNBPZ) for database mappings and even for schema creation and initial load

We decided to use [H2](http://www.h2database.com/html/main.html) database since it's already there, but this is java, it 
connects with everything.

[Gradle](https://gradle.org/) is what java has closest to npm. We used it
there in order to generate initial eclipse support and to manage java
dependencies and classpath behaviors. For example, we don't need to refer the
entire payara server on dependencies inside build.gradle . Instead, we depend
on javaee-api 8.0 and mark it as provided.

Our last development environment piece is a vanilla [eclipse photon](https://www.eclipse.org/) IDE.
Any fresh eclipse, TBH. install the JEE version, grab payara tools from the
[eclipse marketplace](https://marketplace.eclipse.org/content/payara-tools) and unzip payara somewhere so you can point the IDE to it.

## How do i run this?

Now there's the tricky part.

### Option 1: prepare the server, pack the war and manualy deploy

Prepare a machine with at least 1GB of RAM and start your payara server there. 
There are many ways to make that part.

Then configure a connection pool and a jdbc datasource called **jdbc/beer-ds**.

Also do not forget to change app server port from 8080 to 3000.

See payara videos and docs to learn on how to make those configurations.

Once you have a full featured java enterprise application server up and 
running, build the project with the **gradle build** command. I expect you to 
have gradle (4.9 at least) available from command line. ou might use [scoop](https://scoop.sh/) or 
[sdkman](https://sdkman.io/) to get it working nice on your box. 

The build will assemble the **beer-store-service-java-jee.war** file and you'll 
need to theploy it as the context root on payara.

### Option 2: not available yet!

It's quite annoying nowadays the *java ee way* to do things. 

Not every project will need a swarm of clustered servers. 

Not all PaaS vendor offer war deploy on a jee-compliant app server.

Yet jee java is a good standard when we play nice the rules.

[Payara Micro](https://www.payara.fish/payara_micro) is a start, but it's still not as good as a simple *npm run dev*

Maybe a very tweaked build.gradle could do the work, but we'll study it later.
