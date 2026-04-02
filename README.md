# GoMicro

This is a code I wrote for testing. It is a 'useless piece of code' which creates four microservices written in Go and uses a remote DB2 database to store the data.
The solution also uses three local docker images for deploying a Kakfa server for messaging, a PostgreSQL database and a MongoDB database.

* **Members**  
Reads group members from a remote DB2 database.
* **Locations**  
Reads locations from a remote DB2 database.
* **Sender**  
Gets a list of locations using REST to communicate with the Members service and sends them (using a goroutine for each of them) one at a time using Kafka to the Receiver service.
* **Receiver**  
Receives each of the locations sent through Kafka and stores them one at a time (using a goroutine for each of them) in a table in the Postgres database.
* **Receiver2**  
Receives each of the locations sent through Kafka and stores them one at a time (using a goroutine for each of them) in a collection in the Mongo database.

## Docker compose

The project includes a docker-compose file. Once the machine has docker and docker-compose installed, that file starts the three containers.

## Configuration

The project needs a .env file in the root folder which contains the information required to connect to each of the containers. That includes user/passwords and also the addresses for each of the microservices.

```
KAFKA_BROKER=localhost:9092

DB2_HOST=<host>
DB2_USER=<user>
DB2_PWD=<password>

POSTGRES_HOST=localhost
POSTGRES_USER=<user>
POSTGRES_PWD=<password>

MONGO_HOST=localhost
MONGO_USER=<user>
MONGO_PWD=<password>

LOCATIONS_SERVICE=http://localhost:8080
MEMBERS_SERVICE=http://localhost:8081
```





