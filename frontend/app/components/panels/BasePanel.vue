<script lang="ts" setup>
const props = defineProps<{
    x: number
    y: number
}>()

const width = computed(() => 240 * props.x - 20)
const height = computed(() => 240 * props.y - 20)
</script>

<template>
    <div
        v-reveal
        class="group"
        rounded-30px
        duration-300
        hover:scale-99
        bg="neutral-800/25"
        backdrop-blur-3.5px
        border="~ neutral-500"
        p-30px
        :style="{
            width: `${width}px`,
            height: `${height}px`
        }"
    >
        <slot/>
    </div>
</template>

<style>
.v-reveal-init {
  opacity: 0;
  transform: translateY(24px);
  transition: scale 300ms ease, opacity 0.6s ease, transform 0.6s ease;
  transition-delay: calc(var(--reveal-i, 0) * 80ms);
}

.v-reveal-init.is-visible {
  opacity: 1;
  transform: translateY(0);
}

@media (prefers-reduced-motion: reduce) {
  .v-reveal-init {
    opacity: 1 !important;
    transform: none !important;
    transition: none !important;
  }
}
</style>