import { useAppTheme } from "@/context/AppThemeContext";
import { CardPlan } from "@/components/CardPlan";
import { PlanType } from "@/types/PlanType";
import { fetchPlans } from "@/services/servicePlans";
import { Image, KeyboardAvoidingView, Platform, SafeAreaView, ScrollView, Text } from "react-native";
import { useEffect, useState } from "react";
import { ContatoRodape } from "@/components/ContatoRodapeIconesContato";
import { ContatoRodapeCopyRight } from "@/components/ContatoRodapeCopyRight";

export default function PlanosScreen() {
  const { theme } = useAppTheme();
  const [plans, setPlans] = useState<PlanType[]>([]);

  useEffect(() => {
    fetchPlans().then(setPlans);
  }, []);

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
          {/* <Text style={theme.title}>Planos Comerciais</Text> */}

          {/* Banner com imagem */}
          <Image
            source={require("@/assets/solar-facil-antenas.png")}
            resizeMode="cover"
            style={theme.imagePreview}
          />

          <Text style={[theme.subtitle, { marginBottom: 0 }]}>
            Conheça nossos planos de energia solar pensados para diferentes perfis de consumo.
          </Text>

          {plans.map((plan) => (
            <CardPlan key={plan.name} plan={plan} />
          ))}

          {/* Ícones de Contato */}
          <ContatoRodape label="Para saber mais detalhes sobre nossos planos e propostas, de ganhos em economia com energia elétrica, ou como rentabilizar com a sua produção de energia elétrica, entre em contato conosco." />

        </ScrollView>
        <ContatoRodapeCopyRight />
      </SafeAreaView>
    </KeyboardAvoidingView>
  );
}
