export function useVisitTracking() {
  const config = useRuntimeConfig()
  const route = useRoute()

  const trackVisit = (path: string) => {
    if (import.meta.server) return

    const url = `${config.public.apiBase}/api/v1/visits?p=${encodeURIComponent(path)}`
    const sent = navigator.sendBeacon(url)

    if (!sent) {
      fetch(url, { method: 'POST', keepalive: true }).catch(() => {})
    }
  }

  onMounted(() => {
    trackVisit(route.fullPath)
  })

  watch(
    () => route.fullPath,
    (newPath) => trackVisit(newPath)
  )
}