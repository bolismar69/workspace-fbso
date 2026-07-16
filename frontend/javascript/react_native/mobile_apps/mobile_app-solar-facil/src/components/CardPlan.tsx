// src/components/CardPlan.tsx
import { useAppTheme } from "@/context/AppThemeContext";
import { Ionicons } from "@expo/vector-icons";
import { router } from "expo-router";
import React, { useState } from "react";
import { Pressable, StyleSheet, Text, TouchableOpacity, View } from "react-native";
import Animated, { useSharedValue, useAnimatedStyle, withTiming } from "react-native-reanimated";
import { PlanType } from "@/types/PlanType";

type CardPlanProps = {
  plan: PlanType;
};

export function CardPlan({ plan }: CardPlanProps) {
  const { theme } = useAppTheme();
  const [flipped, setFlipped] = useState(false);
  const rotation = useSharedValue(0);

  const frontAnimatedStyle = useAnimatedStyle(() => ({
    transform: [{ rotateY: `${rotation.value}deg` }],
  }));

  const backAnimatedStyle = useAnimatedStyle(() => ({
    transform: [{ rotateY: `${rotation.value + 180}deg` }],
  }));

  const handleFlip = () => {
    const next = flipped ? 0 : 180;
    rotation.value = withTiming(next, { duration: 400 });
    setFlipped(!flipped);
  };

  const dynamicCardStyle = {
    borderColor: plan.commercialIndication ? "#2563EB" : "#ccc",
    borderWidth: plan.commercialIndication ? 2 : 1,
    shadowOpacity: plan.commercialIndication ? 0.25 : 0.1,
    shadowRadius: plan.commercialIndication ? 8 : 4,
    elevation: plan.commercialIndication ? 6 : 2,
  };

  return (
    <View style={styles.container}>
      <Pressable onPress={handleFlip}>
        <View style={styles.cardWrapper}>
          {/* Frente */}
          <Animated.View
            style={[
              styles.face,
              frontAnimatedStyle,
              dynamicCardStyle,
              { backgroundColor: theme.card.backgroundColor },
            ]}
          >
            <View style={{ flex: 1, gap: 8 }}>
              <Text style={[styles.popular, { marginBottom: 24, backgroundColor: plan.commercialIndication ? "#2563EB" : "#94a3b8" }]}>
                {plan.commercialAttraction}
              </Text>

              <View style={styles.titleRow}>
                <Ionicons name={plan.icon} size={24} color="#fbbf24" style={{ marginRight: 12 }} />
                <Text style={[theme.subtitle, { fontWeight: "bold", marginBottom: 0 }]}>
                  Plano {plan.name}
                </Text>
                <TouchableOpacity onPress={handleFlip}>
                  <Ionicons name="information-circle-outline" size={24} color={theme.text.color} />
                </TouchableOpacity>
              </View>

              <Text style={theme.text}>Potência: {plan.powerRange}</Text>
              <Text style={theme.text}>Consumo: {plan.consumption}</Text>
              <Text style={theme.text}>Tarifa: {plan.pricePerKwh}</Text>
              <Text style={theme.text}>Valor Estimado: {plan.monthlyEstimate}</Text>
              <Text style={theme.text}>Custo de Energia: {plan.energyCost}</Text>

              <TouchableOpacity
                style={[styles.button, { backgroundColor: theme.button.backgroundColor }]}
                onPress={() => router.push("/login")}
              >
                <Text style={[theme.buttonText, { fontWeight: "600" }]}>Assinar</Text>
              </TouchableOpacity>
            </View>
          </Animated.View>

          {/* Verso */}
          <Animated.View
            style={[
              styles.back,
              backAnimatedStyle,
              dynamicCardStyle,
              { backgroundColor: theme.card.backgroundColor },
            ]}
          >
            <View style={{ flex: 1, gap: 8 }}>

              <View style={styles.titleRow}>
                <Ionicons name={plan.icon} size={24} color="#fbbf24" style={{ marginRight: 12 }} />
                <Text style={[theme.title, { fontWeight: "bold" }]}>Plano {plan.name}</Text>
                <TouchableOpacity onPress={handleFlip}>
                  <Ionicons name="information-circle-outline" size={24} color={theme.text.color} />
                  {/* <MaterialIcons name="attach-money" size={24} color={theme.text.color} /> */}
                </TouchableOpacity>
              </View>

              <Text style={[theme.text, { fontStyle: "italic", textAlign: "center" }]}>
                {plan.description}
              </Text>

              <TouchableOpacity
                style={[styles.button, { backgroundColor: theme.button.backgroundColor }, { marginTop: 170 }]}
                onPress={() => router.push("/login")}
              >
                <Text style={[theme.buttonText, { fontWeight: "600" }]}>Assinar</Text>
              </TouchableOpacity>
            </View>
          </Animated.View>
        </View>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    marginTop: 30,
    perspective: "1000",
  },
  cardWrapper: {
    width: "100%",
    minHeight: 380,
    position: "relative",
  },
  face: {
    position: "absolute",
    width: "100%",
    height: "100%",
    borderRadius: 12,
    padding: 16,
    justifyContent: "space-between",
    backfaceVisibility: "hidden",
    shadowColor: "#000",
  },
  back: {
    position: "absolute",
    width: "100%",
    height: "100%",
    borderRadius: 12,
    padding: 16,
    justifyContent: "center",
    alignItems: "center",
    backfaceVisibility: "hidden",
    shadowColor: "#000",
  },
  popular: {
    alignSelf: "center",
    color: "white",
    fontSize: 12,
    fontWeight: "bold",
    paddingHorizontal: 12,
    paddingVertical: 4,
    borderRadius: 6,
    marginBottom: 8,
  },
  titleRow: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
    marginBottom: 8,
    marginTop: -30,
  },
  button: {
    marginTop: 12,
    paddingVertical: 10,
    borderRadius: 8,
    alignItems: "center",
  },
});
