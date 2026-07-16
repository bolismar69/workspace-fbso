package tech.fbso.ms_product_catalog_admin_simple.service;

import java.util.List;
import java.util.stream.Collectors;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import tech.fbso.ms_product_catalog_admin_simple.dto.CategoryRequestDto;
import tech.fbso.ms_product_catalog_admin_simple.dto.CategoryResponseDto;
import tech.fbso.ms_product_catalog_admin_simple.entity.Category;
import tech.fbso.ms_product_catalog_admin_simple.exception.ResourceNotFoundException;
import tech.fbso.ms_product_catalog_admin_simple.repository.CategoryRepository;

@Service
@Transactional
public class CategoryService {

    private final CategoryRepository repository;

    public CategoryService(CategoryRepository repository) {
        this.repository = repository;
    }

    public CategoryResponseDto create(CategoryRequestDto request) {
        Category category = new Category();
        category.setAcronym(request.acronym());
        category.setDescription(request.description());

        Category savedCategory = repository.save(category);
        return toCategoryResponse(savedCategory);
    }

    @Transactional(readOnly = true)
    public CategoryResponseDto findById(Long id) {
        Category category = repository.findById(id)
            .orElseThrow(() -> new ResourceNotFoundException("Category not found with id: " + id));
        return toCategoryResponse(category);
    }

    @Transactional(readOnly = true)
    public List<CategoryResponseDto> findAll() {
        return repository.findAll().stream()
            .map(this::toCategoryResponse)
            .collect(Collectors.toList());
    }

    public CategoryResponseDto update(Long id, CategoryRequestDto request) {
        Category category = repository.findById(id)
            .orElseThrow(() -> new ResourceNotFoundException("Category not found with id: " + id));

        category.setAcronym(request.acronym());
        category.setDescription(request.description());

        Category updatedCategory = repository.save(category);
        return toCategoryResponse(updatedCategory);
    }

    public void delete(Long id) {
        if (!repository.existsById(id)) {
            throw new ResourceNotFoundException("Category not found with id: " + id);
        }
        repository.deleteById(id);
    }

    private CategoryResponseDto toCategoryResponse(Category category) {
        return new CategoryResponseDto(
            category.getId(),
            category.getAcronym(),
            category.getDescription(),
            category.getCreatedAt(),
            category.getUpdatedAt()
        );
    }
}
