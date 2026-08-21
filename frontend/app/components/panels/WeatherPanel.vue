<script lang="ts" setup>
const props = defineProps<{
  weather: WeatherData
}>()

const timeFormatter = new Intl.DateTimeFormat('en-US', {
  timeZone: 'Europe/Moscow',
  hour: '2-digit',
  minute: '2-digit',
  hour12: false,
})

const secondsFormatter = new Intl.DateTimeFormat('en-US', {
  timeZone: 'Europe/Moscow',
  second: '2-digit',
})

const weekdayFormatter = new Intl.DateTimeFormat('en-US', {
  timeZone: 'Europe/Moscow',
  weekday: 'long',
})

const dateFormatter = new Intl.DateTimeFormat('en-GB', {
  timeZone: 'Europe/Moscow',
  day: 'numeric',
  month: 'long',
})

const time = ref('')
const seconds = ref('')
const weekday = ref('')
const date = ref('')

function updateValues() {
  const now = new Date()

  time.value = timeFormatter.format(now)
  seconds.value = secondsFormatter.format(now).padStart(2, '0')
  weekday.value = weekdayFormatter.format(now)
  date.value = dateFormatter.format(now)
}

let interval: ReturnType<typeof setInterval>

onMounted(() => {
  updateValues()
  interval = setInterval(updateValues, 1000)
})

onUnmounted(() => {
  clearInterval(interval)
})
</script>

<template>
  <NamedPanel :x="2" :y="1" :title="'Moscow'" :subTitle="`${Math.round(weather.temp)}°`">
    <div flex items-end>
      <Text text="48px white" font-bold line-height-tight>{{ time }}</Text>
      <Text font-bold pb-1>:{{ seconds }}</Text>
    </div>
    <Text>{{ weekday }}, {{ date }}</Text>
  </NamedPanel>
</template>