<script lang="ts" setup>
const props = defineProps<{
  path: string
}>()
</script>

<template>
  <div flex select-none flex-none mt-3 w-640px pl-10px justify-left items-center h-105px rounded-20px bg="neutral-800/25" backdrop-blur-3.5px border="~ neutral-500" duration-300 hover:scale-95>
    <img select-none draggable="false" object-contain :src=props.path w-85px h-85px rounded-10px>
    <slot/>
  </div>
</template>