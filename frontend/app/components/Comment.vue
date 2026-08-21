<script setup lang="ts">
const props = defineProps<{
  comment: Comment
}>()

const dateFormatter = new Intl.DateTimeFormat('ru-RU', {
  hour: '2-digit',
  minute: '2-digit',
  day: '2-digit',
  month: '2-digit',
  year: 'numeric',
})

const formattedDate = computed(() => dateFormatter.format(new Date(props.comment.date)))
</script> 

<template>
    <div w-225 p-20px rounded-20px bg="neutral-800/25" backdrop-blur="[3.5px]" border="~ neutral-500">
        <div flex justify-between items-center>
            <h1 font="[Fira_Code]" text="32px violet-400" ml-20px>{{ comment.author }}</h1>
            <h1 font="[Fira_Code]" text="26px neutral-500" >{{ formattedDate }}</h1>
        </div>
        
        <h1 font="[Fira_Code]" text="26px white" ml-30px>{{ comment.message }}</h1>
        <h1 v-if="comment.response != ''" font="[Fira_Code]" text="26px neutral-500" ml-20px>> {{ comment.response }}</h1>
    </div>
</template>