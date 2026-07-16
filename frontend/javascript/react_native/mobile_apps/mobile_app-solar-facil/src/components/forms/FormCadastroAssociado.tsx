import React, { useState } from "react";
import { useRouter } from "expo-router";
import { Text, TouchableOpacity, View } from "react-native";
import { Ionicons } from "@expo/vector-icons";

import { FormSection } from "@/components/forms/FormSection";
import { AssociadoType } from "@/types/AssociadoType";
import { FieldDefinitionType } from "@/types/FieldDefinitionType";
// import { salvarAssociado } from "@/services/storage/serviceAssociado";
import { isValidCPF } from "@/utils/validators/validatorCPF";
import { isValidCNPJ } from "@/utils/validators/validatorCNPJ";
import { useAssociadosCopilot } from "@/services/database/useAssociadosCopilot"

interface FormCadastroAssociadoProps {
  onSubmit: (data: any) => void; // Define the onSubmit prop
}

export function FormCadastroAssociado({ onSubmit }: FormCadastroAssociadoProps) {
  const router = useRouter();

  const useAssociados = useAssociadosCopilot(); // Call the hook here

  const [mensagem, setMensagem] = useState<{
    tipo: "success" | "error";
    texto: string;
  } | null>(null);

  const [formVisivel, setFormVisivel] = useState(true);

  const [dadosAssociado, setDadosAssociado] = useState<AssociadoType | null>(null);

  const fields: FieldDefinitionType<AssociadoType>[] = [
    {
      name: "tipoPessoa",
      label: "Tipo de Pessoa",
      placeholder: "Selecione",
      type: "select",
      required: true,
      options: [
        { value: "Pessoa Física", label: "Pessoa Física" },
        { value: "Pessoa Jurídica", label: "Pessoa Jurídica" },
      ],
    },
    {
      name: "cpf_cnpj",
      label: "CPF/CNPJ",
      placeholder: "000.000.000-00",
      type: "text",
      required: true,
      keyboardType: "numeric",
      validation: (value, allValues) => {
        if (allValues.tipoPessoa === "fisica") {
          if (!value) return "CPF é obrigatório para pessoa física";
          if (!isValidCPF(value)) return "CPF inválido";
        }
        if (allValues.tipoPessoa === "juridica") {
          if (!value) return "CNPJ é obrigatório para pessoa jurídica";
          if (!isValidCNPJ(value)) return "CNPJ inválido";
        }
        return null;
      },
    },
    {
      name: "nome",
      label: "Nome completo ou Razão Social",
      placeholder: "João Silva",
      type: "text",
      required: true,
      toUpperCase: true,
    },
    {
      name: "email",
      label: "E-mail",
      placeholder: "exemplo@email.com",
      type: "email",
      required: true,
      validation: (value) => {
        if (!value) return "E-mail é obrigatório";
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(value)) return "E-mail inválido";
        return null;
      }
    },
    {
      name: "telefone",
      label: "Telefone",
      placeholder: "(11) 99999-0000",
      type: "text",
      required: true,
      validation: (value) => {
        if (!value) return "Telefone é obrigatório";
        // const phoneRegex = /^\(\d{2}\) \d{5}-\d{4}$/;
        const phoneRegex = /^\d{2}\d{4,5}\d{4}$/;
        if (!phoneRegex.test(value)) return "Telefone inválido. Formato esperado: (XX) XXXXX-XXXX";
        return null;
      }
    },
    {
      name: "senha",
      label: "Senha",
      placeholder: "Digite sua senha",
      type: "password",
      required: true,
      validation: (value) => {
        if (!value) return "Senha é obrigatória";
        if (value.length < 6) return "A senha deve ter pelo menos 6 caracteres";
        return null;
      }
    }
  ];

  const initialValues: AssociadoType = {
    id: 0,
    tipoPessoa: "Pessoa Física",
    cpf_cnpj: "",
    nome: "",
    email: "",
    telefone: "",
    dataCadastro: "", // Example default value
    dataAtualizacao: "", // Example default value
    senha: "",
    status: "Em cadastro", // Example default value
    tipoAssociado: "Beneficiado", // Example default value
    cep: "",
    endereco: "",
    numero: "",
    bairro: "",
    cidade: "",
    estado: "",
    complemento: "",
    aceitaTermos: "Sim",
    observacoes: "",
    dataNascimento: "", // Example default value
    nomeSocial: "", // Example default value
    dataAbertura: "", // Example default value
    razaoSocial: "", // Example default value
    nomeFantasia: "", // Example default value
    nomeConcessionaria: "", // Example default value
    consumoMedio: "", // Example default value
    planoDesejado: "", // Example default value
    potenciaInstalada: "", // Example default value
    disponibilidade: "", // Example default value
    tipoConexao: "", // Example default value
  };

  const handleSubmit = async (data: AssociadoType) => {
    try {
      // cria um id único para o associado
      data.dataCadastro = new Date().toISOString();
      data.dataAtualizacao = new Date().toISOString();
      data.cpf_cnpj = data.cpf_cnpj.replace(/\D/g, ""); // Remove caracteres não numéricos

      const result = await (await useAssociados).insertRecord(data); // Salva o associado no banco de dados
      if (result.success === false) {
        setMensagem({
          tipo: "error",
          texto: result.error || " - Erro ao salvar o associado.",
        });
        setFormVisivel(false);
      }

      data.id = result.data?.rowID || 0; // <== pega o id gerado pelo banco de dados
      setDadosAssociado(data); // <== salva para usar na navegação
      setMensagem({
        tipo: "success",
        texto: "Parabéns, cadastro realizado com sucesso!",
      });
      setFormVisivel(false);
    } catch (error: any) {
      setMensagem({
        tipo: "error",
        texto: error?.message || "Erro ao salvar o associado.",
      });
      setFormVisivel(false);
    }
  };

  const handleMensagemPress = () => {
    setMensagem(null);
    setFormVisivel(true);

    if (mensagem?.tipo === "success" && dadosAssociado) {
      router.push({
        pathname: "/cadastro",
        params: { ...dadosAssociado },
      });
    }
  };

  return (
    <View style={{ gap: 24 }}>
      {mensagem && (
        <TouchableOpacity
          onPress={handleMensagemPress}
          style={{
            backgroundColor: mensagem.tipo === "success" ? "#DBEAFE" : "#FECACA",
            padding: 16,
            borderRadius: 8,
            flexDirection: "row",
            alignItems: "center",
            gap: 8,
          }}
        >
          <Ionicons
            name={mensagem.tipo === "success" ? "checkmark-circle" : "alert-circle"}
            size={24}
            color={mensagem.tipo === "success" ? "#2563EB" : "#DC2626"}
          />
          <Text
            style={{
              color: mensagem.tipo === "success" ? "#2563EB" : "#DC2626",
              fontSize: 16,
              fontWeight: "500",
            }}
          >
            {mensagem.texto}
          </Text>
        </TouchableOpacity>
      )}

      {formVisivel && (
        <FormSection
          title="Cadastro de Associado"
          fields={fields}
          initialValues={initialValues}
          onSubmit={handleSubmit}
          isSubmit={true}
        />
      )}
    </View>
  );
}
