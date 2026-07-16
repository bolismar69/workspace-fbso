// /src/types/DBResponse.ts

export type DBResponse<T> = {
  success: boolean;
  data?: T;
  error?: string;
};
