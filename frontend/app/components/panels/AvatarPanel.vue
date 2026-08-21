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
    h-220px
    w-220px
    flex
    items-center
    justify-center
  >
    <img w-200px h-200px rounded-20px border="~ neutral-700" src="/avatar.webp" alt="">
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