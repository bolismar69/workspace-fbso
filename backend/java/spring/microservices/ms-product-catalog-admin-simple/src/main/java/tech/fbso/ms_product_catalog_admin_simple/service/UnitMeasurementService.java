package tech.fbso.ms_product_catalog_admin_simple.service;

import java.util.List;
import java.util.stream.Collectors;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import tech.fbso.ms_product_catalog_admin_simple.dto.UnitMeasurementRequestDto;
import tech.fbso.ms_product_catalog_admin_simple.dto.UnitMeasurementResponseDto;
import tech.fbso.ms_product_catalog_admin_simple.entity.UnitMeasurement;
import tech.fbso.ms_product_catalog_admin_simple.exception.ResourceNotFoundException;
import tech.fbso.ms_product_catalog_admin_simple.repository.UnitMeasurementRepository;

@Service
@Transactional
public class UnitMeasurementService {

    private final UnitMeasurementRepository repository;

    public UnitMeasurementService(UnitMeasurementRepository repository) {
        this.repository = repository;
    }

    public UnitMeasurementResponseDto create(UnitMeasurementRequestDto request) {
        UnitMeasurement unitMeasurement = new UnitMeasurement();
        unitMeasurement.setAcronym(request.acronym());
        unitMeasurement.setDescription(request.description());

        UnitMeasurement savedUnit = repository.save(unitMeasurement);
        return toUnitMeasurementResponse(savedUnit);
    }

    @Transactional(readOnly = true)
    public UnitMeasurementResponseDto findById(Long id) {
        UnitMeasurement unitMeasurement = repository.findById(id)
            .orElseThrow(() -> new ResourceNotFoundException("UnitMeasurement not found with id: " + id));
        return toUnitMeasurementResponse(unitMeasurement);
    }

    @Transactional(readOnly = true)
    public List<UnitMeasurementResponseDto> findAll() {
        return repository.findAll().stream()
            .map(this::toUnitMeasurementResponse)
            .collect(Collectors.toList());
    }

    public UnitMeasurementResponseDto update(Long id, UnitMeasurementRequestDto request) {
        UnitMeasurement unitMeasurement = repository.findById(id)
            .orElseThrow(() -> new ResourceNotFoundException("UnitMeasurement not found with id: " + id));

        unitMeasurement.setAcronym(request.acronym());
        unitMeasurement.setDescription(request.description());

        UnitMeasurement updatedUnit = repository.save(unitMeasurement);
        return toUnitMeasurementResponse(updatedUnit);
    }

    public void delete(Long id) {
        if (!repository.existsById(id)) {
            throw new ResourceNotFoundException("UnitMeasurement not found with id: " + id);
        }
        repository.deleteById(id);
    }

    private UnitMeasurementResponseDto toUnitMeasurementResponse(UnitMeasurement unitMeasurement) {
        return new UnitMeasurementResponseDto(
            unitMeasurement.getId(),
            unitMeasurement.getAcronym(),
            unitMeasurement.getDescription(),
            unitMeasurement.getCreatedAt(),
            unitMeasurement.getUpdatedAt()
        );
    }
}
