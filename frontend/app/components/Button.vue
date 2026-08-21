<script lang="ts" setup>
const props = defineProps<{
  active: boolean
}>()

const emit = defineEmits<{
  click: []
}>()
</script>

<template>
  <div 
    class="flex cursor-pointer duration-300 hover:scale-105 bg-neutral-500/25 w-250px h-80px justify-center items-center rounded-full"
    :class="{ 'bg-violet-400/25': active }"
    @click="emit('click')"
  >
    <p 
      class="select-none" font="[Fira_Code]" text="32px neutral-500" duration-300
      :class="{ 'text-violet-400 font-bold': active }"
    >
      <slot />
    </p>
  </div>
</template>