export type ConsumoMedioType = {
  id: number;
  name: string;
  initialValue: number;
  finalValue: number;
  unit: string;
  description: string;
  status: "active" | "inactive";
};
