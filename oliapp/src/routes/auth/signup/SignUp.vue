<script setup lang="ts">
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { RouterLink } from "vue-router";
</script>
<template>
	<div class="flex flex-col items-center">
		<h1 class="text-3xl scroll-m-20 font-extrabold tracking-tight">Bienvenido a OLI AI</h1>

		<form class="self-stretch flex flex-col gap-2 mt-4">
			<div class="">
				<Label for="first_name">Nombre(s)</Label>
				<Input id="first_name" type="text" placeholder="" autocomplete="off" />
			</div>
			<div class="">
				<Label for="last_name">Apellido(s)</Label>
				<Input id="last_name" type="text" placeholder="" autocomplete="off" />
			</div>
			<div class="">
				<Label for="email">Correo</Label>
				<Input id="email" type="email" placeholder="ejemplo@mail.com" autocapitalize="off" autocomplete="email" />
			</div>
			<div class="">
				<Label for="password">Contraseña</Label>
				<Input id="password" type="password" placeholder="" />
			</div>
			<Button class="font-bold text-base mt-2">Registrarme</Button>
			<p class="self-end text-base mt-2">
				¿Ya tienes una cuenta?
				<RouterLink :to="{ name: 'signin' }" class="underline underline-offset-4">Ingresar</RouterLink>
			</p>
		</form>
		<div class="h-2 md:h-32"></div>
	</div>
</template>
