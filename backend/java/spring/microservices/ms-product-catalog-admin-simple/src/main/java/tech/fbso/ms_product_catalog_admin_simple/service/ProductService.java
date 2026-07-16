package tech.fbso.ms_product_catalog_admin_simple.service;

import java.util.List;
import java.util.stream.Collectors;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import tech.fbso.ms_product_catalog_admin_simple.dto.ProductRequestDto;
import tech.fbso.ms_product_catalog_admin_simple.dto.ProductResponseDto;
import tech.fbso.ms_product_catalog_admin_simple.entity.Category;
import tech.fbso.ms_product_catalog_admin_simple.entity.Product;
import tech.fbso.ms_product_catalog_admin_simple.entity.UnitMeasurement;
import tech.fbso.ms_product_catalog_admin_simple.exception.ResourceNotFoundException;
import tech.fbso.ms_product_catalog_admin_simple.repository.CategoryRepository;
import tech.fbso.ms_product_catalog_admin_simple.repository.ProductRepository;
import tech.fbso.ms_product_catalog_admin_simple.repository.UnitMeasurementRepository;

@Service
@Transactional
public class ProductService {

    private final ProductRepository productRepository;
    private final CategoryRepository categoryRepository;
    private final UnitMeasurementRepository unitMeasurementRepository;

    public ProductService(
        ProductRepository productRepository,
        CategoryRepository categoryRepository,
        UnitMeasurementRepository unitMeasurementRepository
    ) {
        this.productRepository = productRepository;
        this.categoryRepository = categoryRepository;
        this.unitMeasurementRepository = unitMeasurementRepository;
    }

    public ProductResponseDto create(ProductRequestDto request) {
        Category category = categoryRepository.findById(request.categoryId())
            .orElseThrow(() -> new ResourceNotFoundException("Category not found with id: " + request.categoryId()));

        UnitMeasurement unitMeasurement = unitMeasurementRepository.findById(request.unitMeasurementId())
            .orElseThrow(() -> new ResourceNotFoundException("UnitMeasurement not found with id: " + request.unitMeasurementId()));

        Product product = new Product();
        product.setName(request.name());
        product.setBarcode(request.barcode());
        product.setPrice(request.price());
        product.setQuantityInStock(request.quantityInStock());
        product.setCategory(category);
        product.setUnitMeasurement(unitMeasurement);

        Product savedProduct = productRepository.save(product);
        return toProductResponse(savedProduct);
    }

    @Transactional(readOnly = true)
    public ProductResponseDto findById(Long id) {
        Product product = productRepository.findById(id)
            .orElseThrow(() -> new ResourceNotFoundException("Product not found with id: " + id));
        return toProductResponse(product);
    }

    @Transactional(readOnly = true)
    public List<ProductResponseDto> findAll() {
        return productRepository.findAll().stream()
            .map(this::toProductResponse)
            .collect(Collectors.toList());
    }

    public ProductResponseDto update(Long id, ProductRequestDto request) {
        Product product = productRepository.findById(id)
            .orElseThrow(() -> new ResourceNotFoundException("Product not found with id: " + id));

        Category category = categoryRepository.findById(request.categoryId())
            .orElseThrow(() -> new ResourceNotFoundException("Category not found with id: " + request.categoryId()));

        UnitMeasurement unitMeasurement = unitMeasurementRepository.findById(request.unitMeasurementId())
            .orElseThrow(() -> new ResourceNotFoundException("UnitMeasurement not found with id: " + request.unitMeasurementId()));

        product.setName(request.name());
        product.setBarcode(request.barcode());
        product.setPrice(request.price());
        product.setQuantityInStock(request.quantityInStock());
        product.setCategory(category);
        product.setUnitMeasurement(unitMeasurement);

        Product updatedProduct = productRepository.save(product);
        return toProductResponse(updatedProduct);
    }

    public void delete(Long id) {
        if (!productRepository.existsById(id)) {
            throw new ResourceNotFoundException("Product not found with id: " + id);
        }
        productRepository.deleteById(id);
    }

    private ProductResponseDto toProductResponse(Product product) {
        return new ProductResponseDto(
            product.getId(),
            product.getSku(),
            product.getBarcode(),
            product.getName(),
            product.getPrice(),
            product.getQuantityInStock(),
            product.getCategory().getId(),
            product.getCategory().getDescription(),
            product.getUnitMeasurement().getId(),
            product.getUnitMeasurement().getAcronym(),
            product.getCreatedAt(),
            product.getUpdatedAt()
        );
    }
}
