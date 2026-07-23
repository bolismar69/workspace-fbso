package com.fbso.platform.admin.controller;

import com.fbso.platform.admin.dto.request.ForgotPasswordRequest;
import com.fbso.platform.admin.dto.request.ResetPasswordRequest;
import com.fbso.platform.admin.dto.response.AuthResponse;
import com.fbso.platform.admin.service.AuthService;
import jakarta.validation.Valid;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1/auth")
public class AuthController {

    private final AuthService authService;

    public AuthController(AuthService authService) {
        this.authService = authService;
    }

    /**
     * Login — redireciona para Keycloak Authorization Code Flow.
     * O OAuth2 Client (SecurityConfig) gerencia o fluxo.
     */
    @PostMapping("/login")
    public ResponseEntity<AuthResponse> login() {
        return ResponseEntity.ok(authService.login());
    }

    /**
     * Forgot Password — envia link de redefinição por email (RN13-03: expira 1h).
     */
    @PostMapping("/forgot-password")
    public ResponseEntity<AuthResponse> forgotPassword(@Valid @RequestBody ForgotPasswordRequest request) {
        return ResponseEntity.ok(authService.forgotPassword(request));
    }

    /**
     * Reset Password — redefine a senha com token (RN13-01: 8+ chars, letra+número).
     */
    @PostMapping("/reset-password")
    public ResponseEntity<AuthResponse> resetPassword(@Valid @RequestBody ResetPasswordRequest request) {
        return ResponseEntity.ok(authService.resetPassword(request));
    }
}
