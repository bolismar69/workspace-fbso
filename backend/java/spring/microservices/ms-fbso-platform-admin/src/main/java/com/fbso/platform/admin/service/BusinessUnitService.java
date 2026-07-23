package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.request.BusinessUnitCreateRequest;
import com.fbso.platform.admin.dto.request.BusinessUnitUpdateRequest;
import com.fbso.platform.admin.dto.response.BusinessUnitResponse;
import com.fbso.platform.admin.entity.BusinessUnit;
import com.fbso.platform.admin.exception.BusinessException;
import com.fbso.platform.admin.exception.DuplicateCnpjException;
import com.fbso.platform.admin.exception.TenantNotFoundException;
import com.fbso.platform.admin.repository.BusinessUnitRepository;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.security.annotation.Auditable;
import com.fbso.platform.admin.utils.CnpjValidator;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

/**
 * Serviço de gestão de Unidades de Negócio (F04-05).
 *
 * <h3>RNs cobertas</h3>
 * <ul>
 *   <li>RN17-01: CNPJ único entre ativos do tenant, imutável, soft delete libera reúso</li>
 *   <li>RN17-02: BU desativada não pode ser "pai" de novas filiais</li>
 *   <li>RN17-03: Primeira BU = Matriz (isMatrix=true)</li>
 *   <li>RN17-04: Sem limite de níveis hierárquicos</li>
 *   <li>RN17-05: Seletor BU reflete permissões do usuário</li>
 * </ul>
 */
@Service
public class BusinessUnitService {

    private static final Logger log = LoggerFactory.getLogger(BusinessUnitService.class);

    private final BusinessUnitRepository buRepo;

    public BusinessUnitService(BusinessUnitRepository buRepo) {
        this.buRepo = buRepo;
    }

    /**
     * Cria uma nova Unidade de Negócio.
     *
     * <p>RN17-01: Valida CNPJ (formato + dígitos verificadores) e unicidade.
     * RN17-02: Se parentId informado, valida que a BU pai está ativa.</p>
     */
    @Auditable(entityType = "BUSINESS_UNIT", action = "CREATED")
    @Transactional
    public BusinessUnitResponse create(BusinessUnitCreateRequest req) {
        UUID tenantId = TenantContext.getTenantId();

        // RN17-01: Validar CNPJ
        if (!CnpjValidator.isValid(req.cnpj())) {
            throw new BusinessException("CNPJ inválido",
                    "O CNPJ informado não é válido. Verifique o formato e os dígitos verificadores.");
        }

        // RN17-01: CNPJ único entre ativos
        if (buRepo.existsByCnpj(req.cnpj(), tenantId)) {
            throw new DuplicateCnpjException(req.cnpj());
        }

        // RN17-02: Validar parent_id (se informado)
        if (req.parentId() != null) {
            validateParentActive(req.parentId());
        }

        BusinessUnit bu = new BusinessUnit();
        bu.setTenantId(tenantId);
        bu.setCnpj(req.cnpj());
        bu.setCorporateName(req.corporateName());
        bu.setTaxRegime(req.taxRegime());
        bu.setParentId(req.parentId());
        bu.setStreet(req.street());
        bu.setNumber(req.number());
        bu.setComplement(req.complement());
        bu.setNeighborhood(req.neighborhood());
        bu.setCity(req.city());
        bu.setState(req.state());
        bu.setZipCode(req.zipCode());
        bu.setStatus("ACTIVE");
        bu.setMatrix(req.isMatrix());

        buRepo.save(bu);
        log.info("BusinessUnit criada: id={}, isMatrix={}", bu.getId(), bu.isMatrix());
        return toResponse(bu);
    }

    /**
     * Atualiza dados cadastrais da BU.
     *
     * <p>RN17-01: CNPJ NÃO pode ser alterado (campo não incluso no request).
     * RN17-02: Se parentId alterado, valida que novo pai está ativo.</p>
     */
    @Auditable(entityType = "BUSINESS_UNIT", action = "UPDATED")
    @Transactional
    public BusinessUnitResponse update(UUID id, BusinessUnitUpdateRequest req) {
        BusinessUnit bu = buRepo.findById(id)
                .orElseThrow(() -> new TenantNotFoundException("Unidade de Negócio não encontrada: " + id));

        // Defesa em profundidade: tenant isolation já garantido pelo BaseRepository
        // + RLS. BU-scope authorization refinado em findAccessible() (RN17-05).

        // RN17-02: Validar novo parentId
        if (req.parentId() != null && !req.parentId().equals(bu.getParentId())) {
            validateParentActive(req.parentId());
            bu.setParentId(req.parentId());
        }

        if (req.corporateName() != null) bu.setCorporateName(req.corporateName());
        if (req.taxRegime() != null) bu.setTaxRegime(req.taxRegime());
        if (req.street() != null) bu.setStreet(req.street());
        if (req.number() != null) bu.setNumber(req.number());
        if (req.complement() != null) bu.setComplement(req.complement());
        if (req.neighborhood() != null) bu.setNeighborhood(req.neighborhood());
        if (req.city() != null) bu.setCity(req.city());
        if (req.state() != null) bu.setState(req.state());
        if (req.zipCode() != null) bu.setZipCode(req.zipCode());

        buRepo.update(bu);
        log.info("BusinessUnit atualizada: id={}", id);
        return toResponse(bu);
    }

