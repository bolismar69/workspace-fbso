// src/components/ui/TabIcon.tsx
import React from "react";
import { Image, ImageSourcePropType } from "react-native";

export function TabIcon({ source, focused }: { source: ImageSourcePropType; focused: boolean }) {
  return (
    <Image
      source={source}
      style={{
        width: 24,
        height: 24,
        opacity: focused ? 1 : 0.6,
        tintColor: focused ? "#F5B000" : "#4CAF50",
      }}
      resizeMode="contain"
    />
  );
}
