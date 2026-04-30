import { useQuery } from "@tanstack/react-query";
import { toast } from "sonner";

export type Plan = {
  id?: string | number;
  title?: string;
  name?: string;
  [key: string]: unknown; // Allow additional properties without TypeScript errors
};

const queryKey = ["plans"];

export function usePlansQuery() {
  return useQuery({
    queryKey: queryKey,
    queryFn: getPlans,
    retry: false,
  });
}

async function getPlans(): Promise<Plan[]> {
  const response = await fetch(import.meta.env.VITE_BASE + "/plans", {
    method: "GET",
    headers: {
      Accept: "application/json",
    },
  });

  if (!response.ok) {
    toast.error(`Failed to fetch plans: ${response.status}`, {
      id: queryKey.join("-"), // Unique ID to prevent duplicate toasts
    });
    throw new Error(`Failed to fetch plans: ${response.status}`);
  }

  const payload: unknown = await response.json();
  return toPlanArray(payload);
}

function toPlanArray(payload: unknown): Plan[] {
  if (Array.isArray(payload)) {
    return payload as Plan[];
  }

  if (payload && typeof payload === "object" && "data" in payload) {
    const data = (payload as { data?: unknown }).data;
    if (Array.isArray(data)) {
      return data as Plan[];
    }
  }

  throw new Error("Invalid plans response format.");
}
