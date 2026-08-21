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
        class="group"
        rounded-30px
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