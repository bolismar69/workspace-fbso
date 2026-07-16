# Microservice : ms-product-catalog-admin-service

## 📋 Resumo do que foi criado:

Endpoints criados:

```
POST   /api/v1/categories                    (201 Created)
GET    /api/v1/categories                    (200 OK)
GET    /api/v1/categories/{id}               (200 OK)
PUT    /api/v1/categories/{id}               (200 OK)
DELETE /api/v1/categories/{id}               (204 No Content)

POST   /api/v1/products                      
GET    /api/v1/products
GET    /api/v1/products/{id}
PUT    /api/v1/products/{id}
DELETE /api/v1/products/{id}

POST   /api/v1/unit-measurements
GET    /api/v1/unit-measurements
GET    /api/v1/unit-measurements/{id}
PUT    /api/v1/unit-measurements/{id}
DELETE /api/v1/unit-measurements/{id}
```

### **Exception Handling**
- `ResourceNotFoundException` - para recursos não encontrados (404)
- `GlobalExceptionHandler` - com suporte para validação, 404 e erros genéricos

## ✨ Best Practices implementados:
- ✅ Validação com annotations (`@NotBlank`, `@Size`, `@Positive`, `@DecimalMin`)
- ✅ `@Transactional` para consistência de dados
- ✅ Constructor Injection 
- ✅ ProblemDetail para respostas de erro padronizadas
- ✅ HTTP status codes corretos
- ✅ Read-only transactions para queries
- ✅ Tratamento de Foreign Keys com erro apropriado

---
