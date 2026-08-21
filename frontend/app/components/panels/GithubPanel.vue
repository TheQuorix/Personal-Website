<script lang="ts" setup>
const props = defineProps<{
  github: GitHubData
}>()

const recentWeeks = computed(() => props.github.calendar.slice(-21))
</script>

<template>
  <NamedPanel :x="4" :y="2" :title="'GitHub'">
    <div flex gap-x-5px w-min>
      <div flex v-for="week in recentWeeks" flex-col gap="y-5px">
        <GithubDay v-for="day in week.days" :day="day"/>
      </div>
    </div>
    <div flex items-end mt="15px">
      <Text text="48px white" font-bold line-height-tight>{{ github.contributions }}</Text>
      <Text font-bold pb-1 ml="3px">contributions</Text>

      <Text text="48px white" ml="25px" font-bold line-height-tight>{{ github.followers }}</Text>
      <Text font-bold pb-1 ml="3px">followers</Text>

      <Text text="48px white" ml="25px" font-bold line-height-tight>{{ github.repos }}</Text>
      <Text font-bold pb-1 ml="3px">public repos</Text>
    </div>
  </NamedPanel>
</template>