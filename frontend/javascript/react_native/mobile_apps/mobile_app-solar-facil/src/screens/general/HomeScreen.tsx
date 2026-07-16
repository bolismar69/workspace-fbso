import { useAppTheme } from "@/context/AppThemeContext";
import { Ionicons, MaterialCommunityIcons } from "@expo/vector-icons";
import { Image, KeyboardAvoidingView, Platform, SafeAreaView, ScrollView, Text, TouchableOpacity, View } from "react-native";
import { ContatoRodape } from "@/components/ContatoRodapeIconesContato";
import { ContatoRodapeCopyRight } from "@/components/ContatoRodapeCopyRight";

export default function HomeScreen() {
  const { theme } = useAppTheme();

  return (
    <KeyboardAvoidingView
      style={{ flex: 1 }}
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      // ou ajustar conforme cabeçalho
      keyboardVerticalOffset={Platform.OS === "ios" ? 64 : 0}
    >
      <SafeAreaView style={[{ flex: 1 }]}>
        <ScrollView
          style={{ flex: 1, backgroundColor: theme.basicView.backgroundColor }}
          contentContainerStyle={{ padding: 16, paddingBottom: 100 }}
        >

          {/* Apresentação principal */}
          {/* <Text style={theme.title}>Bem vindo a Solar Fácil</Text> */}

          {/* Banner com imagem */}
          <Image
            source={require("@/assets/solar-facil-antenas.png")}
            style={theme.imagePreview}
          />

          {/* Descrição */}
          <Text style={theme.subtitle}>
            Sua plataforma inteligente, para economizar, com energia limpa e ajudar o planeta.
          </Text>

          {/* Cards com ícones e destaques */}
          <View style={theme.card}>
            {/* <View style={[theme.iconContainer, { backgroundColor: "yellow" }]}> */}
            <Ionicons name="sunny-outline" size={28} color={theme.title.color} />
            {/* </View> */}

            {/*           <View style={[theme.iconContainer, { backgroundColor: "yellow", position: "relative" }]}>
            <View
              style={{
                position: "absolute",
                width: 20,
                height: 20,
                borderRadius: 10,
                backgroundColor: "yellow",
                top: 14,
                left: 14,
              }}
            />
            <Ionicons name="sunny-outline" size={56} color="black" />
          </View> */}
            {/* TODO: colocar a cor amarelo dentro do outline */}

            <Text style={theme.highlightText}>Energia Sustentável</Text>
            <Text style={theme.text}>
              Conectamos você a fontes de energia solar sem instalação e sem investimento inicial.
            </Text>
          </View>

          <View style={theme.card}>
            <MaterialCommunityIcons name="cash-refund" size={28} color={theme.title.color} />
            <Text style={theme.highlightText}>Economize até 20%</Text>
            <Text style={theme.text}>
              Pague menos na sua conta de luz e acompanhe tudo de forma transparente.
            </Text>
            <TouchableOpacity style={theme.button}>
              <Text style={theme.buttonText}>Junte-se a nós e comece a economizar</Text>
            </TouchableOpacity>
          </View>

          <View style={theme.card}>
            <Ionicons name="rocket-outline" size={28} color={theme.title.color} />
            <Text style={theme.highlightText}>Retorno para Produtores</Text>
            <Text style={theme.text}>
              Geração excedente? Venha compartilhar com a comunidade e tenha retorno financeiro.
            </Text>
            <TouchableOpacity style={theme.button}>
              <Text style={theme.buttonText}>Ingresse como Fornecedor</Text>
            </TouchableOpacity>
          </View>

          {/* Ícones de Contato */}
          <ContatoRodape label="Nos contate ou nos siga nas redes sociais, e venha conhecer mais a Solar Fácil, e lhe mostraremos como economizar com energia elétrica, ou como rentabilizar sua produção." />

          {/* Seção Sobre */}
          {/*           <Text style={theme.subtitle}>Nos contate ou nos siga nas redes sociais para mais informações</Text>
          <View style={{ flexDirection: "row", justifyContent: "space-around", marginTop: 8 }}>
            <TouchableOpacity onPress={() => handlePress("mailto:contato@solarfacil.com")}>
              <Ionicons name="mail" size={32} color={theme.title.color} />
            </TouchableOpacity>
            <TouchableOpacity onPress={() => handlePress("https://wa.me/5511956781234?text=Olá%2C%20estou%20interessado%20em%20saber%20mais%20sobre%20o%20Solar%20Fácil.%20Me%20passe%20por%20aqui%20o%20link%20para%20ser%20adicionado.")}>
              <Ionicons name="logo-whatsapp" size={32} color="#25D366" />
            </TouchableOpacity>
            <TouchableOpacity onPress={() => handlePress("https://www.instagram.com/instagram")}>
              <Ionicons name="logo-instagram" size={32} color="#C13584" />
            </TouchableOpacity>
            <TouchableOpacity onPress={() => handlePress("https://www.facebook.com/AdoroCinema")}>
              <Ionicons name="logo-facebook" size={32} color="#1877F2" />
            </TouchableOpacity>
            <TouchableOpacity onPress={() => handlePress("https://www.youtube.com/")}>
              <FontAwesome name="youtube" size={32} color="#FF0000" />
            </TouchableOpacity>
          </View> */}

        </ScrollView>
        <ContatoRodapeCopyRight />
      </SafeAreaView>
    </KeyboardAvoidingView>
  );
}
