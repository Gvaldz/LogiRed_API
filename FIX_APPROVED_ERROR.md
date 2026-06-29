# 🔧 FIX: "error al verificar aprobación" - Solución Completa

## 🐛 Problema Reportado

**Error recibido:**
```json
{
  "error": "error al verificar aprobación"
}
```

**Token decodificado mostraba:**
```json
{
  "approved": true,
  "email": "atun@gmail.com",
  "exp": 4936098861,
  "user_id": 27,
  "usertype": 2
}
```

**Contradicción:** El token tenía `approved: true` pero la API devolvía error de aprobación.

---

## 🔍 Causa Raíz

El `AuthMiddleware` NO estaba extrayendo `usertype` ni `approved` del token JWT. Esto causaba que:

1. El token se validaba correctamente ✅
2. Pero `usertype` y `approved` NO se guardaban en el contexto ❌
3. Las middlewares que esperaban estos valores en el contexto fallaban ❌

**Flujo Incorrecto:**
```
Token: {approved: true, usertype: 2, ...}
    ↓
AuthMiddleware valida token pero NO extrae approved
    ↓
c.Get("approved") = no existe en contexto
    ↓
driver_approval middleware: "error al verificar aprobación"
```

---

## ✅ Solución Implementada

### 1. Agregados métodos de extracción a `token_service.go`

```go
// ExtractUserType - Lee usertype del token
func (s *JWTService) ExtractUserType(tokenString string) (int, error) {
    // Parse token
    // Extrae claims["usertype"]
    // Retorna el valor
}

// ExtractApproved - Lee approved del token
func (s *JWTService) ExtractApproved(tokenString string) (bool, error) {
    // Parse token
    // Extrae claims["approved"]
    // Retorna el valor (o false si no existe)
}
```

### 2. Actualizado `TokenService` interface

```go
type TokenService interface {
    GenerateToken(userID int32, email string, usertype int, approved bool) (Token, error)
    ValidateToken(tokenString string) (int32, error)
    ExtractUserType(tokenString string) (int, error)      // ← NUEVO
    ExtractApproved(tokenString string) (bool, error)     // ← NUEVO
}
```

### 3. Actualizado `AuthMiddleware` en `middleware.go`

```go
func AuthMiddleware(tokenService tokenService.TokenService, userRepo users_domain.UserRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        // ... validación del token ...
        
        // NUEVO: Extrae usertype del token
        userType, err := tokenService.ExtractUserType(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{...})
            return
        }

        // NUEVO: Extrae approved del token
        approved, err := tokenService.ExtractApproved(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{...})
            return
        }

        // NUEVO: Guarda en contexto
        c.Set("userType", userType)     // ← Ahora disponible para otras middlewares
        c.Set("approved", approved)     // ← Ahora disponible para driver_approval
        c.Next()
    }
}
```

---

## 🔄 Flujo Correcto (Después del Fix)

```
Token JWT: {approved: true, usertype: 2, user_id: 27, ...}
    ↓
AuthMiddleware.ExtractUserType(token) → 2
AuthMiddleware.ExtractApproved(token) → true
    ↓
c.Set("userType", 2)
c.Set("approved", true)
    ↓
driver_approval middleware:
    approved := c.Get("approved") → true ✅
    Valida exitosamente ✅
    ↓
Acceso permitido ✅
```

---

## ✨ Cambios Realizados

| Archivo | Cambio |
|---------|--------|
| `src/core/token_service.go` | + 2 métodos de extracción |
| `src/core/services/auth/domain/tokenservice.go` | + 2 métodos en interfaz |
| `src/server/middleware/middleware.go` | Extrae usertype y approved del token |

---

## 🧪 Verificación

Después de estos cambios:

✅ **Token generado correctamente:**
```json
{
  "user_id": 27,
  "email": "atun@gmail.com",
  "exp": 4936098861,
  "usertype": 2,
  "approved": true
}
```

✅ **AuthMiddleware extrae y guarda en contexto:**
```
c.Get("userType") → 2
c.Get("approved") → true
c.Get("user") → User object
c.Get("userID") → 27
```

✅ **Middlewares pueden acceder a los valores:**
```go
approved, _ := c.Get("approved")
if !approved.(bool) {
    c.AbortWithStatusJSON(403, gin.H{"error": "no aprobado"})
}
```

✅ **¡Error "error al verificar aprobación" RESUELTO!**

---

## 📝 Resumen de Funcionamiento

1. **Login:** Usuario se autentica → Recibe JWT con `approved`
2. **Almacenamiento:** Cliente guarda token
3. **Request:** Envía `Authorization: Bearer <token>`
4. **AuthMiddleware:** 
   - Valida token ✅
   - Extrae `usertype` ✅
   - Extrae `approved` ✅
   - Guarda en contexto ✅
5. **Driver Approval Middleware:** 
   - Lee `approved` del contexto ✅
   - Valida acceso ✅
6. **Response:** Permite acceso si está aprobado ✅

---

## 🚀 Cómo Probar

1. **Login con conductor aprobado:**
   ```bash
   POST /auth/login
   {email, password}
   ```

2. **Usar endpoint protegido:**
   ```bash
   GET /rides/available
   Authorization: Bearer <token_del_login>
   ```

3. **Esperado:**
   - ✅ Sin error "error al verificar aprobación"
   - ✅ Respuesta normal del endpoint

---

## 💾 Commit

```
408baf2 - fix: restore token extraction methods and update AuthMiddleware
```

---

**¡Problema solucionado! ✅**
