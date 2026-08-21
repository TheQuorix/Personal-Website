<script lang="ts" setup>
const technologyVectors = [
  '/vector/technologies/java.svg',
  '/vector/technologies/kotlin.svg',
  '/vector/technologies/csharp.svg',
  '/vector/technologies/nuxt.svg',
  '/vector/technologies/typescript.svg',
  '/vector/technologies/python.svg',
  '/vector/technologies/rust.svg',
  '/vector/technologies/flutter.svg',
  '/vector/technologies/arduino.svg',
]

const scrollContainer = ref<HTMLElement | null>(null)
let isDragging = false
let startX = 0
let startScrollLeft = 0

function handleMouseDown(e: MouseEvent) {
  if (!scrollContainer.value) return
  isDragging = true
  startX = e.pageX
  startScrollLeft = scrollContainer.value.scrollLeft
}

function handleMouseMove(e: MouseEvent) {
  if (!isDragging || !scrollContainer.value) return
  e.preventDefault()
  const delta = e.pageX - startX
  scrollContainer.value.scrollLeft = startScrollLeft - delta
}

function stopDragging() {
  isDragging = false
}
</script>

<template>
  <NamedPanel :x="4" :y="1" :title="'Technologies'">
    <div
      ref="scrollContainer"
      class="no-scrollbar"
      cursor-grab
      select-none
      overflow-x-scroll
      flex
      gap-10px
      w-880px
      rounded-20px
      h-106px
      @mousedown="handleMouseDown"
      @mousemove="handleMouseMove"
      @mouseup="stopDragging"
      @mouseleave="stopDragging"
    >
      <Badge v-for="technology in technologyVectors" :path="technology"/>
    </div>
  </NamedPanel>
</template>

<style>
.no-scrollbar {
  scrollbar-width: none;
  -ms-overflow-style: none;
}
.no-scrollbar::-webkit-scrollbar {
  display: none;
}
</style>