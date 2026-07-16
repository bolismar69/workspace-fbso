// src/components/layout/KeyboardSafeScreen.tsx
import React from "react";
import {
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  SafeAreaView,
  ViewStyle,
  ScrollViewProps,
} from "react-native";

interface Props extends ScrollViewProps {
  children: React.ReactNode;
  style?: ViewStyle;
  contentContainerStyle?: ViewStyle;
}

export function KeyboardSafeScreen({
  children,
  style,
  contentContainerStyle,
  ...rest
}: Props) {
  console.log("KeyboardSafeScreen rendered");
  return (
    <KeyboardAvoidingView
      style={{ flex: 1 }}
      behavior={Platform.OS === "ios" ? "padding" : "height"}
    >
      <SafeAreaView style={{ flex: 1 }}>
        <ScrollView
          style={style}
          contentContainerStyle={[{ padding: 16 }, contentContainerStyle]}
          keyboardShouldPersistTaps="handled"
          {...rest}
        >
          {children}
        </ScrollView>
      </SafeAreaView>
    </KeyboardAvoidingView>
  );
}
