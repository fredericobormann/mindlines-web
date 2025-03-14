<script setup lang="ts">
import { type Answer, answerLine, type Character, type Line, type SceneContent } from "~/types/scenes.type";

  type LearningState = 'asking' | 'checking';

  const route = useRoute();
  const { data: scene, status, error } = await useFetch<SceneContent>(`/api/scenes/${route.params.id}`);
  const MY_CHARACTER: Character = 'BIEDERMANN';
  const currentLine = ref<number | undefined>(findNextCharacterLine(scene.value?.content, 0, MY_CHARACTER));
  const learningState = ref<LearningState>('asking')

  watch(scene, s => {
    if(currentLine.value == undefined && s) {
      currentLine.value = findNextCharacterLine(s.content, 0, MY_CHARACTER);
    }
  })

  function increaseLine() {
    currentLine.value = findNextCharacterLine(scene.value?.content, (currentLine.value ?? -1) + 1, MY_CHARACTER);
  }

  function reveal() {
    console.log(currentLine.value);
    if(currentLine.value == undefined) return;
    currentLine.value += 1;
    learningState.value = 'checking';
  }


  function findNextCharacterLine(lines?: Line[], startFrom: number = 0, character: Character = 'BIEDERMANN') {
    const lineIx = lines?.slice(startFrom).findIndex(l => l.character === character);
    if (lineIx == undefined) {
      return;
    }
    return lineIx + startFrom;
  }

  function saveAnswer(answer: Answer) {
    if (currentLine.value == undefined) return;
    if (!scene.value?.content?.[currentLine.value - 1]) return;
    scene.value.content[currentLine.value - 1] = answerLine(scene.value?.content?.[currentLine.value - 1], answer);
    console.log(scene.value);
    learningState.value = 'asking';
    increaseLine();
  }
</script>

<template>
  <div>
    <TechnicalError v-if="error">{{ error.message }}</TechnicalError>
    <h1 class="text-4xl font-extrabold font-serif">{{ scene?.name }}</h1>
    <div v-if="scene" class="flex-col flex gap-y-4 items-start py-12">
      <p v-for="line of scene.content.slice(0, currentLine ?? 0)" class="font-serif">
        <mark v-if="line.character === 'BIEDERMANN'" class="mark-teal mark-rounded">{{ line.character }}: {{ line.line }}</mark>
        <template v-else>
          {{ line.character }}: {{ line.line }}
        </template>
      </p>
    </div>
    <div class="justify-center flex flex-row">
      <UButton v-if="learningState === 'asking'" class="fixed bottom-8" @click="reveal" size="xl">Continue</UButton>
      <UButtonGroup v-else-if="learningState === 'checking'" class="fixed bottom-8" size="xl">
        <UButton @click="saveAnswer('wrong')" color="red">Wrong</UButton>
        <UButton @click="saveAnswer('correct')">Correct</UButton>
      </UButtonGroup>
    </div>
  </div>
</template>

<style lang="postcss">
mark {
  @apply pr-1 pl-1 py-0;
  margin-top: 0.15rem;
  margin-bottom: 0.15rem;
}

.mark-rounded {
  @apply inline-block rounded-md;
}

.mark-teal {
  @apply bg-yellow-300 bg-opacity-40;
}
</style>
