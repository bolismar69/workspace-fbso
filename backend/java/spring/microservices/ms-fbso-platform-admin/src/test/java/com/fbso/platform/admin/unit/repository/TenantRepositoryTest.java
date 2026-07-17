package com.fbso.platform.admin.unit.repository;

import com.fbso.platform.admin.entity.Tenant;
import com.fbso.platform.admin.enums.TenantStatus;
import com.fbso.platform.admin.repository.TenantRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;

import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@DisplayName("TenantRepository")
class TenantRepositoryTest {

    @Mock private JdbcTemplate jdbc;

    private TenantRepository repo;

    @BeforeEach
    void setUp() {
        repo = new TenantRepository(jdbc);
    }

    @Nested
    @DisplayName("findAllPaginated")
    class FindAllPaginated {

        @Test
        @DisplayName("deve executar query com filtros de status e busca textual")
        void shouldFilterByStatusAndSearch() {
            when(jdbc.query(anyString(), any(RowMapper.class),
                    eq("ACTIVE"), eq("%Mercado%"), eq("%Mercado%"), eq(25), eq(0)))
                    .thenReturn(List.of());

            List<Tenant> result = repo.findAllPaginated(0, 25, "ACTIVE", null, "Mercado");

            assertThat(result).isEmpty();
        }

        @Test
        @DisplayName("não deve aplicar busca textual com menos de 3 caracteres")
        void shouldNotApplySearchWithLessThan3Chars() {
            when(jdbc.query(anyString(), any(RowMapper.class),
                    eq("ACTIVE"), eq(25), eq(0)))
                    .thenReturn(List.of());

            List<Tenant> result = repo.findAllPaginated(0, 25, "ACTIVE", null, "ab");

            assertThat(result).isEmpty();
        }

        @Test
        @DisplayName("deve incluir JOIN com plan quando filtro de plano informado")
        void shouldJoinPlanWhenFilteringByPlan() {
            when(jdbc.query(anyString(), any(RowMapper.class),
                    eq("Básico"), eq(25), eq(0)))
                    .thenReturn(List.of());

            List<Tenant> result = repo.findAllPaginated(0, 25, null, "Básico", null);

            assertThat(result).isEmpty();
        }
    }

    @Nested
    @DisplayName("countFiltered")
    class CountFiltered {

        @Test
        @DisplayName("deve retornar contagem com filtros")
        void shouldReturnFilteredCount() {
            when(jdbc.queryForObject(anyString(), eq(Integer.class),
                    eq("ACTIVE"), eq("Básico"),
                    eq("%Busca%"), eq("%Busca%")))
                    .thenReturn(15);

            int count = repo.countFiltered("ACTIVE", "Básico", "Busca");

            assertThat(count).isEqualTo(15);
        }
    }

    @Nested
    @DisplayName("findByNameCorporate")
    class FindByName {

        @Test
        @DisplayName("deve buscar tenant por razão social exata")
        void shouldFindByNameCorporate() {
            Tenant tenant = new Tenant();
            tenant.setNameCorporate("Empresa Teste");
            tenant.setStatus(TenantStatus.ACTIVE);

            when(jdbc.query(anyString(), any(RowMapper.class), eq("Empresa Teste")))
                    .thenReturn(List.of(tenant));

            var result = repo.findByNameCorporate("Empresa Teste");

            assertThat(result).isPresent();
            assertThat(result.get().getNameCorporate()).isEqualTo("Empresa Teste");
        }

        @Test
        @DisplayName("deve retornar empty quando não encontrado")
        void shouldReturnEmptyWhenNotFound() {
            when(jdbc.query(anyString(), any(RowMapper.class), eq("Inexistente")))
                    .thenReturn(List.of());

            var result = repo.findByNameCorporate("Inexistente");

            assertThat(result).isEmpty();
        }
    }
}
