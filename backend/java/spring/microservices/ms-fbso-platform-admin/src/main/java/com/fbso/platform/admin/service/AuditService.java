package com.fbso.platform.admin.service;

import com.fbso.platform.admin.dto.response.AuditEntryResponse;
import com.fbso.platform.admin.repository.AuditRepository;
import org.springframework.stereotype.Service;

import java.time.OffsetDateTime;
import java.util.List;

/**
 * Serviço de consulta de auditoria (F02-05).
 *
 * <p>Registros são imutáveis — apenas leitura.</p>
 */
@Service
public class AuditService {

    private final AuditRepository auditRepo;

    public AuditService(AuditRepository auditRepo) {
        this.auditRepo = auditRepo;
    }

    /**
     * Consulta registros de auditoria com filtros e paginação.
     *
     * @param startDate  data inicial (opcional)
     * @param endDate    data final (opcional)
     * @param action     ação (opcional)
     * @param entityType tipo de entidade (opcional)
     * @param page       página (0-based)
     * @param size       registros por página (padrão 25, max 100)
     * @return lista de registros de auditoria
     */
    public List<AuditEntryResponse> search(String startDate, String endDate,
                                            String action, String entityType,
                                            int page, int size) {
        OffsetDateTime start = startDate != null && !startDate.isBlank()
                ? OffsetDateTime.parse(startDate) : null;
        OffsetDateTime end = endDate != null && !endDate.isBlank()
                ? OffsetDateTime.parse(endDate) : null;

        return auditRepo.findByFilters(start, end, action, entityType, page, size)
                .stream().map(AuditEntryResponse::from).toList();
    }
}
