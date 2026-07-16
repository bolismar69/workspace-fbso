// /src/types/PlanType.ts
import { Ionicons } from "@expo/vector-icons";

export type PlanType = {
  id: number;
  name: string;
  icon: React.ComponentProps<typeof Ionicons>["name"];
  powerRange: string;
  consumption: string;
  pricePerKwh: string;
  monthlyEstimate: string;
  energyCost: string;
  commercialIndication: boolean;
  description?: string;
  commercialAttraction: string;
  status: "active" | "inactive";
};
