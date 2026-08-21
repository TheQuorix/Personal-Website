import type { ObjectDirective } from 'vue'

let observer: IntersectionObserver | null = null

function getObserver() {
  if (import.meta.client && !observer) {
    observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            entry.target.classList.add('is-visible')
            observer?.unobserve(entry.target)
          }
        })
      },
      {
        threshold: 0.15,
        rootMargin: '0px 0px -50px 0px',
      }
    )
  }
  return observer
}

export const vReveal: ObjectDirective<HTMLElement> = {
  mounted(el) {
    el.classList.add('v-reveal-init')
    getObserver()?.observe(el)
  },
  unmounted(el) {
    getObserver()?.unobserve(el)
  },
  getSSRProps() {
    return {}
  },
}