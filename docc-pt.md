# Configuração do Gateway (gateway.yaml)

Este arquivo define as rotas dos serviços, a porta do servidor e as configurações globais de CORS. Abaixo, você verá exemplos e explicações tanto em **Português** quanto em **Inglês**.

---

## 🇧🇷 Português

### 📍 Porta do Servidor

```yaml
server:
  port: 3001
```

Define a porta onde o gateway será iniciado.

---

### 🔁 Rotas dos Serviços

```yaml
routes:
  - name: users
    path: /users
    uri: http://localhost:3001

  - name: auth
    path: /auth
    uri: http://localhost:3000
```

Cada rota mapeia um caminho (`path`) para um serviço de destino (`uri`). Essas rotas serão automaticamente gerenciadas pelo proxy.

---

### 🌐 Configurações Globais de CORS

```yaml
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

#### Parâmetros:

* `allowedOrigins`: Lista de origens permitidas. Use `"*"` para permitir todas.
* `allowedMethods`: Métodos HTTP permitidos.
* `allowedHeaders`: Cabeçalhos permitidos.
* `allowCredentials`: Define se cookies/token de sessão devem ser permitidos (CORS com credenciais).

️ Atenção: usar `"*"` em `allowedOrigins` junto com `allowCredentials: true` pode causar **problemas de segurança** e **viola a especificação CORS**.

---

## 🇺🇸 English

### 📍 Server Port

```yaml
server:
  port: 3001
```

Defines the port on which the gateway will run.

---

### 🔁 Service Routes

```yaml
routes:
  - name: users
    path: /users
    uri: http://localhost:3001

  - name: auth
    path: /auth
    uri: http://localhost:3000
```

Each route maps a request `path` to a target service `uri`. These routes are automatically proxied by the system.

---

### 🌐 Global CORS Configuration

```yaml
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

#### Parameters:

* `allowedOrigins`: List of allowed origins. Use `"*"` to allow all.
* `allowedMethods`: Allowed HTTP methods.
* `allowedHeaders`: Allowed HTTP headers.
* `allowCredentials`: Whether to allow cookies/session tokens (CORS credentials).

️ Warning: Using `"*"` with `allowCredentials: true` can be a **security risk** and **violates the CORS specification**.

---

### ✅ Exemplo Completo / Full Example

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
