<script setup lang="ts">
const props = defineProps<{
  day: GitHubDay
}>()

const colorClass = computed(() => {
  switch (props.day.level) {
    case 1: return '#6d28d9'
    case 2: return '#7c3aed'
    case 3: return '#8b5cf6'
    case 4: return '#a78bfa'
    default: return '#404040'
  }
})

const dateFormatter = new Intl.DateTimeFormat('en-US', {
  month: 'short',
  day: 'numeric',
  year: 'numeric',
})

const formattedDate = computed(() => dateFormatter.format(new Date(props.day.date)))

const tooltipText = computed(() => {
  const noun = props.day.count === 1 ? 'contribution' : 'contributions'
  return `${props.day.count} ${noun} on ${formattedDate.value}`
})

const isHovered = ref(false)
</script>

<template>
  <div
    relative
    inline-block
    @mouseenter="isHovered = true"
    @mouseleave="isHovered = false"
  >
    <div
      class="peer"
      rounded-5px
      duration-300
      :opacity="isHovered ? '100' : '50'"
      :style="{ backgroundColor: colorClass }"
      w-35px
      h-35px
      m-0
    />
    <div
      opacity-0
      translate-y-5px
      duration-300
      peer-hover:opacity-100
      peer-hover:translate-y-0
      absolute
      bottom-full
      left="1/2"
      -translate-x="1/2"
      mb-8px
      px-12px
      py-6px
      rounded-30px
      bg="neutral-800/85"
      backdrop-blur-3.5px
      border="~ neutral-500"
      text="12px white"
      whitespace-nowrap
      pointer-events-none
      z-10
    >
      {{ tooltipText }}
    </div>
  </div>
</template>