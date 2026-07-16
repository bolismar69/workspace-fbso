import React from "react";
import { KeyboardAvoidingView, SafeAreaView, ScrollView, Platform } from "react-native";
import { useAppTheme } from "@/context/AppThemeContext";
import { FormDadosCadastraisAssociado } from "@/components/forms/FormDadosCadastraisAssociado";
import { ContatoRodapeCopyRight } from "@/components/ContatoRodapeCopyRight";
import { useAuth } from "@/context/AuthContext";
import { FormCadastroAssociado } from "@/components/forms/FormCadastroAssociado";
import { AssociadoType } from "@/types/AssociadoType";

// Refactored to accept a single props object:
type AssociadoDadosCadastraisScreenProps = {
  actionForm?: string | null; // possiveis valores Default, New, Edit
  itemAssociado?: AssociadoType | null;
  editable?: boolean | true; // Default to true if not provided
};

export default function AssociadoDadosCadastraisScreen({ actionForm, itemAssociado, editable }: AssociadoDadosCadastraisScreenProps) {
  console.log("AssociadoDadosCadastraisScreen - Renderizando tela de dados cadastrais do associado");
  const { theme } = useAppTheme();
  const { isLoggedIn, userID, userName, associado } = useAuth();
  console.log("AssociadoDadosCadastraisScreen - isLoggedIn:", isLoggedIn);
  console.log("AssociadoDadosCadastraisScreen - userID:", userID);
  console.log("AssociadoDadosCadastraisScreen - userName:", userName);
  console.log("AssociadoDadosCadastraisScreen - associado:", associado);
  console.log("AssociadoDadosCadastraisScreen - param: actionForm:", actionForm);
  console.log("AssociadoDadosCadastraisScreen - param: itemAssociado:", itemAssociado);
  console.log("AssociadoDadosCadastraisScreen - param: editable:", editable);

  const handleSubmit = (data: any) => {
    console.log("AssociadoDadosCadastraisScreen - Associado cadastrado com sucesso:", data);
  };

  return (
    <KeyboardAvoidingView
      style={{ flex: 1 }}
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      keyboardVerticalOffset={Platform.OS === "ios" ? 64 : 0}
    >
      <SafeAreaView style={{ flex: 1 }}>
        <ScrollView
          contentContainerStyle={{
            padding: 16,
            flexGrow: 1,
            backgroundColor: theme.screen?.backgroundColor,
          }}
          keyboardShouldPersistTaps="handled"
        >

          {actionForm
            ? (actionForm === "New"
              ? (
                console.log("AssociadoDadosCadastraisScreen - vai param actionForm[New] - Renderizando FormCadastroAssociado"),
                <FormCadastroAssociado
                  onSubmit={handleSubmit}
                />
              )
              : (
                console.log("AssociadoDadosCadastraisScreen - vai param actionForm[Edit] - Renderizando FormDadosCadastraisAssociado"),
                <FormDadosCadastraisAssociado
                  associado={itemAssociado!}
                  onSubmit={handleSubmit}
                  editable={editable}
                />
              )
            )
            : (isLoggedIn === true && associado
              ? (
                console.log("AssociadoDadosCadastraisScreen - Renderizando FormDadosCadastraisAssociado"),
                <FormDadosCadastraisAssociado
                  associado={associado}
                  onSubmit={handleSubmit}
                />
              ) : (
                console.log("AssociadoDadosCadastraisScreen - Renderizando FormCadastroAssociado"),
                <FormCadastroAssociado
                  onSubmit={handleSubmit}
                />
              )
            )
          }
        </ScrollView>
        <ContatoRodapeCopyRight />
      </SafeAreaView>
    </KeyboardAvoidingView>
  );
}
