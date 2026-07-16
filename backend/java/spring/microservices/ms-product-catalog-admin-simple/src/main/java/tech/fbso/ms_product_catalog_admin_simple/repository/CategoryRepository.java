package tech.fbso.ms_product_catalog_admin_simple.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import tech.fbso.ms_product_catalog_admin_simple.entity.Category;

@Repository
public interface CategoryRepository extends JpaRepository<Category, Long> {
}
