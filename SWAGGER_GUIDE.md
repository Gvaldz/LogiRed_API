# 📚 Guía de Swagger - LogiRed API

## ✅ Problema Solucionado

**Antes:** Swagger no podía enviar el Authorization header con Bearer token
**Ahora:** Swagger está correctamente configurado para Bearer token authentication

---

## 🔧 Cambios Realizados

### 1. Corregido en `main.go`
```go
// ANTES (incorrecto):
// @securityDefinitions.apikey Bearer

// AHORA (correcto):
// @securityDefinitions.apiKey Bearer
```

### 2. Estandarizado en todos los controladores
```go
// TODOS los controladores ahora usan:
// @Security     Bearer

// En lugar de la mezcla anterior:
// @Security     ApiKeyAuth  ❌ (eliminado)
// @Security     Bearer      ✅ (consistente)
```

---

## 🚀 Cómo Usar Swagger Correctamente

### Paso 1: Abre Swagger UI
```
http://localhost:8080/swagger/index.html
```

### Paso 2: Login en Swagger

1. Scroll down to `POST /auth/login`
2. Click "Try it out"
3. Ingresa credenciales:
```json
{
  "email": "cliente@example.com",
  "password": "password123"
}
```
4. Click "Execute"

### Paso 3: Copia el Token del Header Authorization

**Response Headers:**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

✅ Copia el valor completo: `Bearer eyJhbGc...`

### Paso 4: Autoriza Swagger con el Token

**En la esquina superior derecha de Swagger UI:**
- Click en el botón 🔒 **Authorize**
- Pega el token completo:
```
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
- Click "Authorize"
- Click "Close"

### Paso 5: Usa cualquier endpoint protegido

Ahora puedes usar cualquier endpoint que requiera autenticación:
- `GET /rides/client/me`
- `POST /proposals`
- `GET /rides/available`
- etc.

El token se enviará automáticamente en el header Authorization

---

## 📸 Paso a Paso Visual

### 1. Login en Swagger
```
1. Busca POST /auth/login
2. Click "Try it out"
3. Ingresa email y password
4. Click "Execute"
5. Ves Response 200 OK
```

### 2. Copia Token del Header
```
En "Response headers":
Authorization: Bearer eyJhbGc...
↓
Copia todo: Bearer eyJhbGc...
```

### 3. Autoriza el Swagger
```
Click en 🔒 (arriba a la derecha)
Pega: Bearer eyJhbGc...
Click "Authorize"
Click "Close"
```

### 4. Usa endpoints protegidos
```
Ejemplo: GET /rides/client/me
Click "Try it out"
Click "Execute"
↓
Token se envía automáticamente
```

---

## ✅ Verificación

Después de realizar los pasos anteriores, deberías:

✅ Ver el ícono 🔒 en verde (significa autorizado)
✅ Poder usar endpoints sin errores de "token no autorizado"
✅ Ver respuestas correctas de los endpoints
✅ El token se envía automáticamente en cada request

---

## ❌ Errores Comunes en Swagger

### Error 1: "token de autorización requerido"
**Causa:** No autorizaste Swagger con el token
**Solución:**
1. Click 🔒 (Authorize button)
2. Pega el token: `Bearer eyJhbGc...`
3. Click "Authorize"

### Error 2: Token inválido
**Causa:** Token expirado o corrupto
**Solución:** Haz login de nuevo y obtén nuevo token

### Error 3: El token no se envía
**Causa:** Swagger no está autorizado
**Solución:**
1. Verifica que el ícono 🔒 esté en verde
2. Si no, click en 🔒 y autoriza de nuevo

---

## 🔍 Cómo Verificar que Funciona

### En Swagger:
1. **Después de autorizar** (🔒 en verde)
2. Abre cualquier endpoint protegido
3. Click "Try it out"
4. Click "Execute"
5. **Fíjate en "Request headers":**
```
Authorization: Bearer eyJhbGc...
```

✅ Si ves el Authorization header = está funcionando

---

## 📝 Endpoints Disponibles en Swagger

### Públicos (sin token necesario):
- `POST /auth/login` - Login
- `POST /auth/register` - Register (si existe)

### Protegidos (necesitan token):
- `GET /rides/client/me` - Ver mis viajes (Cliente)
- `POST /rides` - Crear viaje (Cliente)
- `GET /rides/available` - Ver viajes disponibles (Conductor aprobado)
- `POST /proposals` - Crear propuesta (Conductor aprobado)
- `GET /proposals/ride/{id}` - Ver propuestas (Conductor aprobado)
- etc.

---

## 🛠️ Información Técnica

**Configuración Swagger:**
- Token Type: **Bearer**
- Location: **Header**
- Header Name: **Authorization**
- Formato: **Bearer <token>**

**Controladores:**
- Todos usan: `@Security     Bearer`
- Consistente en todos los archivos
- Swagger reconoce automáticamente

---

## ✨ Resumen

| Acción | Cómo |
|--------|------|
| **Login** | POST /auth/login → obtén token |
| **Copiar token** | Del header Authorization |
| **Autorizar Swagger** | Click 🔒 → Paste token → Authorize |
| **Usar endpoints** | Click Try it out → Execute |
| **Token se envía** | Automático en el header Authorization |

✅ **¡Listo! Swagger ahora funciona correctamente con Bearer token authentication!**
