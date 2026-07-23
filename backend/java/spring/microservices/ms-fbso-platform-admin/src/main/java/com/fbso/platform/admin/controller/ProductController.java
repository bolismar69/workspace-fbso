package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.request.ProductCreateRequest;
import com.fbso.platform.admin.dto.request.ProductUpdateRequest;
import com.fbso.platform.admin.dto.response.ProductResponse;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import com.fbso.platform.admin.service.ProductService;
import jakarta.validation.Valid;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

/**
 * Controller REST para Catálogo de Produtos/Serviços (F04-06).
 *
 * <h3>RBAC (RN10-01)</h3>
 * <ul>
 *   <li>ADMIN_TENANT, MANAGER_BU: create, edit, deactivate</li>
 *   <li>OPERATOR_BU: view</li>
 * </ul>
 */
@RestController
@RequestMapping("/api/v1/products")
public class ProductController {

    private final ProductService productService;

    public ProductController(ProductService productService) {
        this.productService = productService;
    }

    @GetMapping
    @RequiresPermission(resource = "PRODUCT_SERVICE", action = "view")
    public ResponseEntity<List<ProductResponse>> list(
            @RequestParam(required = false) UUID business_unit_id) {
        return ResponseEntity.ok(productService.findAll(business_unit_id));
    }

    @GetMapping("/{id}")
    @RequiresPermission(resource = "PRODUCT_SERVICE", action = "view")
    public ResponseEntity<ProductResponse> getById(@PathVariable UUID id) {
        return ResponseEntity.ok(productService.findById(id));
    }

    @PostMapping
    @RequiresPermission(resource = "PRODUCT_SERVICE", action = "create")
    public ResponseEntity<ProductResponse> create(
            @Valid @RequestBody ProductCreateRequest req) {
        ProductResponse response = productService.create(req);
        return ResponseEntity.status(HttpStatus.CREATED).body(response);
    }

    @PatchMapping("/{id}")
    @RequiresPermission(resource = "PRODUCT_SERVICE", action = "edit")
    public ResponseEntity<ProductResponse> update(
            @PathVariable UUID id,
            @Valid @RequestBody ProductUpdateRequest req) {
        return ResponseEntity.ok(productService.update(id, req));
    }

    @PostMapping("/{id}/deactivate")
    @RequiresPermission(resource = "PRODUCT_SERVICE", action = "delete")
    public ResponseEntity<ProductResponse> deactivate(@PathVariable UUID id) {
        return ResponseEntity.ok(productService.deactivate(id));
    }
}
