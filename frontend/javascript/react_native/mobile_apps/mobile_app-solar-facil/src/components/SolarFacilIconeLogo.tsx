import React from "react";
import { View, Text, Image, StyleSheet } from "react-native";

const SolarFacilIconeLogo: React.FC = () => {
  return (
    <View style={styles.container}>
      {/* <Image source={require("@/assets/images/solar-facil-favicon.png")} style={styles.logo} /> */}
      <Text style={styles.title}>Solar Fácil</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    alignItems: "center",
    justifyContent: "center",
    padding: 16,
    backgroundColor: "#f5f5f5",
  },
  logo: {
    width: 50,
    height: 50,
    marginBottom: 8,
  },
  title: {
    fontSize: 18,
    fontWeight: "bold",
    color: "#333",
  },
});

export default SolarFacilIconeLogo;
