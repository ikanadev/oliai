import type { CompanyWithTimeData } from "@/domain";
import { createMutation, createQuery } from "@tanstack/solid-query";
import { api } from "./mande";
import { queryClient } from "./queryClient";

export const companyQueryKeys = {
	companies: () => ["companies"],
};

export function useCompanies() {
	return createQuery(() => ({
		queryKey: companyQueryKeys.companies(),
		queryFn: () => api.get<CompanyWithTimeData[]>("/api/admin/companies"),
	}));
}

export function useCreateCompany() {
	return createMutation(() => ({
		mutationFn: (data: {
			name: string;
			logoUrl: string;
		}) => api.post<void>("/api/admin/companies", data),
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: companyQueryKeys.companies() });
		},
	}))
}
