// src/components/ContatoRodape.tsx
import React from "react";
import { View, Text, TouchableOpacity, Linking } from "react-native";
import { Ionicons, FontAwesome, MaterialCommunityIcons, AntDesign, Feather, MaterialIcons } from "@expo/vector-icons";
import { useAppTheme } from "@/context/AppThemeContext";

interface ContatoRodapeProps {
  label: string; // Ensure 'label' is defined
}

export const ContatoRodape: React.FC<ContatoRodapeProps> = ({ label }) => {
  const { theme } = useAppTheme();

  const handlePress = (url: string) => {
    Linking.openURL(url).catch((err) =>
      console.error("Erro ao abrir link:", err)
    );
  };

  return (
    <View style={{
      marginTop: 48,
      marginBottom: 16,
      borderRadius: 8,
      gap: 12,
      borderWidth: 0.5,
      borderTopColor: theme.basicView.borderColor,
      // elevation: 2, // Mantido para compatibilidade com Android
      // boxShadow: "0px 2px 4px rgba(0, 0, 0, 0.1)", // Substituído por boxShadow
    }}
    >
      <Ionicons
        name="chatbubble-ellipses-outline"
        size={64}
        color={theme.title.color}
        style={{ alignSelf: "center", marginTop: 16 }}
      />

      {/* <MaterialIcons
        name="chat"
        size={64}
        color={theme.title.color}
        style={{ alignSelf: "center", marginTop: 16 }}
      /> */}

      {/* <Feather
        name="user"
        size={64}
        color={theme.title.color}
        style={{ alignSelf: "center", marginTop: 16 }}
      /> */}

      <Text style={[theme.basicText, { textAlign: "center", marginTop: 0 }]}>
        {label}
      </Text>

      <View
        style={{
          flexDirection: "row",
          justifyContent: "space-around",
          marginTop: 8,
          marginBottom: 8,
        }}
      >
        <TouchableOpacity
          onPress={() => handlePress("mailto:contato@solarfacil.com")}
        >
          <Ionicons name="mail" size={32} color={theme.title.color} />
        </TouchableOpacity>

        <TouchableOpacity
          onPress={() =>
            handlePress(
              "https://wa.me/5511956781234?text=Olá%2C%20estou%20interessado%20em%20saber%20mais%20sobre%20o%20Solar%20Fácil.%20Me%20passe%20por%20aqui%20o%20link%20para%20ser%20adicionado."
            )
          }
        >
          <Ionicons name="logo-whatsapp" size={32} color="#25D366" />
        </TouchableOpacity>

        <TouchableOpacity
          onPress={() => handlePress("https://www.instagram.com/instagram")}
        >
          <Ionicons name="logo-instagram" size={32} color="#C13584" />
        </TouchableOpacity>

        <TouchableOpacity
          onPress={() => handlePress("https://www.facebook.com/AdoroCinema")}
        >
          <Ionicons name="logo-facebook" size={32} color="#1877F2" />
        </TouchableOpacity>

        <TouchableOpacity onPress={() => handlePress("https://www.youtube.com/")}>
          <FontAwesome name="youtube" size={32} color="#FF0000" />
        </TouchableOpacity>
      </View>
    </View>
  );
};
