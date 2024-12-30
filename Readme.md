# Simple Key-Value Database

A lightweight, file-based key-value database implementation in Go with REST API support.

## Features

- Persistent storage using file system
- Thread-safe operations with mutex locks
- RESTful API endpoints using Gin framework
- Basic CRUD operations (Set, Get, Delete)
- Environment variable configuration
- JSON-based data storage
- Frontend SDK support via @agastyagaur/gogaur-sdk (built by me)

## Prerequisites

- Go 1.16 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone <https://github.com/Agastya18/gogaur-db.git>
cd <repository-name>
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create a `.env` file in the root directory:
```env
PORT=8080
```

## API Endpoints

### Set a Key-Value Pair
```http
POST /set
Content-Type: application/json

{
    "key": "example_key",
    "value": "example_value"
}
```

### Get a Value by Key
```http
GET /get?key=example_key
```

### Delete a Key
```http
DELETE /delete?key=example_key
```

## Usage

1. Start the server:
```bash
go run cmd/main.go
```

2. The server will start on the configured port (default: 8080)

### Usage with React

```javascript

npm i @agastyagaur/gogaur-sdk

import { GogaurDB } from '@agastyagaur/gogaur-sdk';

// Initialize the client
const db = new GogaurDB('http://localhost:8080');

// Set a value
await db.set('user1', { name: 'John Doe' });

// Get a value
const user = await db.get('user1');

// Delete a value
await db.delete('user1');
```


## Data Structure

The database uses a simple file-based storage system with an in-memory index:
- Data is stored in a file (`mydb.data`) as JSON-encoded key-value pairs
- An in-memory map maintains the file offsets for quick lookups
- Concurrent access is handled using read-write mutex locks

## Error Handling

The API returns appropriate HTTP status codes:
- 200: Successful operation
- 400: Invalid request
- 404: Key not found
- 500: Internal server error

## Limitations

- No data compression
- No automatic data cleanup (deleted entries remain in file)
- In-memory index needs to be rebuilt on startup
- Basic error handling

## Contributing

Feel free to submit issues, fork the repository, and create pull requests for any improvements.

## License

[Your chosen license]