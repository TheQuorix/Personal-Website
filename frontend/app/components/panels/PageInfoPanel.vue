<script setup lang="ts">
const { stats, pending, error, fetchStats, todayStat } = useVisitStats()

await fetchStats()
</script>

<template>
  <NamedPanel :x="4" :y="1" :title="'Page Info'">
    <div flex items-end mt-20px>
      <Text v-if="stats" text="48px white" font-bold line-height-tight>{{ stats.total_visits }}</Text>
      <Text v-else text="48px white" font-bold line-height-tight>?</Text>
      <Text font-bold pb-1 ml-3px>total</Text>

      <Text v-if="stats" text="48px white" ml-25px font-bold line-height-tight>{{ todayStat?.total ?? 0 }}</Text>
      <Text v-else text="48px white" ml-25px font-bold line-height-tight>?</Text>
      <Text font-bold pb-1 ml-3px>today</Text>

      <Text v-if="stats" text="48px white" ml-25px font-bold line-height-tight>{{ stats.unique_visits }}</Text>
      <Text v-else text="48px white" ml-25px font-bold line-height-tight>?</Text>
      <Text font-bold pb-1 ml-3px>unique</Text>
    </div>
  </NamedPanel>
</template>