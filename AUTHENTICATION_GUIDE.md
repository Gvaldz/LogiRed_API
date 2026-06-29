# 🔐 Guía de Autenticación - LogiRed API

## 📋 Flujo de Login Actualizado

### 1️⃣ Cliente inicia sesión (POST /auth/login)

**Request:**
```bash
POST /auth/login
Content-Type: application/json

{
  "email": "cliente@example.com",
  "password": "password123"
}
```

**Response (200 OK):**
```json
{
  "expires_at": 4103020800,
  "message": "login exitoso. El token se encuentra en el header Authorization"
}
```

**Response Headers (IMPORTANTE):**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImNsaWVudGVAZXhhbXBsZS5jb20iLCJleHAiOjQxMDMwMjA4MDAsInVzZXJ0eXBlIjoxfQ.signature...
```

✅ **El token está en el header `Authorization`, NO en el body**

---

## 🔑 Cómo usar el Token en Futuras Requests

### Para Clientes (UserType = 1):

```bash
GET /rides/client/me
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

```bash
POST /rides
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Content-Type: application/json

{
  "origin_city": "Bogotá",
  "origin_address": "Calle 1 #2-3",
  ...
}
```

### Para Conductores (UserType = 2):

```bash
POST /proposals
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Content-Type: application/json

{
  "price": 150.50,
  "id_ride": 1,
  "comment": "Puedo llevar tu paquete"
}
```

---

## 📱 Ejemplo con cURL

### 1. Login:
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "cliente@example.com",
    "password": "password123"
  }' \
  -v
```

**Respuesta** (fíjate en los headers):
```
< HTTP/1.1 200 OK
< Authorization: Bearer eyJhbGc...
< Content-Type: application/json
<
{
  "expires_at": 4103020800,
  "message": "login exitoso..."
}
```

### 2. Usar el token en siguiente request:
```bash
curl -X GET http://localhost:8080/rides/client/me \
  -H "Authorization: Bearer eyJhbGc..." \
  -v
```

---

## 🔬 Estructura del Token JWT

El token contiene:
```json
{
  "user_id": 1,
  "email": "cliente@example.com",
  "exp": 4103020800,
  "usertype": 1
}
```

Para conductores, el token también contiene:
```json
{
  "user_id": 2,
  "email": "conductor@example.com",
  "exp": 4103020800,
  "usertype": 2,
  "approved": true    // ← Solo para conductores
}
```

---

## ✅ Checklist: Cómo Usar Correctamente

- [x] **Login**: POST /auth/login con email y password
- [x] **Extraer token**: Del header `Authorization: Bearer ...`
- [x] **Guardar token**: En memoria, localStorage, etc.
- [x] **Usar en requests**: Enviar header `Authorization: Bearer <token>`
- [x] **NO usar**: Token del body (no viene en el body)
- [x] **NO olvidar**: "Bearer " antes del token en el header

---

## ❌ Errores Comunes

### Error 1: "token de autorización requerido"
**Causa**: No envías el header Authorization
**Solución**: Asegúrate de incluir el header en CADA request:
```
Authorization: Bearer eyJhbGc...
```

### Error 2: "formato de token inválido"
**Causa**: Falta "Bearer " o formato incorrecto
**Solución**: El header debe ser exactamente:
```
Authorization: Bearer <token>
```
No:
```
Authorization: <token>                    ❌ Falta "Bearer"
Authorization: token <token>              ❌ Incorrecto
Authorization: JWT <token>                ❌ Incorrecto
```

### Error 3: "token inválido"
**Causa**: Token expirado o corrupto
**Solución**: Haz login de nuevo para obtener nuevo token

---

## 🧪 Ejemplo Completo con Postman

### Request 1: Login
```
POST http://localhost:8080/auth/login
Content-Type: application/json

{
  "email": "cliente@example.com",
  "password": "password123"
}
```

**Copiar el valor del header `Authorization` de la respuesta**

### Request 2: Ver mis viajes
```
GET http://localhost:8080/rides/client/me
Authorization: Bearer <pega_el_token_aquí>
```

---

## 🚀 Cliente (JavaScript/Frontend)

```javascript
// 1. Login
const response = await fetch('http://localhost:8080/auth/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    email: 'cliente@example.com',
    password: 'password123'
  })
});

// 2. Extraer token del header
const token = response.headers.get('Authorization');
// Resultado: "Bearer eyJhbGc..."

// 3. Guardar token
localStorage.setItem('token', token);

// 4. Usar en futuras requests
const ridesResponse = await fetch('http://localhost:8080/rides/client/me', {
  method: 'GET',
  headers: {
    'Authorization': localStorage.getItem('token')
  }
});
```

---

## 🔐 Resumen

| Aspecto | Detalle |
|---------|---------|
| **Token ubicación** | Header `Authorization` (NO body) |
| **Formato** | `Authorization: Bearer <token>` |
| **Qué hay en body del login** | `expires_at` y `message` |
| **Qué hay en token** | `user_id`, `email`, `exp`, `usertype`, `approved` (para conductores) |
| **Validación** | AuthMiddleware extrae datos del token |
| **Cada request necesita** | Header Authorization con token válido |

✅ **Ya todo está configurado. Solo recuerda:**
- ✅ Extraer token del header Authorization en el login
- ✅ Enviar token en header Authorization en futuras requests
- ✅ Usar formato `Bearer <token>`
