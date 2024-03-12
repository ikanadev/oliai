import { Mail, Lock, User } from "@/icons";
import { A } from "@solidjs/router";
import { createStore, unwrap } from "solid-js/store";

type JSXEvent<Ev, El> = Ev & { currentTarget: El, target: Element };

export default function SignUp() {
	const [form, setForm] = createStore({
		firstName: "",
		lastName: "",
		email: "",
		password: "",
	});

	const handleChange = (field: string, event: JSXEvent<InputEvent, HTMLInputElement>) => {
		setForm({ [field]: event.currentTarget.value });
	};

	const handleSubmit = (event: JSXEvent<SubmitEvent, HTMLFormElement>) => {
		event.preventDefault();

		console.log(unwrap(form));
	}

	return (
		<div class="flex flex-col items-center">
			<h1 class="text-3xl scroll-m-20 font-extrabold tracking-tight">Crear una cuenta</h1>

			<form class="self-stretch flex flex-col gap-2 mt-4" onSubmit={handleSubmit}>
				<label class="form-control">
					<div class="label"><span class="label-text">Nombre(s)</span></div>
					<span class="input input-bordered flex items-center gap-2">
						<User class="text-lg" />
						<input
							value={form.firstName}
							onInput={[handleChange, "firstName"]}
							type="text"
							placeholder="Ej. Pedro"
							autocomplete="off"
							class="grow"
							required
						/>
					</span>
				</label>
				<label class="form-control">
					<div class="label"><span class="label-text">Apellido(s)</span></div>
					<span class="input input-bordered flex items-center gap-2">
						<User class="text-lg" />
						<input
							value={form.lastName}
							onInput={[handleChange, "lastName"]}
							type="text"
							placeholder="Ej. Pérez"
							autocomplete="off"
							class="grow"
							required
						/>
					</span>
				</label>
				<label class="form-control">
					<div class="label"><span class="label-text">Correo</span></div>
					<span class="input input-bordered flex items-center gap-2">
						<Mail class="text-lg" />
						<input
							value={form.email}
							onInput={[handleChange, "email"]}
							type="email"
							placeholder="ejemplo@mail.com"
							autocapitalize="off"
							autocomplete="email"
							class="grow"
							required
						/>
					</span>
				</label>
				<label class="form-control">
					<div class="label"><span class="label-text">Contraseña</span></div>
					<span class="input input-bordered flex items-center gap-2">
						<Lock class="text-lg" />
						<input
							value={form.password}
							onInput={[handleChange, "password"]}
							type="password"
							placeholder="* * * * * * *"
							required
						/>
					</span>
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
