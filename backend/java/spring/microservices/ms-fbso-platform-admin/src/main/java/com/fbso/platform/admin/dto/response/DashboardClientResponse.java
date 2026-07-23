package com.fbso.platform.admin.dto.response;

import java.util.List;

public record DashboardClientResponse(
    int activeUnits, String activeUnitsLink,
    int productCount, String productsLink,
    String planName, String planStatus, String planLink,
    List<NotificationResponse> notifications, String notificationsLink
) {}
