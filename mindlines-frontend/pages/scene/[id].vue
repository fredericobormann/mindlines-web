<script setup lang="ts">
import { type Answer, answers, type Character, type Line, type SceneContent } from "~/types/scenes.type";

type LearningState = 'asking' | 'checking';

  const route = useRoute();
  const queryClient = useQueryClient();
  const backendUrl = useRuntimeConfig().public.backendUrl;
  const { data: scene, isLoading, error } = useQuery<SceneContent>({
    queryFn: async () => (await fetch(`${backendUrl}/scenes/${route.params.id}`)).json(),
    queryKey: ['scene', route.params.id],
  });

  const { mutate: saveCardRating, isPending: isSaveCardRatingPending } = useMutation({
    mutationKey: ['scene', route.params.id, 'learn'],
    mutationFn: async (vars : {lineIndex: number, rating: Answer}) => (await fetch(`${backendUrl}/scenes/${route.params.id}?lineIndex=${vars.lineIndex}&rating=${answers[vars.rating]}`, {
      method: 'POST',
    })).json(),
    onSuccess(data, variables, context) {
        queryClient.setQueryData(['scene', route.params.id], data);
        increaseLineIfNecessary();
        learningState.value = 'asking';
    },
  })

  const MY_CHARACTER: Character = 'BIEDERMANN';
  const currentLine = ref<number | undefined>(findNextCharacterLine(scene.value?.content, 0, MY_CHARACTER));
  const learningState = ref<LearningState>('asking')

  watch(scene, s => {
    if(currentLine.value == undefined && s) {
      currentLine.value = findNextCharacterLine(s.content, 0, MY_CHARACTER);
    }
  })

  const visibleLines = computed(() => scene.value?.content.slice(0, currentLine.value ?? 0))

  watch(visibleLines, async _ => {
    await nextTick();
    window.scrollTo(0, document.body.scrollHeight);
  })

  function increaseLineIfNecessary() {
    currentLine.value = findNextCharacterLine(scene.value?.content, (currentLine.value ?? - 1), MY_CHARACTER);
  }

  function reveal() {
    if(currentLine.value == undefined) return;
    currentLine.value += 1;
    learningState.value = 'checking';
  }

  function jumpToEarliestDueLine() {
    if (earliestDueLine.value == undefined) {
      return;
    }
    currentLine.value = earliestDueLine.value;
  }

  const earliestDueLine = computed(() => {
    return scene.value?.content.findIndex(l => l.character == MY_CHARACTER && new Date(l.dueTime).getTime() < new Date().getTime());
  });

  const couldJumpBack = computed(() => {
    return earliestDueLine.value != undefined && currentLine.value != undefined && earliestDueLine.value + 1 < currentLine.value;
  });

  const currentReviewTimeValues = computed(() => {
    if (scene.value == undefined || currentLine.value == undefined) {
      return;
    }
    return scene.value.content[currentLine.value].reviewTimes;
  })

  function findNextCharacterLine(lines?: Line[], startFrom: number = 0, character: Character = 'BIEDERMANN') {
    const lineIx = lines?.slice(startFrom).findIndex(l => l.character === character && new Date(l.dueTime).getTime() < new Date().getTime());
    if (lineIx == undefined) {
      return;
    }
    return lineIx + startFrom;
  }

  function saveAnswer(answer: Answer) {
    if (currentLine.value == undefined) return;
    if (!scene.value?.content?.[currentLine.value - 1]) return;
    saveCardRating({lineIndex: currentLine.value - 1, rating: answer })
  }
</script>

<template>
  <div>
    <TechnicalError v-if="error">{{ error.message }}</TechnicalError>
    <div class="flex flex-row justify-center">
      <UButton
          v-if="couldJumpBack"
          color="red"
          class="fixed top-8 rounded-2xl"
          @click="jumpToEarliestDueLine"
      >
        Jump to due line
      </UButton>
    </div>
    <h1 class="text-4xl font-extrabold font-serif">{{ scene?.name }}</h1>
    <div v-if="scene" class="flex-col flex gap-y-4 items-start py-20">
      <p v-for="line of visibleLines" class="font-serif">
        <mark v-if="line.character === 'BIEDERMANN'" class="mark-teal mark-rounded">{{ line.character }}: {{ line.line }}</mark>
        <template v-else>
          {{ line.character }}: {{ line.line }}
        </template>
      </p>
    </div>
    <div class="justify-center flex flex-row">
      <UButton v-if="learningState === 'asking'" class="fixed bottom-8" @click="reveal" size="xl">Continue</UButton>
      <UButtonGroup v-else-if="learningState === 'checking' && !isSaveCardRatingPending" class="fixed bottom-8" size="xl">
        <UButton @click="saveAnswer('again')" color="red">Again<br />{{ currentReviewTimeValues?.again }}</UButton>
        <UButton @click="saveAnswer('hard')" color="gray">Hard<br />{{ currentReviewTimeValues?.hard }}</UButton>
        <UButton @click="saveAnswer('good')">Good<br />{{ currentReviewTimeValues?.good }}</UButton>
        <UButton @click="saveAnswer('easy')" color="blue">Easy<br />{{ currentReviewTimeValues?.easy }}</UButton>
      </UButtonGroup>
      <NuxtLoadingIndicator v-if="isSaveCardRatingPending" />
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
