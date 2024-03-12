import { Mail, Lock } from "@/icons";
import { A } from "@solidjs/router";

export default function SignIn() {
	return (
		<div class="flex flex-col items-center">
			<h1 class="text-3xl scroll-m-20 font-extrabold tracking-tight">Bienvenido a OLIAI</h1>

			<form class="self-stretch flex flex-col gap-2 mt-4">
				<label class="form-control">
					<div class="label">
						<span class="label-text">Correo</span>
					</div>
					<span class="input input-bordered flex items-center gap-2">
						<Mail class="text-lg" />
						<input
							type="email"
							placeholder="ejemplo@mail.com"
							autocapitalize="off"
							autocomplete="email"
							class="grow"
						/>
					</span>
				</label>
				<label class="form-control">
					<div class="label">
						<span class="label-text">Contraseña</span>
					</div>
					<span class="input input-bordered flex items-center gap-2">
						<Lock class="text-lg" />
						<input type="password" placeholder="* * * * * * *" />
					</span>
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
