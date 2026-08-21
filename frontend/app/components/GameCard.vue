<script lang="ts" setup>
const props = defineProps<{
  img: string
  title: string
  twoWeekPlaytime?: number
  totalPlaytime?: number
}>()

const twoWeekHours = computed(() => (props.twoWeekPlaytime ?? 0) / 60)
const totalHours = computed(() => (props.totalPlaytime ?? 0) / 60)
</script>

<template>
  <div flex w="240px" h="100px" items-center>
    <img :src="props.img" :alt="props.title" w-80px h-80px object-contain rounded-10px />
    <div ml-5px>
      <Text w-145px text="16px white" truncate>{{ props.title }}</Text>
      <Text w-145px text="12.5px neutral-500" truncate v-if="twoWeekHours > 0">
        {{ twoWeekHours.toFixed(1) }}h last 2 weeks
      </Text>
      <Text w-145px text="12.5px neutral-500" truncate v-if="totalHours > 0">
        {{ totalHours.toFixed(1) }}h total
      </Text>
    </div>
  </div>
</template>