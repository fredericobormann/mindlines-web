<template>
  <div class="flex flex-col gap-y-4">
    <h1 class="text-4xl font-extrabold">Select a scene</h1>
    <TechnicalError v-if="isError">{{ error }}</TechnicalError>
    <UButtonGroup orientation="vertical" size="xl">
      <UButton color="white" v-for="scene of scenes" variant="ghost" :key="scene.identifier" :to="`/scene/${scene.index}`">{{ scene.name }}</UButton>
    </UButtonGroup>
  </div>
</template>

<script setup lang="ts">
  import type { Scene } from "~/types/scenes.type";
  import TechnicalError from "~/components/TechnicalError.vue";

  const backendUrl = useRuntimeConfig().public.backendUrl;
  const { data: scenes, isError, error } = useQuery<Scene[]>({
    queryKey: ['scenes'],
    queryFn: async () => (await fetch(`${backendUrl}/scenes`)).json(),
  });
</script>
