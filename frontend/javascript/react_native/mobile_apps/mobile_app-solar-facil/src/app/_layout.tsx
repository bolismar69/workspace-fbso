// /src/app/_layout.tsx
import { Tabs } from "expo-router";
import React from "react";
import { StatusBar } from "expo-status-bar";
import { SafeAreaProvider } from "react-native-safe-area-context";

import { Ionicons, SimpleLineIcons, MaterialIcons, } from "@expo/vector-icons";
import { AppThemeProvider } from "@/context/AppThemeContext";
import { AuthProvider, useAuth } from "@/context/AuthContext";
import { GestureHandlerRootView } from "react-native-gesture-handler";
// import SolEnergiaIcon from "@/assets/icones/sol_energia.png";
// import { TabIcon } from "@/components/ui/TabIcon";
// import { FaHouseDamage } from "react-icons/fa";
// import { LiaHouseDamageSolid } from "react-icons/lia";
import { useColorScheme } from "react-native";
import { DatabaseProvider } from "@/context/DatabaseContext";
import { ReactQueryProvider } from "@/context/ReactQueryProvider";

export default function RootLayout() {
  console.log("RootLayout");
  return (
    <ReactQueryProvider>
      <AppThemeProvider>
        <AuthProviderWrapper />
      </AppThemeProvider>
    </ReactQueryProvider>
  );
}

function AuthProviderWrapper() {
  console.log("AuthProviderWrapper");
  return (
    <AuthProvider>
      <SafeAreaProvider>
        <StatusBar style="auto" />
        <DatabaseProvider autoInitialize={true} >
          <AuthProtectedSlot />
        </DatabaseProvider>
      </SafeAreaProvider>
    </AuthProvider>
  );
}

// function DatabaseInitializerWrapper() {
//   const [loading, setLoading] = useState(true);

//   useEffect(() => {
//     const setup = async () => {
//       try {
//         await initializeDatabase(); // ⬅️ Inicializa o banco aqui
//       } catch (error) {
//         console.error("Erro ao inicializar banco de dados", error);
//       } finally {
//         setLoading(false);
//       }
//     };
//     setup();
//   }, []);

//   if (loading) {
//     return (
//       <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
//         <ActivityIndicator size="large" />
//       </View>
//     );
//   }
//   return <AuthProtectedSlot />;
// }

