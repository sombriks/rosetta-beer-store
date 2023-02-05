# beer-store-service-java-spring-boot

This sample repo shows how to build the beer store service using java with
spring boot, jpa and flywaydb in a relational way.

The [spring boot](http://spring.io/projects/spring-boot) handles routing, data
access, while [flywaydb](https://flywaydb.org/) does database migrations. Also,
we adopted the lombok agent to make java classes less verbose.

It's a [gradle](https://gradle.org/install/) based project and you must have the
**5.1.1 version** or superior.

If you wish to develop using visual studio code, don't forget to run gradle
tasks to generate `.classpath` eclipse files, since it uses it to properly
configure a java project.

See `gradle tasks` output to see complete list of available tasks.

## How to run this

Simply open a terminal on this folder and type:

```bash
gradle build
gradle bootRun
```

Unlike JEE version, the output jar is a ready-to-deploy solution more aligned
with modern web development. It can be executed as a simple jar file, converted
into a systemd service quite easily without previously environment configuration
rather than database or similar things.
