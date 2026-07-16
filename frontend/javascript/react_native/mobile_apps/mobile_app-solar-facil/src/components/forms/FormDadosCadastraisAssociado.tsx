import React, { useState, useEffect } from "react";
import { useRouter } from "expo-router";
import { View, Text } from "react-native";

import { FormSection } from "@/components/forms/FormSection";
import { AssociadoType } from "@/types/AssociadoType";
import { FieldDefinitionType } from "@/types/FieldDefinitionType";
import { validateFormData } from "@/utils/validates/validateFormData";
import { fetchPlanOptions } from "@/services/servicePlans";
import { fetchConcessionariasOptions } from "@/services/serviceConcessionarias";
import { fetchConsumoMedioOptions } from "@/services/serviceConsumoMedio";
import { isValidCPF } from "@/utils/validators/validatorCPF";
import { isValidCNPJ } from "@/utils/validators/validatorCNPJ";
import { brazilianStates } from "@/constants/states";
import { useAuth } from "@/context/AuthContext";
import { useAssociadosCopilot } from "@/services/database/useAssociadosCopilot";

interface FormCadastroDadosCadastraisAssociadoProps {
  associado: AssociadoType;
  onSubmit?: (data: AssociadoType) => void;
  editable?: boolean;
}

export const FormDadosCadastraisAssociado: React.FC<FormCadastroDadosCadastraisAssociadoProps> = ({
  associado,
  onSubmit,
  editable
}) => {
  const router = useRouter();
  const [mensagem, setMensagem] = useState<string | null>(null);
  const { updatelogin } = useAuth();

  const useAssociados = useAssociadosCopilot(); // Call the hook here

  // Carrega as opções de planos, concessionárias e consumo médio
  const [planOptions, setPlanOptions] = useState<{ label: string; value: number }[]>([]);
  const [concessionariasOptions, setConcessionariasOptions] = useState<{ label: string; value: number }[]>([]);
  const [consumoMedioOptions, setConsumoMedioOptions] = useState<{ label: string; value: number }[]>([]);

  const [erros, setErros] = useState<{ [key: string]: string }>({});

  useEffect(() => {
    fetchPlanOptions().then(setPlanOptions);
    fetchConcessionariasOptions().then(setConcessionariasOptions);
    fetchConsumoMedioOptions().then(setConsumoMedioOptions);
  }, []);

  console.log("FormCadastroDadosCadastraisAssociado - VAI SETAR OS FIELDS - associado:", associado);
  const editableFields = editable !== undefined ? editable : true; // Default to true if not provided
  const fields: FieldDefinitionType<AssociadoType>[] = [
    // campos não editaveis
    {
      // Valor padrão caso associado.id seja undefined ou null
      defaultValue: associado.id
        ? associado.id.toString().padStart(15, "0").replace(/(\d{13})(\d{2})$/, "$1-$2")
        : "00000000000000-00",
      editable: false,
      name: "id",
      label: "id",
      type: "number",
      placeholder: "ID do associado",
    },
    {
      defaultValue: associado.dataCadastro
        ? new Date(associado.dataCadastro).toLocaleDateString("pt-BR", {
          day: "2-digit",
          month: "2-digit",
          year: "numeric",
        })
        : "", // Valor padrão caso associado.dataCadastro seja undefined ou null
      editable: false,
      name: "dataCadastro",
      label: "Data de Cadastro",
      type: "text",
      placeholder: "Data de cadastro do associado",
      formattedValue: (value) => new Date(value).toLocaleDateString("pt-BR"),
    },
    {
      // setar a data no formato brasileiro dd/mm/aaaa
      defaultValue: associado.dataAtualizacao
        ? new Date(associado.dataAtualizacao).toLocaleDateString("pt-BR", {
          day: "2-digit",
          month: "2-digit",
          year: "numeric",
        })
        : "", // Valor padrão caso associado.dataCadastro seja undefined ou null
      editable: false,
      name: "dataAtualizacao",
      label: "Última Atualização",
      type: "text",
      placeholder: "Data da última atualização",
      formattedValue: (value) => new Date(value).toLocaleDateString("pt-BR"),
    },
    {
      defaultValue: associado.status,
      editable: false,
      name: "status",
      label: "Status",
      type: "select",
      placeholder: "Status do associado",
      options: [
        { label: "Em cadastro", value: "Em cadastro" },
        { label: "Ativo", value: "Ativo" },
        { label: "Inativo", value: "Inativo" },
        { label: "Bloqueado", value: "Bloqueado" },
        { label: "Encerrado", value: "Encerrado" },
      ],
      required: true,
      validation: (value) => {
        if (!value) return "Status é obrigatório.";
        const validStatuses = ["Em cadastro", "Ativo", "Inativo", "Bloqueado", "Encerrado"];
        if (!validStatuses.includes(value)) return "Status inválido.";
        return null;
      }
    },
    {
      defaultValue: (associado.tipoPessoa !== "Pessoa Jurídica" ? "Pessoa Física" : "Pessoa Jurídica"),
      editable: false,
      name: "tipoPessoa",
      label: "Tipo de Pessoa",
      type: "select",
      placeholder: "Tipo de Pessoa",
      options: [
        { label: "Pessoa Física", value: "Pessoa Física" },
        { label: "Pessoa Jurídica", value: "Pessoa Jurídica" },
      ],
      required: true,
      validation: (value) => {
        if (!value) return "Tipo de pessoa é obrigatório.";
        return null;
      }
    },
    {
      defaultValue: associado.cpf_cnpj,
      editable: false,
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
      defaultValue: associado.tipoAssociado,
      name: "tipoAssociado",
      label: "Tipo Associado",
      type: "select",
      editable: false,
      options: [
        { label: "Fornecedor", value: "Fornecedor" },
        { label: "Beneficiado", value: "Beneficiado" },
        { label: "Hibrido", value: "Hibrido" },
      ],
      required: true,
      validation: (value) => {
        if (!value) return "Tipo de associado é obrigatório.";
        const validTypes = ["Fornecedor", "Beneficiado", "Hibrido"];
        if (!validTypes.includes(value)) return "Tipo de associado inválido.";
        return null;
      },
      placeholder: "Selecione o tipo de associado",
    },


    // campos editaveis 
    {
      defaultValue: associado.nome,
      editable: editableFields,
      name: "nome",
      label: "Nome completo ou Razão Social",
      placeholder: "João Silva",
      type: "text",
      required: true,
      toUpperCase: true,
    },
    {
      defaultValue: associado.telefone,
      editable: editableFields,
      required: true,
      name: "telefone",
      label: "Telefone",
      placeholder: "(11) 99999-0000",
      type: "text",
      validation: (value) => {
        if (!value) return "Telefone é obrigatório";
        const phoneRegex = /^\d{2}\d{5}\d{4}$/;
        if (!phoneRegex.test(value)) return "Telefone inválido. Numero dever ter 10 ou 11 digitos";
        return null;
      }
    },
    {
      defaultValue: associado.email,
      editable: editableFields,
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

    // Campos de endereço
    {
      defaultValue: associado.cep,
      editable: editableFields,
      name: "cep",
      label: "CEP",
      placeholder: "00000-000",
      type: "text",
      required: true,
      keyboardType: "numeric",
      validation: (value) => {
        if (!value) return "CEP é obrigatório";
        const cepRegex = /^\d{5}\d{3}$/;
        if (!cepRegex.test(value)) return "CEP inválido. Formato esperado: 00000000";
        return null;
      }
    },
    {
      defaultValue: associado.endereco,
      editable: editableFields,
      name: "endereco",
      label: "Endereço",
      placeholder: "Rua Exemplo, 123",
      type: "text",
      required: true,
      toUpperCase: true,
    },
    {
      defaultValue: associado.numero,
      editable: editableFields,
      name: "numero",
      label: "Número",
      placeholder: "123",
      type: "text",
      required: true,
    },
    {
      defaultValue: associado.bairro,
      editable: editableFields,
      name: "bairro",
      label: "Bairro",
      placeholder: "Informe o bairro",
      type: "text",
      required: true,
      toUpperCase: true,
    },
    {
      defaultValue: associado.cidade,
      editable: editableFields,
      name: "cidade",
      label: "Cidade",
      placeholder: "Informe a cidade",
      type: "text",
      required: true,
      toUpperCase: true,
    },
    {
      defaultValue: associado.estado,
      editable: editableFields,
      name: "estado",
      label: "Estado",
      placeholder: "São Paulo",
      type: "select",
      required: true,
      keyboardType: "default",
      options: brazilianStates,
    },
    {
      defaultValue: associado.complemento,
      editable: editableFields,
      name: "complemento",
      label: "Complemento (opcional)",
      placeholder: "Apto 101, Bloco B, etc.",
      type: "text",
      toUpperCase: true,
    },

    // Campos adicionais
    {
      defaultValue: associado.tipoConexao,
      editable: editableFields,
      name: "tipoConexao",
      label: "Tipo de Conexão",
      placeholder: "Selecione o tipo",
      type: "select",
      required: true,
      options: [
        { label: "Monofásica", value: "monofasica" },
        { label: "Bifásica", value: "bifasica" },
        { label: "Trifásica", value: "trifasica" },
      ],
    },
    {
      defaultValue: associado.aceitaTermos || "Não", // Valor padrão
      editable: editableFields,
      name: "aceitaTermos",
      label: "Aceita os termos e condições?",
      type: "select",
      placeholder: "Selecione uma opção",
      required: true,
      options: [
        { label: "Sim", value: "Sim" },
        { label: "Não", value: "Não" },
      ],
    },
    {
      defaultValue: associado.observacoes,
      editable: editableFields,
      name: "observacoes",
      label: "Observações (opcional)",
      placeholder: "",
      type: "textarea", // Campo de texto grande
    },

    // campos específicos de pessoa física
    {
      defaultValue: associado.dataNascimento
        ? new Date(associado.dataNascimento).toLocaleDateString("pt-BR", {
          day: "2-digit",
          month: "2-digit",
          year: "numeric",
        })
        : "", // Valor padrão caso associado.dataCadastro seja undefined ou null
      editable: editableFields,
      name: "dataNascimento",
      label: "Data de Nascimento",
      placeholder: "DD/MM/AAAA",
      type: "date",
      required: true,
      formattedValue: (value) => new Date(value).toLocaleDateString("pt-BR"),
    },
    {
      defaultValue: associado.nomeSocial,
      editable: editableFields,
      name: "nomeSocial",
      label: "Nome Social (opcional)",
      placeholder: "Nome social do associado",
      type: "text",
      toUpperCase: true,
    },

    // campos específicos de pessoa jurídica
    {
      defaultValue: associado.razaoSocial,
      editable: editableFields,
      name: "razaoSocial",
      label: "Razão Social",
      placeholder: "Razão social da empresa",
      type: "text",
      toUpperCase: true,
    },
    {
      defaultValue: associado.nomeFantasia,
      editable: editableFields,
      name: "nomeFantasia",
      label: "Nome Fantasia (opcional)",
      placeholder: "Nome fantasia da empresa",
      type: "text",
      toUpperCase: true,
    },
    {
      defaultValue: associado.dataAbertura
        ? new Date(associado.dataAbertura).toLocaleDateString("pt-BR", {
          day: "2-digit",
          month: "2-digit",
          year: "numeric",
        })
        : "", // Valor padrão caso associado.dataCadastro seja undefined ou null
      editable: editableFields,
      name: "dataAbertura",
      label: "Data de Abertura",
      placeholder: "DD/MM/AAAA",
      type: "date",
      required: true,
      formattedValue: (value) => new Date(value).toLocaleDateString("pt-BR"),
    },

    // dados especificos de beneficiado
    {
      editable: editableFields,
      name: "nomeConcessionaria",
      label: "Informe a concessionaria",
      type: "select",
      required: true,
      placeholder: "Selecione a concessionária",
      options: concessionariasOptions.map(option => ({
        ...option,
        value: String(option.value),
      })),
    },
    {
      editable: editableFields,
      name: "consumoMedio",
      label: "Consumo médio mensal (kWh)",
      type: "select",
      required: true,
      placeholder: "Selecione o consumo médio",
      options: consumoMedioOptions.map(option => ({
        ...option,
        value: String(option.value),
      })),
    },
    {
      editable: editableFields,
      name: "planoDesejado",
      label: "Plano desejado",
      placeholder: "Selecione um plano",
      type: "select",
      required: true,
      options: planOptions.map(option => ({
        ...option,
        value: String(option.value),
      })),
    },

    // campos específicos de fornecedor
    {
      defaultValue: associado.potenciaInstalada,
      editable: editableFields,
      name: "potenciaInstalada",  // Campo de potência instalada
      label: "Potência Instalada (kWh)",
      placeholder: "Potência instalada",  // Placeholder para potência instalada
      type: "text", // Tipo numérico para potência
      keyboardType: "numeric", // Teclado numérico
      required: true, // Campo obrigatório
    },
    {
      defaultValue: associado.disponibilidade,
      editable: editableFields,
      name: "disponibilidade",
      label: "Disponibilidade de Energia (kWh)",
      placeholder: "Ex: 150",
      keyboardType: "numeric",
      required: true,
      type: "number",
    },
  ];

  const handleSubmit = async (data: AssociadoType) => {
    console.log("FormCadastroDadosCadastraisAssociado - handleSubmit - Iniciando ... :", data);
    const erros = validateFormData(data, fields);
    if (Object.keys(erros).length > 0) {
      console.log("FormCadastroDadosCadastraisAssociado - handleSubmit - erros:", erros);
      setMensagem("Por favor, corrija os erros antes de salvar.");
      return;
    }

    const dadosAtualizados = {
      ...associado,
      ...data,
      senha: associado.senha, // mantém senha original
    };

    try {
      console.log("FormCadastroDadosCadastraisAssociado - Vai chamar o serviço de atualização:", dadosAtualizados);

      // await atualizarAssociado(dadosAtualizados);
      const result = (await useAssociados).updateRecord(dadosAtualizados);
      if ((await result).success === false) {
        console.error("FormCadastroDadosCadastraisAssociado - Erro ao atualizar os dados:", (await result).error);
        setMensagem("Erro ao atualizar os dados.");
        return;
      }

      console.log("FormCadastroDadosCadastraisAssociado - Dados atualizados com sucesso: ", result, "==> ", dadosAtualizados);
      setMensagem("Dados atualizados com sucesso!");

      updatelogin(dadosAtualizados); // Atualiza o associado no contexto de autenticação

      // agora precisamos limpar as mensagens de erros de todos os campos
      setErros({});

      if (onSubmit) onSubmit(dadosAtualizados);
      console.log("FormCadastroDadosCadastraisAssociado - Dados enviados com sucesso:", dadosAtualizados);

      // se recebeu o parametro de editable indica que veio no modo de edição
      if (editable !== undefined && editable === true) {
        // se veio no modo de edição, então vamos voltar para a tela anterior
        console.log("FormCadastroDadosCadastraisAssociado - Navegando para a tela anterior");
        router.back(); // Descomente se estiver usando React Router
      }
    } catch (error) {
      console.error("FormCadastroDadosCadastraisAssociado - Erro ao atualizar os dados.", error);
      setMensagem("Erro ao atualizar os dados.");
    }
  };

  return (
    <View style={{ gap: 0 }}>
      {mensagem && <Text style={{ color: "blue", marginBottom: 12 }}>{mensagem}</Text>}
      <FormSection
        fields={fields}
        title="Dados Cadastrais do Associado"
        initialValues={associado}
        onSubmit={handleSubmit}
        isSubmit={true}
      />
    </View>
  );
};