function AuthProtectedSlot() {
  console.log("AuthProtectedSlot");

  const colorScheme = useColorScheme();
  const { isLoggedIn } = useAuth();
  console.log("AuthProtectedSlot isLoggedIn:", isLoggedIn);

  // const iconFocusedColor = "#2E7D32"; // Verde para o ícone focado
  // const iconFocusedColor = "#388E3C"; // Verde para o ícone focado
  const iconFocusedColor = "#43A047"; // Verde para o ícone focado
  // const iconFocusedColor = "#04f611"; // Verde para o ícone focado

  // opcoes: "#ffff56", "#ffff42", "#ffff52", "#ffff7a", "#ffff9e", "#ffffbf", "#ffffdf"
  const iconActiveBackgroundColor = "#ffffbf"; // Cor de fundo para ícones focados --default:"#f0f0f0"
  const iconInactiveBackgroundColor = "#fff"; // Cor de fundo para ícones não foc
  const headerBackgroundColor = colorScheme === "dark" ? "#1e1e1e" : "#ffffbf"; // Cor de fundo do cabeçalho

  return (
    <GestureHandlerRootView style={{ flex: 1 }}>
      {/* <AuthProvider> */}
      {/* <AppThemeProvider> */}
      <React.Fragment>
        <StatusBar style="auto" />
        <Tabs
          screenOptions={{
            headerShown: true, // Exibir o cabeçalho
            headerTitleAlign: "center", // Alinhar o título do cabeçalho ao
            headerStyle: {
              backgroundColor: headerBackgroundColor, // Cor de fundo do cabeçalho
            },
            // headerTintColor: colorScheme === "dark" ? "#fff" : "#000", // Cor do texto do cabeçalho
            headerTitleStyle: {
              color: colorScheme === "dark" ? "#fff" : "#000", // Cor do título do cabeçalho
              fontSize: 24, // Tamanho da fonte do título do cabeçalho
              fontWeight: "bold", // Negrito para o título do cabeçalho
              marginBottom: 16,
              marginTop: 16,
              alignContent: "center",
            },

            // configuracao da tab-bar
            tabBarInactiveTintColor: "#888", // Cor padrão para ícones não focados
            tabBarActiveTintColor: iconFocusedColor, // Cor para ícones focados
            tabBarActiveBackgroundColor: iconActiveBackgroundColor, // Cor de fundo para ícones focados
            tabBarInactiveBackgroundColor: iconInactiveBackgroundColor, // Cor de fundo para ícones não focados
            // cor da tab bar
            tabBarStyle: {
              backgroundColor: headerBackgroundColor, // Cor de fundo da tab bar
              borderTopWidth: 0, // Remover a borda superior da tab bar
            },
            // tabBarShowLabel: true, // Exibir rótulos na tab bar
            tabBarLabelStyle: {
              fontSize: 12, // Tamanho da fonte dos rótulos da tab bar
              fontWeight: "bold", // Negrito para os rótulos da tab bar
            },
            tabBarItemStyle: {
              marginBottom: 0, // Remover a margem inferior dos itens da tab bar
              marginTop: 0, // Remover a margem superior dos itens da tab bar
              backgroundColor: headerBackgroundColor, // Cor de fundo dos itens da tab bar
              borderTopWidth: 0, // Remover a borda superior dos itens da tab bar
              borderRadius: 0, // Remover borda arredondada dos itens da tab bar
              padding: 0, // Remover o preenchimento dos itens da tab bar

            },
            tabBarIconStyle: {
              marginBottom: 0, // Remover a margem inferior dos ícones da tab bar
              marginTop: 0, // Remover a margem superior dos ícones da tab bar
            },
          }}
        >
          <Tabs.Screen
            name="index"
            options={{
              title: "Bem vindo ao Solar Fácil",
              // headerShown: false,
              tabBarLabel: "Solar",
              tabBarIcon: ({ color, size, focused, }) => (
                // <Ionicons name="sunny-outline"
                //   color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                //   size={size}
                // />
                <MaterialIcons
                  name="solar-power"                     // "attach-money"
                  color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                  size={size}
                />
              ),
            }}
          />
          <Tabs.Screen
            name="planos"
            options={{
              title: "Planos Comerciais",
              // headerShown: false,
              tabBarLabel: "Plano",
              tabBarIcon: ({ color, size, focused, }) => (
                // <Ionicons name="pricetag-outline"
                //   color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                //   size={size}
                // />
                <MaterialIcons
                  name="currency-exchange"                     // "attach-money"
                  color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                  size={size}
                />
              ),
            }}
          />
          <Tabs.Screen
            name="saibamais"
            options={{
              title: "Saiba Mais",
              tabBarLabel: "Saiba",
              tabBarIcon: ({ color, size, focused, }) => (
                <Ionicons name="information-circle-outline"
                  color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                  size={size}
                />
              ),
            }}
          />
          <Tabs.Screen
            name="faq"
            options={{
              title: "Perguntas Frequentes",
              tabBarLabel: "FAQ",
              tabBarIcon: ({ color, size, focused, }) => (
                <Ionicons name="help-circle-outline"
                  color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                  size={size}
                />
              ),
            }}
          />
          <Tabs.Screen
            name="login"
            options={{
              title: "Faça seu Login",
              tabBarLabel: isLoggedIn === true ? "Logout" : "Login",
              headerShown: true,
              tabBarIcon: ({ color, size, focused, }) => (
                <Ionicons
                  // name={isLoggedIn === true ? "log-out-outline" : "log-in-outline"}
                  // name={isLoggedIn === true ? "lock-open-outline" : "lock-closed-outline"}
                  name={isLoggedIn === true ? "flash-off-outline" : "flash-outline"}
                  color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                  size={size}
                />
              ),
            }}
          />
          <Tabs.Screen
            name="cadastro"
            options={{
              title: "Cadastro",
              tabBarLabel: "User",
              headerShown: false,
              tabBarIcon: ({ color, size, focused, }) => (
                <Ionicons name="person-outline"
                  color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                  size={size}
                />
              ),
            }}
          />
          <Tabs.Screen
            name="movimentacao"
            options={{
              title: "Movimentações Mensais",
              tabBarLabel: "...",
              headerShown: true,
              tabBarIcon: ({ color, size, focused, }) => (
                // <Ionicons name="book-outline"
                //   color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                //   size={size}
                // />
                // <Entypo name="spreadsheet"        /// "line-graph"
                //   color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                //   size={size}
                // />
                <SimpleLineIcons
                  name="book-open" // "book-open"  //  "notebook"
                  color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                  size={size}
                />
              ),
            }}
          />
          <Tabs.Screen
            name="listatodos"
            options={{
              title: "Lista",
              tabBarLabel: "Lista",
              headerShown: true,
              tabBarIcon: ({ color, size, focused, }) => (
                <Ionicons name="list-outline"
                  color={focused ? iconFocusedColor : color} // Verde se focado, cor padrão caso contrário
                  size={size}
                />
              ),
            }}
          />
        </Tabs>
      </React.Fragment>
      {/* </AppThemeProvider> */}
      {/* </AuthProvider> */}
    </GestureHandlerRootView >
  )
}
