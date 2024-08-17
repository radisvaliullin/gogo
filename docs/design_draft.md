Design Draft of GoGo repo.

Abstract application demonstrating features and best practices of Golang Application.

The app contains components:
* HTTP Server exposing API. API provides methods to create `task`. Task will be put to the worker queue.
* Queue client (external service kafka or nats).
* Workers read tasks from the queue, handle and save to DB.
* Storage manager to interact with DB (Postgres).

Requirements:
* The App should follow best practices of Go.
* Scalable and concurrent.
* Testable.
* Implement graceful shutdown.
