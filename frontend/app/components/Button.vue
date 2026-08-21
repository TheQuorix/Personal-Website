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
    <Text 
      class="select-none text-neutral-500 text-36px" 
      :class="{ 'text-violet-400': active }"
    >
      <slot />
    </Text>
  </div>
</template>