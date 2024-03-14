import { signIn } from "@/api";
import { Alert } from "@/components";
import { type JSXEvent, MessageType } from "@/domain";
import { useForm } from "@/hooks";
import { Lock, Mail } from "@/icons";
import { useAppState } from "@/store";
import { saveToken } from "@/utils";
import { emailValidator, minLenValidator, nonEmptyValidator } from "@/utils/validators";
import { A, useNavigate } from "@solidjs/router";
import { Show, createSignal } from "solid-js";

export default function SignIn() {
	const navigate = useNavigate();
	const { handleApiError, setUser } = useAppState();
	const [errMsg, setErrMsg] = createSignal("");
	const { form, isValid } = useForm({
		email: {
			validators: [nonEmptyValidator, emailValidator],
		},
		password: {
			validators: [nonEmptyValidator, minLenValidator(6)],
		},
	});

	const closeErrMsg = () => setErrMsg("");

	const handleSubmit = (event: JSXEvent<SubmitEvent, HTMLFormElement>) => {
		event.preventDefault();
		if (!isValid()) {
			return;
		}
		signIn({ email: form.email.value(), password: form.password.value() })
			.then((data) => {
				saveToken(data.token);
				setUser(data.user);
				navigate("/home", { replace: true });
			}).catch((err) => {
				handleApiError(err, (msg) => setErrMsg(msg));
			});
	}

	return (
		<div class="flex flex-col items-center">
			<h1 class="text-3xl scroll-m-20 font-extrabold tracking-tight mb-2">Bienvenido a OLIAI</h1>

			<Show when={errMsg().length > 0}>
				<Alert message={errMsg()} variant={MessageType.WARNING} onClose={closeErrMsg} />
			</Show>

			<form class="self-stretch flex flex-col gap-2 mt-2" onSubmit={handleSubmit}>
				<label class="form-control">
					<div class="label">
						<span class="label-text">Correo</span>
					</div>
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
					<div class="label">
						<span class="label-text">Contraseña</span>
					</div>
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
					Ingresar
				</button>
				<p class="self-end text-base mt-2">
					¿Aún no tienes una cuenta?
					<A href="/auth/signup" class="link link-primary font-semibold"> Registrarme</A>
				</p>
			</form>
			<div class="h-2 md:h-32" />
		</div >
	);
}