    /**
     * Desativa uma BU (soft delete).
     *
     * <p>RN17-01: Soft delete libera CNPJ para reúso — índice parcial
     * {@code WHERE deleted_dt IS NULL} garante unicidade apenas entre ativos.</p>
     */
    @Auditable(entityType = "BUSINESS_UNIT", action = "DEACTIVATED")
    @Transactional
    public BusinessUnitResponse deactivate(UUID id) {
        BusinessUnit bu = buRepo.findById(id)
                .orElseThrow(() -> new TenantNotFoundException("Unidade de Negócio não encontrada: " + id));

        buRepo.softDelete(id, TenantContext.getUserId());
        log.info("BusinessUnit desativada: id={}", id);
        return toResponse(bu);
    }

    /**
     * Retorna BU por ID com validação de tenant (via BaseRepository).
     */
    public BusinessUnitResponse findById(UUID id) {
        BusinessUnit bu = buRepo.findById(id)
                .orElseThrow(() -> new TenantNotFoundException("Unidade de Negócio não encontrada: " + id));
        return toResponse(bu);
    }

    /**
     * Retorna árvore hierárquica completa do tenant (ADR-L08).
     */
    public List<BusinessUnitResponse> findTree() {
        UUID tenantId = TenantContext.getTenantId();
        List<BusinessUnit> all = buRepo.findTree(tenantId);
        return buildTree(all);
    }

    /**
     * Lista BUs acessíveis ao usuário (RN17-05).
     * Admin vê todas. Manager/Operator veem apenas as autorizadas.
     */
    public List<BusinessUnitResponse> findAccessible() {
        UUID tenantId = TenantContext.getTenantId();
        // Fase 0: Admin vê todas. Filtro por permissão será refinado.
        List<BusinessUnit> all = buRepo.findTree(tenantId);
        return all.stream().map(this::toResponse).toList();
    }

    // -- helpers --

    private void validateParentActive(UUID parentId) {
        buRepo.findById(parentId).ifPresentOrElse(parent -> {
            if (!"ACTIVE".equals(parent.getStatus())) {
                throw new BusinessException("Unidade pai inativa",
                        "Não é possível vincular a uma Unidade de Negócio desativada. RN17-02.");
            }
        }, () -> {
            throw new TenantNotFoundException("Unidade de Negócio pai não encontrada: " + parentId);
        });
    }

    private BusinessUnitResponse toResponse(BusinessUnit bu) {
        return new BusinessUnitResponse(
                bu.getId(), bu.getTenantId(), bu.getParentId(),
                bu.getCnpj(), bu.getCorporateName(), bu.getTaxRegime(),
                bu.getStreet(), bu.getNumber(), bu.getComplement(),
                bu.getNeighborhood(), bu.getCity(), bu.getState(), bu.getZipCode(),
                bu.getStatus(), bu.isMatrix(),
                null, // children preenchido via buildTree
                bu.getCreatedDt(), bu.getUpdatedDt()
        );
    }

    /**
     * Monta a árvore hierárquica a partir de uma lista plana.
     * Processa as BUs em ordem (raízes primeiro) e aninha filhas.
     */
    private List<BusinessUnitResponse> buildTree(List<BusinessUnit> all) {
        List<BusinessUnitResponse> roots = new ArrayList<>();
        for (BusinessUnit bu : all) {
            if (bu.getParentId() == null) {
                roots.add(toTreeResponse(bu, all));
            }
        }
        return roots;
    }

    private BusinessUnitResponse toTreeResponse(BusinessUnit bu, List<BusinessUnit> all) {
        List<BusinessUnitResponse> children = all.stream()
                .filter(c -> bu.getId().equals(c.getParentId()))
                .map(c -> toTreeResponse(c, all))
                .toList();
        return new BusinessUnitResponse(
                bu.getId(), bu.getTenantId(), bu.getParentId(),
                bu.getCnpj(), bu.getCorporateName(), bu.getTaxRegime(),
                bu.getStreet(), bu.getNumber(), bu.getComplement(),
                bu.getNeighborhood(), bu.getCity(), bu.getState(), bu.getZipCode(),
                bu.getStatus(), bu.isMatrix(),
                children.isEmpty() ? null : children,
                bu.getCreatedDt(), bu.getUpdatedDt()
        );
    }
}
