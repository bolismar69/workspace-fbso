import React from "react";
// import AssociadoCadastroScreen from "@/screens/associado/AssociadoCadastroScreen";
import AssociadoDadosCadastraisScreen from "@/screens/associado/AssociadoDadosCadastraisScreen";
export default function AppCadastroScreen() {

  return (
    console.log("AppFaqScreen"),
    // <AssociadoCadastroScreen />
    <AssociadoDadosCadastraisScreen itemAssociado={null} editable={true} />
  );
}
