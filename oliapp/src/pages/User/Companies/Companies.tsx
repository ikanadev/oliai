import { createEffect, Show, For } from "solid-js";
import { useCompanies } from "@/api";
import { useAppState } from "@/store";
import { EmptyMessage } from "@/components";
import { A } from "@solidjs/router";
import { Add } from "@/icons";

export default function Companies() {
	const { handleApiError } = useAppState();
	const companiesQuery = useCompanies();

	companiesQuery.data;

	createEffect(() => {
		if (companiesQuery.error) {
			handleApiError(companiesQuery.error);
		}
	});
	return (
		<div class="">
			<div class="flex justify-between items-center">
				<h1 class="text-3xl font-extrabold">EMPRESAS</h1>
				<A href="/companies/add" class="btn btn-primary btn-sm" type="button">
					<Add class="text-xl" /> Agregar
				</A>
			</div>
			<Show when={companiesQuery.data}>
				{(companies) => (
					<>
						<For each={companies()}>
							{(company) => (
								<p>{company.name}</p>
							)}
						</For>
						<Show when={companies().length === 0}>
							<EmptyMessage message="No hay empresas." />
						</Show>
					</>
				)}
			</Show>
		</div>
	);
}
