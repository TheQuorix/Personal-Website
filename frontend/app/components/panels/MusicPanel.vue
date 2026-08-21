<script setup lang="ts">
import type { LastFmTrack } from '#shared/types/lastfm'

const props = defineProps<{
  track: LastFmTrack
}>()

const now = ref(new Date())

let interval: ReturnType<typeof setInterval>

onMounted(() => {
  interval = setInterval(() => {
    now.value = new Date()
  }, 60_000)
})

onUnmounted(() => {
  clearInterval(interval)
})

const trackState = computed(() => {
  if (props.track.now_playing) return 'Now Playing'

  const playedAt = new Date(props.track.date)
  const diffSec = Math.floor((now.value.getTime() - playedAt.getTime()) / 1000)

  if (diffSec < 60) return `${diffSec}s ago`

  const diffMin = Math.floor(diffSec / 60)
  if (diffMin < 60) return `${diffMin}m ago`

  const diffHour = Math.floor(diffMin / 60)
  if (diffHour < 24) return `${diffHour}h ago`

  const diffDay = Math.floor(diffHour / 24)
  return `${diffDay}d ago`
})
</script>

<template>
  <NamedPanel :x="3" :y="1" title="Music" :miniSubTitle="trackState">
    <a target="_blank" :href="track.song_url">
      <div flex>
        <img :src="track.image_url" alt="" w-110px h-110px rounded-20px border="~ neutral-700" mr-10px>
        <div flex flex-col justify-center w-500px>
          <Text truncate text="24px white" font-600>{{ track.name }}</Text>
          <Text truncate>{{ track.artist }}</Text>
        </div>
      </div>
    </a>
  </NamedPanel>
</template>