import React from "react";
import { InputText } from "@/components/inputs/InputText";
import { InputDate } from "@/components/inputs/InputDate";
import { InputRadio } from "@/components/inputs/InputRadio";
import { InputSelect } from "@/components/inputs/InputSelect";
import { InputSwitch } from "@/components/inputs/InputSwitch";
import { InputTextarea } from "@/components/inputs/InputTextarea";
import { InputPasswordWithToggle } from "@/components/inputs/InputPasswordWithToggle";

interface DynamicInputProps {
  type: string;
  label?: string;
  placeholder?: string;
  value?: string | number | boolean | Date | null | undefined;
  onChange?: (value: any) => void;
  options?: { label: string; value: string }[];
  error?: string;
  editable?: boolean;
}

export const DynamicInput: React.FC<DynamicInputProps> = ({
  type,
  label,
  placeholder,
  value,
  onChange,
  options,
  error,
  editable = true, // Default value for editable
}) => {

  // Renderiza um componente específico com base no tipo
  switch (type) {
    case "date":
      // Garante que value seja um objeto Date ou undefined
      const dateValue = value ? new Date(value as string | number | Date) : "";
      console.log("DynamicInput - label:", label, " - value:", value, " - dateValue:", dateValue);

      return (
        <InputDate
          label={label || ""}
          value={value as Date}
          onChangeText={onChange ?? (() => { })}
          error={error}
          editable={editable}
        />
      );

    case "radio":
      return (
        <InputRadio
          label={label || ""}
          value={value as string}
          onChangeText={onChange ?? (() => { })}
          options={options || []}
          error={error}
          editable={editable}
        />
      );

    case "select":
      return (
        <InputSelect
          label={label || ""}
          value={value as string}
          onChangeText={onChange ?? (() => { })}
          options={options || []}
          placeholder={placeholder}
          error={error}
          editable={editable}
        />
      );

    case "switch":
      return (
        <InputSwitch
          label={label || ""}
          value={value as boolean}
          onChangeText={onChange ?? (() => { })}
          error={error}
          editable={editable}
        />
      );

    case "textarea":
      return (
        <InputTextarea
          label={label || ""}
          value={value as string}
          onChangeText={onChange ?? (() => { })}
          placeholder={placeholder}
          error={error}
          editable={editable}
        />
      );
    case "password":
      return (
        <InputPasswordWithToggle
          label={label}
          placeholder={placeholder}
          value={value as string}
          onChangeText={onChange}
          error={error}
          editable={editable}
        />
      );

    default:
      // Default para InputText
      return (
        console.log("DynamicInput - InputText - Default para InputText { label, value } - ", { label, value }),
        <InputText
          label={label}
          placeholder={placeholder}
          value={value as string}
          onChangeText={onChange}
          error={error}
          options={options}
          editable={editable}
        />
      );
  }
};
