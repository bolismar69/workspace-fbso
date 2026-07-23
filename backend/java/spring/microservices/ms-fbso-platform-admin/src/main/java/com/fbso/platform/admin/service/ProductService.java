package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.request.ProductCreateRequest;
import com.fbso.platform.admin.dto.request.ProductUpdateRequest;
import com.fbso.platform.admin.dto.response.ProductResponse;
import com.fbso.platform.admin.entity.BusinessUnit;
import com.fbso.platform.admin.enums.ProductType;
import com.fbso.platform.admin.exception.BusinessException;
import com.fbso.platform.admin.exception.TenantNotFoundException;
import com.fbso.platform.admin.repository.BusinessUnitRepository;
import com.fbso.platform.admin.repository.ProductRepository;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.security.annotation.Auditable;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.UUID;

/**
 * Serviço de catálogo de Produtos/Serviços (F04-06).
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN18-01: Catálogo segmentado por Unidade de Negócio</li>
 *   <li>RN18-02: SKU opcional, único por BU ativo (índice parcial)</li>
 *   <li>RN18-03: Indicador "Não mapeado" (placeholder fiscal)</li>
 *   <li>RN18-04: Soft delete em produtos</li>
 * </ul>
 */
@Service
public class ProductService {

    private static final Logger log = LoggerFactory.getLogger(ProductService.class);

    private final ProductRepository productRepo;
    private final BusinessUnitRepository buRepo;

    public ProductService(ProductRepository productRepo, BusinessUnitRepository buRepo) {
        this.productRepo = productRepo;
        this.buRepo = buRepo;
    }

    /**
     * Cria um produto vinculado à BU ativa (RN18-01).
     * RN18-02: SKU se informado, deve ser único na BU.
     */
    @Auditable(entityType = "PRODUCT_SERVICE", action = "CREATED")
    @Transactional
    public ProductResponse create(ProductCreateRequest req) {
        // Validar que BU existe e está ativa. Tenant isolation: BaseRepository
        // filtra por tenant_id + RLS. Se BU de outro tenant → 404.
        BusinessUnit bu = buRepo.findById(req.businessUnitId())
                .orElseThrow(() -> new TenantNotFoundException("Unidade de Negócio não encontrada"));

        // RN18-02: SKU único por BU
        if (req.sku() != null && !req.sku().isBlank()) {
            if (productRepo.existsBySku(req.sku(), req.businessUnitId())) {
                throw new BusinessException("SKU duplicado",
                        "Já existe um produto com o SKU '" + req.sku() + "' nesta Unidade de Negócio.");
            }
        }

        com.fbso.platform.admin.entity.ProductService ps =
                new com.fbso.platform.admin.entity.ProductService();
        ps.setTenantId(bu.getTenantId());
        ps.setBusinessUnitId(req.businessUnitId());
        ps.setName(req.name());
        ps.setSku(req.sku());
        ps.setType(ProductType.valueOf(req.type()));
        ps.setDescription(req.description());
        ps.setStatus("ACTIVE");

        productRepo.save(ps);
        log.info("Produto criado: id={}, name={}, buId={}", ps.getId(), ps.getName(), ps.getBusinessUnitId());
        return toResponse(ps);
    }

    @Auditable(entityType = "PRODUCT_SERVICE", action = "UPDATED")
    @Transactional
    public ProductResponse update(UUID id, ProductUpdateRequest req) {
        com.fbso.platform.admin.entity.ProductService ps = productRepo.findById(id)
                .orElseThrow(() -> new TenantNotFoundException("Produto não encontrado: " + id));

        // RN18-02: Validar SKU se alterado
        if (req.sku() != null && !req.sku().isBlank() && !req.sku().equals(ps.getSku())) {
            if (productRepo.existsBySku(req.sku(), ps.getBusinessUnitId())) {
                throw new BusinessException("SKU duplicado",
                        "Já existe um produto com o SKU '" + req.sku() + "' nesta Unidade de Negócio.");
            }
            ps.setSku(req.sku());
        }

        if (req.name() != null) ps.setName(req.name());
        if (req.type() != null) ps.setType(ProductType.valueOf(req.type()));
        if (req.description() != null) ps.setDescription(req.description());

        productRepo.update(ps);
        log.info("Produto atualizado: id={}", id);
        return toResponse(ps);
    }

    /**
     * Soft delete de produto (RN18-04).
     */
    @Auditable(entityType = "PRODUCT_SERVICE", action = "DEACTIVATED")
    @Transactional
    public ProductResponse deactivate(UUID id) {
        com.fbso.platform.admin.entity.ProductService ps = productRepo.findById(id)
                .orElseThrow(() -> new TenantNotFoundException("Produto não encontrado: " + id));

        productRepo.softDelete(id, TenantContext.getUserId());
        log.info("Produto desativado: id={}, name={}", id, ps.getName());
        return toResponse(ps);
    }

    public ProductResponse findById(UUID id) {
        com.fbso.platform.admin.entity.ProductService ps = productRepo.findById(id)
                .orElseThrow(() -> new TenantNotFoundException("Produto não encontrado: " + id));
        return toResponse(ps);
    }

    /**
     * Lista produtos (RN18-01 — segmentado por BU se informado).
     */
    public List<ProductResponse> findAll(UUID businessUnitId) {
        List<com.fbso.platform.admin.entity.ProductService> products;
        if (businessUnitId != null) {
            products = productRepo.findByBusinessUnit(businessUnitId);
        } else {
            products = productRepo.findAll(0, 100, "name");
        }
        return products.stream().map(this::toResponse).toList();
    }

    private ProductResponse toResponse(com.fbso.platform.admin.entity.ProductService ps) {
        return new ProductResponse(
                ps.getId(), ps.getTenantId(), ps.getBusinessUnitId(),
                ps.getName(), ps.getSku(),
                ps.getType() != null ? ps.getType().name() : null,
                ps.getDescription(), ps.getStatus(),
                "NOT_MAPPED", // RN18-03: placeholder fiscal
                ps.getCreatedDt(), ps.getUpdatedDt()
        );
    }
}
