<script setup lang="ts">
import { RouterView } from "vue-router";
</script>
<template>
	<div class="min-h-screen flex">
		<div class="flex-1 hidden md:block">
			<img src="/chatbot.svg" class="h-full object-fill" />
		</div>
		<div class="flex-1 flex items-start md:items-center justify-center p-4">
			<div class="w-full sm:w-[400px] max-w-full mt-20 md:mt-0">
				<RouterView />
			</div>
		</div>
	</div>
</template>
