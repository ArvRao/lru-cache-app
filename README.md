## About
An LRU Cache full stack application with **get** and **set** methods.
The cache supports a configurable maximum capacity (1024 items) and expiration times for each cache entry.
The cache is thread-safe, using a mutex to protect against concurrent access.
Backend built using Golang Fiber framework. The server is set up to handle API requests and listens on port `3000`.
It includes optional improvements like graceful shutdown handling.

### Setup

### Backend
> Make sure to have Go 1.21.1 installed in local machine
1. Navigate to `/backend` directory from root directory:

`cd /backend`

2. Download the necessary dependencies:

`go mod download`

3. Run the application on localhost at port 3001

`go run *.go`

---

### Frontend
> Make sure to have npm installed in local machine
1. Navigate to `/frontend` directory from root directory:

`cd /frontend`

2. Install the necessary dependencies
 
`npm install`

3. Run the application on localhost at port 3000

`npm start`


**Application preview**

![image](https://github.com/user-attachments/assets/4b18371f-e2bb-43e9-ad59-f4be3e3a88f4)
