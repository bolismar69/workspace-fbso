package tech.fbso.ms_product_catalog_admin_simple.controller;

import jakarta.validation.Valid;
import java.util.List;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import tech.fbso.ms_product_catalog_admin_simple.dto.UnitMeasurementRequestDto;
import tech.fbso.ms_product_catalog_admin_simple.dto.UnitMeasurementResponseDto;
import tech.fbso.ms_product_catalog_admin_simple.service.UnitMeasurementService;

@RestController
@RequestMapping("/api/v1/unit-measurements")
public class UnitMeasurementController {

    private final UnitMeasurementService service;

    public UnitMeasurementController(UnitMeasurementService service) {
        this.service = service;
    }

    @PostMapping
    public ResponseEntity<UnitMeasurementResponseDto> create(@Valid @RequestBody UnitMeasurementRequestDto request) {
        UnitMeasurementResponseDto response = service.create(request);
        return ResponseEntity.status(HttpStatus.CREATED).body(response);
    }

    @GetMapping("/{id}")
    public ResponseEntity<UnitMeasurementResponseDto> findById(@PathVariable Long id) {
        UnitMeasurementResponseDto response = service.findById(id);
        return ResponseEntity.ok(response);
    }

    @GetMapping
    public ResponseEntity<List<UnitMeasurementResponseDto>> findAll() {
        List<UnitMeasurementResponseDto> response = service.findAll();
        return ResponseEntity.ok(response);
    }

    @PutMapping("/{id}")
    public ResponseEntity<UnitMeasurementResponseDto> update(
        @PathVariable Long id,
        @Valid @RequestBody UnitMeasurementRequestDto request
    ) {
        UnitMeasurementResponseDto response = service.update(id, request);
        return ResponseEntity.ok(response);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> delete(@PathVariable Long id) {
        service.delete(id);
        return ResponseEntity.noContent().build();
    }
}
