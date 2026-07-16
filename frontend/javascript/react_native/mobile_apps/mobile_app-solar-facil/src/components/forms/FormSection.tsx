import { DynamicInput } from "@/components/inputs/DynamicInput";
import { View, Text, TouchableOpacity } from "react-native";
import { Controller, useForm, DefaultValues, Path } from "react-hook-form";
import { FieldDefinitionType } from "@/types/FieldDefinitionType";
import { useAppTheme } from "@/context/AppThemeContext";
import { validateFormData } from "@/utils/validates/validateFormData";

interface FormSectionProps<T extends Record<string, any>> {
  title?: string;
  fields: FieldDefinitionType<T>[];
  initialValues: DefaultValues<T>;
  onSubmit?: (data: T) => void;
  isSubmit?: boolean;
}

export function FormSection<T extends Record<string, any>>({
  title,
  fields,
  initialValues,
  onSubmit,
  isSubmit,
}: FormSectionProps<T>) {
  const { theme } = useAppTheme();

  const {
    control,
    setError,
    getValues,
    formState: { errors },
  } = useForm<T>({ defaultValues: initialValues });

  const handleValidationAndSubmit = () => {
    const formData = getValues();
    const validationErrors = validateFormData(formData, fields);

    if (Object.keys(validationErrors).length > 0) {
      Object.entries(validationErrors).forEach(([fieldName, message]) => {
        setError(fieldName as Path<T>, { message });
      });
      return;
    }

    onSubmit?.(formData);
  };

  return (
    <View style={{ gap: 16 }}>
      {title && <Text style={[theme.title]}>{title}</Text>}
      {fields.map((field) => (
        <Controller
          key={field.name as string}
          control={control}
          name={field.name as Path<T>}
          rules={{
            required: field.required ? `${field.label} é obrigatório` : undefined,
            pattern: field.pattern
              ? { value: field.pattern, message: field.errorMessage || "Formato inválido" }
              : undefined,
            validate: field.validation
              ? (value) => {
                const allValues = getValues();
                const result = field.validation?.(value, allValues);
                return result || true;
              }
              : undefined,
          }}
          render={({ field: { onChange, onBlur, value } }) => (
            <DynamicInput
              type={field.type}
              label={field.label}
              placeholder={field.placeholder}
              // value={value}
              value={field.defaultValue || value}
              onChange={onChange}
              options={field.options}
              error={errors?.[field.name as keyof T]?.message as string}
              editable={field.editable}
            />
          )}
        />
      ))}

      {onSubmit && isSubmit && (
        <TouchableOpacity
          style={[theme.button, { marginTop: 16 }]}
          onPress={handleValidationAndSubmit}
        >
          <Text style={theme.buttonText}>Enviar</Text>
        </TouchableOpacity>
      )}
    </View>
  );
}
