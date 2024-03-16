import type { JSXEvent } from "@/domain";
import { ArrowBack } from "@/icons";
import { A, useNavigate } from "@solidjs/router";
import { useForm } from "@/hooks";
import { nonEmptyValidator, minLenValidator } from "@/utils/validators";
import { useCreateCompany } from "@/api";

export default function AddCompany() {
	const navigate = useNavigate();
	const companyMutation = useCreateCompany();
	const { form, isValid } = useForm({
		name: {
			validators: [nonEmptyValidator, minLenValidator(3)],
		},
		logoUrl: {
			validators: [nonEmptyValidator, minLenValidator(3)],
		}
	});

	const handleSubmit = (event: JSXEvent<SubmitEvent, HTMLFormElement>) => {
		event.preventDefault();
		if (!isValid()) {
			return;
		}
		companyMutation.mutate(
			{ name: form.name.value(), logoUrl: form.logoUrl.value() },
			{ onSuccess: () => navigate("/companies") },
		);
	}
	return (
		<div class="">
			<div class="flex items-center gap-2">
				<A href="/companies" class="btn btn-ghost btn-sm"><ArrowBack class="text-3xl" /></A>
				<h1 class="font-extrabold text-3xl">NUEVA EMPRESA</h1>
			</div>
			<form class="self-stretch flex flex-col mt-2" onSubmit={handleSubmit}>
				<label class="form-control">
					<div class="label">
						<span class="label-text">Nombre</span>
					</div>
					<input
						value={form.name.value()}
						onInput={form.name.onInput}
						autocomplete="off"
						class="input input-bordered"
						required
					/>
					<div class="label">
						<span class="label-text-alt text-error">{form.name.error()}</span>
					</div>
				</label>
				<label class="form-control">
					<div class="label">
						<span class="label-text">Logo Url</span>
					</div>
					<input
						value={form.logoUrl.value()}
						onInput={form.logoUrl.onInput}
						class="input input-bordered"
						autocomplete="off"
						required
					/>
					<div class="label">
						<span class="label-text-alt text-error">{form.logoUrl.error()}</span>
					</div>
				</label>
				<button type="submit" class="btn btn-primary mt-2">
					Guardar
				</button>
			</form>
		</div>
	);
}
