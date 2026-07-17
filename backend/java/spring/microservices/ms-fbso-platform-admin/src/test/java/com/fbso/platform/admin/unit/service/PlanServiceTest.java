package com.fbso.platform.admin.unit.service;

import com.fbso.platform.admin.dto.request.PlanCreateRequest;
import com.fbso.platform.admin.dto.request.PlanUpdateRequest;
import com.fbso.platform.admin.entity.Plan;
import com.fbso.platform.admin.enums.Recurrence;
import com.fbso.platform.admin.exception.BusinessException;
import com.fbso.platform.admin.exception.PlanHasActiveSubscribersException;
import com.fbso.platform.admin.repository.PlanRepository;
import com.fbso.platform.admin.service.PlanService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.math.BigDecimal;
import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
@org.mockito.junit.jupiter.MockitoSettings(strictness = org.mockito.quality.Strictness.LENIENT)
@DisplayName("PlanService")
class PlanServiceTest {

    @Mock private PlanRepository repo;
    private PlanService service;

    @BeforeEach
    void setUp() {
        service = new PlanService(repo);
    }

    @Nested
    @DisplayName("create")
    class Create {

        @Test
        @DisplayName("deve criar plano com version=1 e status ACTIVE")
        void shouldCreate() {
            doNothing().when(repo).save(any());

            var req = new PlanCreateRequest("Básico", "Plano básico",
                    new BigDecimal("99.00"), Recurrence.MONTHLY);
            var result = service.create(req);

            assertThat(result.name()).isEqualTo("Básico");
            assertThat(result.status()).isEqualTo("ACTIVE");
            assertThat(result.version()).isEqualTo(1);
        }
    }

    @Nested
    @DisplayName("update")
    class Update {

        @Test
        @DisplayName("deve versionar quando preço altera")
        void shouldVersionOnPriceChange() {
            UUID id = UUID.randomUUID();
            Plan existing = new Plan();
            existing.setId(id);
            existing.setName("Básico");
            existing.setPrice(new BigDecimal("99.00"));
            existing.setStatus("ACTIVE");
            existing.setVersion(1);

            when(repo.findById(id)).thenReturn(Optional.of(existing));
            doNothing().when(repo).update(any());

            var req = new PlanUpdateRequest(null, null, new BigDecimal("149.00"), null);
            var result = service.update(id, req);

            assertThat(result.version()).isEqualTo(2);
        }
    }

    @Nested
    @DisplayName("deactivate")
    class Deactivate {

        @Test
        @DisplayName("deve desativar plano sem assinantes")
        void shouldDeactivateWithNoSubscribers() {
            UUID id = UUID.randomUUID();
            Plan plan = new Plan();
            plan.setId(id);
            plan.setName("Básico");
            plan.setStatus("ACTIVE");

            when(repo.findById(id)).thenReturn(Optional.of(plan));
            when(repo.hasActiveSubscribers(id)).thenReturn(false);
            when(repo.countActive()).thenReturn(2);
            doNothing().when(repo).update(any());

            var result = service.deactivate(id);

            assertThat(result.status()).isEqualTo("DISCONTINUED");
        }

        @Test
        @DisplayName("deve lançar PlanHasActiveSubscribersException quando tem assinantes (RN06-01)")
        void shouldThrowWithActiveSubscribers() {
            UUID id = UUID.randomUUID();
            Plan plan = new Plan();
            plan.setId(id);
            plan.setName("Enterprise");
            plan.setStatus("ACTIVE");

            when(repo.findById(id)).thenReturn(Optional.of(plan));
            when(repo.hasActiveSubscribers(id)).thenReturn(true);

            assertThatThrownBy(() -> service.deactivate(id))
                    .isInstanceOf(PlanHasActiveSubscribersException.class);
        }
    }
}
