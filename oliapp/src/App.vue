<script setup lang="ts">
import { RouterView } from "vue-router";
import { useColorMode } from "@vueuse/core";

useColorMode();
</script>

<template>
	<RouterView />
</template>

<style scoped></style>
