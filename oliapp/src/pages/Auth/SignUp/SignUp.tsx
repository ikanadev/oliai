import { signUp } from "@/api";
import { Alert } from "@/components";
import type { ApiError, JSXEvent } from "@/domain";
import { MessageType } from "@/domain";
import { useForm } from "@/hooks";
import { Lock, Mail, User } from "@/icons";
import { isApiError } from "@/utils";
import { emailValidator, minLenValidator, nonEmptyValidator } from "@/utils/validators";
import { A } from "@solidjs/router";
import { Show, createSignal } from "solid-js";

export default function SignUp() {
	const [errMsg, setErrMsg] = createSignal("");
	const { form, isValid } = useForm({
		firstName: { validators: [nonEmptyValidator, minLenValidator(3)] },
		lastName: { validators: [nonEmptyValidator, minLenValidator(3)] },
		email: { validators: [nonEmptyValidator, emailValidator] },
		password: { validators: [nonEmptyValidator, minLenValidator(6)] },
	});

	const closeErrMsg = () => setErrMsg("");

	const handleSubmit = (event: JSXEvent<SubmitEvent, HTMLFormElement>) => {
		event.preventDefault();
		if (!isValid()) {
			return;
		}
		const data = {
			firstName: form.firstName.value(),
			lastName: form.lastName.value(),
			email: form.email.value(),
			password: form.password.value(),
		};
		signUp(data).then((res) => {
			console.log(res);
		}).catch((err) => {
			if (isApiError(err.body)) {
				setErrMsg((err.body as ApiError).message);
				return;
			}
			// Handle unexpected error
			console.log(err);
			console.log(err.body);
			console.log(err.response);
		});
	}

	return (
		<div class="flex flex-col items-center">
			<h1 class="text-3xl scroll-m-20 font-extrabold tracking-tight mb-2">Crear una cuenta</h1>
			<Show when={errMsg().length > 0}>
				<Alert message={errMsg()} variant={MessageType.WARNING} onClose={closeErrMsg} />
			</Show>

			<form class="self-stretch flex flex-col mt-2" onSubmit={handleSubmit}>
				<label class="form-control">
					<div class="label"><span class="label-text">Nombre(s)</span></div>
					<span class="input input-bordered flex items-center gap-2">
						<User class="text-lg" />
						<input
							value={form.firstName.value()}
							onInput={form.firstName.onInput}
							type="text"
							placeholder="Ej. Pedro"
							autocomplete="off"
							class="grow"
							required
						/>
					</span>
					<div class="label">
						<span class="label-text-alt text-error">{form.firstName.error()}</span>
					</div>
				</label>
				<label class="form-control">
					<div class="label"><span class="label-text">Apellido(s)</span></div>
					<span class="input input-bordered flex items-center gap-2">
						<User class="text-lg" />
						<input
							value={form.lastName.value()}
							onInput={form.lastName.onInput}
							type="text"
							placeholder="Ej. Pérez"
							autocomplete="off"
							class="grow"
							required
						/>
					</span>
					<div class="label">
						<span class="label-text-alt text-error">{form.lastName.error()}</span>
					</div>
				</label>
				<label class="form-control">
					<div class="label"><span class="label-text">Correo</span></div>
					<span class="input input-bordered flex items-center gap-2">
						<Mail class="text-lg" />
						<input
							value={form.email.value()}
							onInput={form.email.onInput}
							type="email"
							placeholder="ejemplo@mail.com"
							autocapitalize="off"
							autocomplete="email"
							class="grow"
							required
						/>
					</span>
					<div class="label">
						<span class="label-text-alt text-error">{form.email.error()}</span>
					</div>
				</label>
				<label class="form-control">
					<div class="label"><span class="label-text">Contraseña</span></div>
					<span class="input input-bordered flex items-center gap-2">
						<Lock class="text-lg" />
						<input
							value={form.password.value()}
							onInput={form.password.onInput}
							type="password"
							placeholder="* * * * * * *"
							required
						/>
					</span>
					<div class="label">
						<span class="label-text-alt text-error">{form.password.error()}</span>
					</div>
				</label>
				<button type="submit" class="btn btn-primary mt-2">
					Registrarme
				</button>
				<p class="text-sm">
					Al crear una cuenta, aceptas nuestros
					<A href="#" class="link link-primary font-semibold"> Términos y Condiciones </A>
					de uso.
				</p>
				<p class="self-end text-base mt-2">
					¿Ya tienes una cuenta?
					<A href="/auth/signin" class="link link-primary font-semibold"> Ingresar</A>
				</p>
			</form>
			<div class="h-2 md:h-32" />
		</div>
	);
}
