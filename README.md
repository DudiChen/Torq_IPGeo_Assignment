# Interview Assignment: IP to Country Service in Go Language

## Overview

This project implements a simple and extensible **IP-to-Country REST API service**.  
Given an IP address, the service returns its associated **country** and **city**, while enforcing a configurable **rate-limiting mechanism**.

The service was implemented as part of a coding exercise, with a strong focus on:
- Clear and readable code
- Environment-based configuration
- Extensibility for multiple IP-to-country data stores
- Production-grade structure and behavior

---

## Features

- REST API endpoint to resolve IP addresses into location data
- Configurable request rate limiting (without external rate-limiting libraries)
- Pluggable datastore design (CSV-based store implemented)
- Environment variable–driven configuration
- Dockerized execution for fast startup
- Partial unit test coverage for the CSV datastore

---

## API Specification

### Endpoint

```
GET /v1/find-country?ip=<IP_ADDRESS>
```

### Successful Response

```json
{
  "country": "United-States",
  "city": "New-York"
}
```

### Error Response

```json
{
  "error": "error message"
}
```

Appropriate HTTP status codes are returned, including:
- `400` for invalid requests
- `404` if the IP address is not found
- `429` when the rate limit is exceeded
- `500` for internal server errors

---

## Configuration

The service is fully configured using **environment variables**.  
For local development, these can be set directly or via a `.env` file in the project root.

### Required Environment Variables

```env
DATABASE_TYPE="csv"
DATABASE_PATH="./csv/data.csv"
RATE_REQUESTS=1
RATE_INTERVAL="1s"
PORT=8080
```

---

## CSV Datastore

### Format

```
ip,city,country
```

### Example Entry

```
67.250.186.196,New-York,United-States
```

---

## Running the Service

### Using Docker

```bash
./run_ipgeo_service.sh
```

Service will be available at:

```
http://localhost:8080
```

---

## Testing

Partial unit testing is included under:

```
csv/store_test.go
```

---

## Author

**David (Dudi) Chen**  
Date: July 28th, 2024
