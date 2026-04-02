# GoMicro

This is a code I wrote for testing. It is a 'useless piece of code' which creates four microservices written in Go and uses a remote DB2 database to store the data.
The solution also uses two local docker images for deploying a Kakfa server for messaging and a PostgreSQL database.

* Members  
Reads group members from a remote DB2 database.
* Locations  
Reads locations from a remote DB2 database.
* Sender  
Gets a list of locations using REST to communicate with the Members service and sends them (using a goroutine for each of them) one at a time using Kafka to the Receiver service.
* Receiver  
Receives each of the locations sent through Kafka and stores them one at a time (using a goroutine for each of them) in a table in the Postgres database.
* Receiver2  
Receives each of the locations sent through Kafka and stores them one at a time (using a goroutine for each of them) in a collection in the Mongo database.


