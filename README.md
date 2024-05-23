# L0 Task

## Overview
This is a demo service project designed to display order data through a simple interface. The data model is provided in JSON format.

## Objectives
1. Deploy PostgreSQL locally.
2. Create and configure a database and user.
3. Set up tables to store the provided data.
4. Develop the service.
5. Implement connection and subscription to a channel in NATS Streaming.
6. Store received data in the database.
7. Implement in-memory caching of received data.
8. Recover cache from the database in case of service failure.
9. Start an HTTP server to serve data by ID from the cache.
10. Develop a basic interface to display order data by ID.

## Steps to Complete

### 1. Setup PostgreSQL
- Install PostgreSQL locally.
- Create a new database.
- Create a new user and configure access.

### 2. Database Configuration
- Define and create the necessary tables to store the JSON data.

### 3. Service Development
- Develop the service using a language and framework of your choice.
- Ensure it can connect to the PostgreSQL database.

### 4. NATS Streaming Integration
- Implement connection to NATS Jetstream.
- Subscribe to the appropriate channel to receive order data.
- Store the received data into the PostgreSQL database.

### 5. Caching Mechanism
- Implement in-memory caching for the received data.
- Ensure the service recovers the cache from the database if it crashes.

### 6. HTTP Server
- Develop an HTTP server to provide order data by ID from the cache.
- Ensure efficient retrieval and minimal latency.

### 7. Interface Development
- Create a simple user interface to display order data by ID.
- Ensure the interface is user-friendly and functional.

## Tools and Technologies
- PostgreSQL
- NATS Jetstream
- HTTP server 
- Simple Flutter UI

## Running the Service
1. Start PostgreSQL and ensure the database is running.
2. Run the service to start the HTTP server.
3. Access the user interface to query order data by ID.