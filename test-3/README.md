### [ Test-3 ] Book Library API Application

<a href="https://orange-moon-296181.postman.co/workspace/ETC~6af78728-9fca-43de-b5c7-7caeb249ff69/collection/63c3057096ae5eadc4a63276"> Test-3 Postman Collection Link (gRPC) </a>

<!-- HOW TO RUN -->
## How To Run
Before running the application, please configure the database connection config in the config file [.config.yml](.config.yml) to your own database configuration.

After that, you can execute this Make command to run the app: 
```
    make run-grpc
```

To run the application using docker, you have to pull the docker image first, run this command:
```
    make docker-pull
```
This will pull the docker image from docker repository to your computer. The default repository used is: sir-avdul.ddns.net:8555/etcetera/test-3-grpc:latest

You can run this command if you prefer to build the image yourself (*the docker repository target is inside [Makefile](Makefile)):
```
    make docker-build
```

After docker done building the image, you can use that image to run app via docker-compose:
```
    make docker-run
```
<!-- TESTING COMMANDS -->
## Testing Commands
Execute this command to run the unit-test:
```
    make test
```
For integration test you can run this command:
```
    make test-integration
```
## NOTES

Sorry the integration test is currently unavailable as I don't have enough time to make it because of time conflict with other techincal test from other company.

As well as the database migration is not available, and the REDIS caching is not properly implemented(only insert is implemented, but the fetch is still to the PSQL not to REDIS).

Also the Validator has been set in the entity struct but I'm running out of time to finish it.
