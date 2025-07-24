# Gateway API

This project is a **framework** in Go, focused on creating **API Gateways** and **Reverse Proxies**. It simplifies route management and network configurations, allowing you to define your gateway's behavior declaratively through a single YAML file.

## Key Features

- **Service Routing**: Easily configure the mapping of request paths to target service URIs.
- **Port Configuration**: Define the port on which the gateway will run.
- **Global CORS Configuration**: Centrally manage Cross-Origin Resource Sharing (CORS) policies.

## Route Configuration

All gateway routes and configurations are defined in the `gateway.yaml` file. You only need to configure this file for the gateway to function.

Example `gateway.yaml`:

```yaml
server:
  port: 3001

routes:
  - name: users
    path: /users
    uri: http://localhost:3001

  - name: auth
    path: /auth
    uri: http://localhost:3000

globalCors:
  allowedOrigins:
    - "http://localhost:3000"
    - "http://192.168.219.213:3000"
    - "http://localhost:3001"
  allowedMethods:
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
    - "OPTIONS"
  allowedHeaders:
    - "Content-Type"
    - "Authorization"
    - "x-api-key"
  allowCredentials: true
```

## How to Set Up the Project

To set up the project, you will need to have Go installed on your machine.

1.  **Clone the repository:**
    ```bash
    git clone <YOUR_REPOSITORY_URL>
    cd gateway-api
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

## How to Run the Project

1.  **Build the executable:**
    ```bash
    go build -o pulse-proxy .
    ```

2.  **Execute the gateway:**
    ```bash
    ./pulse-proxy
    ```
    Or, to run in the background (Linux/macOS):
    ```bash
    ./pulse-proxy &
    ```

## How to Use the Project

Once the gateway is running, it will start routing requests according to the configurations in `gateway.yaml`.

For example, if you configured a route like this:

```yaml
routes:
  - name: users
    path: /users
    uri: http://localhost:3001
```

Any request to `http://localhost:3001/users` (assuming the gateway port is 3001) will be forwarded to `http://localhost:3001/users` (the target service).

You can test the routes using tools like `curl` or Postman.

Example with `curl`:

```bash
curl http://localhost:3001/users
```

Ensure that the target services (`uri` in `gateway.yaml`) are running for the gateway to successfully forward requests.
