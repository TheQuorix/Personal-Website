<script setup lang="ts">
const props = defineProps<{
  author: string,
  message: string,
  response: string
}>()
</script> 

<template>
    <div w-225 p-20px rounded-20px bg="neutral-800/25" backdrop-blur="[3.5px]" border="~ neutral-500">
        <h1 font="[Fira_Code]" text="32px violet-400" ml-20px>{{ author }}</h1>
        <h1 font="[Fira_Code]" text="26px white" ml-30px>{{ message }}</h1>
        <h1 v-if="response != ''" font="[Fira_Code]" text="26px neutral-500" ml-20px>> {{ response }}</h1>
    </div>
</template>