# beer-store-service-java-spring-boot

This sample repo shows how to build the beer store service using java with
spring boot, jpa and flywaydb in a relational way.

The [spring boot](http://spring.io/projects/spring-boot) handles routing, data
access, while [flywaydb](https://flywaydb.org/) does database migrations. Also,
we adopted the lombok agent to make java classes less verbose.

It's a [gradle](https://gradle.org/install/) based project and you must have the
**5.1.1 version** or superior.

If you wish to develop using visual studio code, don't forget to run gradle
tasks to generate `.classpath` eclipse files, since it uses it to proper
configure a java project.

## How to run this

Simply open a terminal on this folder and type:

```bash
gradle build
gradle bootRun
```

