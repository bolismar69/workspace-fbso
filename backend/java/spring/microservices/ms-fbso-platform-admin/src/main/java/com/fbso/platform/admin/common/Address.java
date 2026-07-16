package com.fbso.platform.admin.common;

import java.util.Objects;

/**
 * Value Object representando um endereço.
 * <p>
 * Imutável — uma vez construído, não pode ser alterado.
 * Usado como componente embutido em entidades que possuem endereço
 * (ex: {@code BusinessUnit}).
 */
public final class Address {

    private final String street;
    private final String number;
    private final String complement;
    private final String neighborhood;
    private final String city;
    private final String state;
    private final String zipCode;

    private Address(Builder builder) {
        this.street = builder.street;
        this.number = builder.number;
        this.complement = builder.complement;
        this.neighborhood = builder.neighborhood;
        this.city = builder.city;
        this.state = builder.state;
        this.zipCode = builder.zipCode;
    }

    // ---- Getters ----

    public String getStreet() { return street; }
    public String getNumber() { return number; }
    public String getComplement() { return complement; }
    public String getNeighborhood() { return neighborhood; }
    public String getCity() { return city; }
    public String getState() { return state; }
    public String getZipCode() { return zipCode; }

    // ---- Builder ----

    public static Builder builder() {
        return new Builder();
    }

    public static class Builder {
        private String street;
        private String number;
        private String complement;
        private String neighborhood;
        private String city;
        private String state;
        private String zipCode;

        public Builder street(String street) { this.street = street; return this; }
        public Builder number(String number) { this.number = number; return this; }
        public Builder complement(String complement) { this.complement = complement; return this; }
        public Builder neighborhood(String neighborhood) { this.neighborhood = neighborhood; return this; }
        public Builder city(String city) { this.city = city; return this; }
        public Builder state(String state) { this.state = state; return this; }
        public Builder zipCode(String zipCode) { this.zipCode = zipCode; return this; }

        public Address build() {
            Objects.requireNonNull(street, "street is required");
            Objects.requireNonNull(city, "city is required");
            Objects.requireNonNull(state, "state is required");
            Objects.requireNonNull(zipCode, "zipCode is required");
            return new Address(this);
        }
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof Address that)) return false;
        return Objects.equals(street, that.street)
            && Objects.equals(number, that.number)
            && Objects.equals(city, that.city)
            && Objects.equals(state, that.state)
            && Objects.equals(zipCode, that.zipCode);
    }

    @Override
    public int hashCode() {
        return Objects.hash(street, number, city, state, zipCode);
    }

    @Override
    public String toString() {
        return "%s, %s — %s/%s — %s".formatted(street, number, city, state, zipCode);
    }
}
