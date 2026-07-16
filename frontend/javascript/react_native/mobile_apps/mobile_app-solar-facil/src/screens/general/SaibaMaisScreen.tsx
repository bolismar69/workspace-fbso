import { useAppTheme } from "@/context/AppThemeContext";
import { Image, KeyboardAvoidingView, Platform, SafeAreaView, ScrollView, Text } from "react-native";
import CardIconeAmarelo from "@/components/CardIconeAmarelo";
import CardIconePadrao from "@/components/CardIconePadrao";
import { ContatoRodape } from "@/components/ContatoRodapeIconesContato";
import { ContatoRodapeCopyRight } from "@/components/ContatoRodapeCopyRight";

export default function SaibaMaisScreen() {
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

          {/* Título */}
          {/* <Text style={theme.title}>Sobre a Solar Fácil</Text> */}

          {/* Imagem institucional */}
          {/* <Image
            source={require("@/assets/solar-facil-institucional.png")}
            style={theme.imagePreview}
          /> */}
          {/* Banner com imagem */}
          <Image
            source={require("@/assets/solar-facil-antenas.png")}
            resizeMode="cover"
            style={theme.imagePreview}
          />

          {/* Apresentação institucional */}
          <Text style={theme.text}>
            A Solar Fácil é uma cooperativa moderna que conecta Produtores de Energia Solar aos Beneficiados — pessoas físicas ou jurídicas que desejam economizar na conta de luz.
          </Text>
          <Text style={theme.text}>
            Os produtores injetam energia excedente na rede da Distribuidora local, e a Solar Fácil distribui esse crédito entre os cooperados de forma transparente, segura e sustentável.
          </Text>
          <Text style={theme.text}>
            Nosso modelo é descentralizado, colaborativo e regulado pela ANEEL. Geramos impacto ambiental positivo, economia para os beneficiados e retorno para os fornecedores.
          </Text>

          {/* Nossos Valores */}
          <Text style={theme.subtitle}>Nossos Valores</Text>

          {/* Cards com ícones e destaques */}
          <CardIconeAmarelo
            iconName="people-circle-outline"
            iconColor="green"
            iconBackground="yellow"
            title="Cooperação"
            description="Valorizamos a colaboração entre os membros para construir um futuro energético melhor."
            theme={theme}
          />
          <CardIconeAmarelo
            iconName="shield-checkmark-outline"
            iconColor="green"
            iconBackground="yellow"
            title="Transparência"
            description="Todas as operações são claras, com rastreabilidade dos créditos e repasses."
            theme={theme}
          />
          <CardIconeAmarelo
            iconName="sunny-outline"
            iconColor="green"
            iconBackground="yellow"
            title="Sustentabilidade"
            description="Reduzimos a pegada de carbono ao promover o uso de energia limpa e renovável."
            theme={theme}
          />

          {/* Nossa História */}
          <Text style={theme.subtitle}>Nossa História</Text>

          <CardIconePadrao
            iconName="calendar-start"
            iconColor={theme.title.color}
            iconBackground=""
            title="2022 — Ideia e Pesquisa"
            description="Estudo sobre modelos de cooperativas energéticas e viabilidade técnica no Brasil."
            theme={theme}
          />

          <CardIconePadrao
            iconName="lightbulb-on-outline"
            iconColor={theme.title.color}
            iconBackground=""
            title="2023 — Criação da Solar Fácil"
            description="Lançamento da cooperativa com foco em energia solar compartilhada."
            theme={theme}
          />

          <CardIconePadrao
            iconName="account-group-outline"
            iconColor={theme.title.color}
            iconBackground=""
            title="2024 — Primeiros Cooperados"
            description="Início da operação com os primeiros produtores e beneficiados em São Paulo."
            theme={theme}
          />

          <CardIconePadrao
            iconName="rocket-launch"
            iconColor={theme.title.color}
            iconBackground=""
            title="2025 — Expansão Nacional"
            description="Expansão da atuação para outros estados, integrando mais redes e cooperados."
            theme={theme}
          />

          <ContatoRodape label="Nos contate ou nos siga nas redes sociais para conhecer melhor a nossa proposta de ajudar o planeta, ajudar as pessoas e a sociedade, em como enconomizar e agregar ganhos com a energia elétrica." />

        </ScrollView>
        <ContatoRodapeCopyRight />
      </SafeAreaView>
    </KeyboardAvoidingView>
  );
}
