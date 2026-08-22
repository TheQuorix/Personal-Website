import type { VisitStats } from '../../shared/types/visits'

export function useVisitStats() {
  const config = useRuntimeConfig()

  const stats = ref<VisitStats | null>(null)
  const pending = ref(false)
  const error = ref<string | null>(null)

  const fetchStats = async () => {
    pending.value = true
    error.value = null
    try {
      stats.value = await $fetch<VisitStats>(`${config.public.apiBase}/api/v1/visits/stats`)
    } catch (e) {
      error.value = 'Не удалось загрузить статистику'
    } finally {
      pending.value = false
    }
  }

  const todayStat = computed(() => {
    if (!stats.value) return null
    const today = new Date().toISOString().slice(0, 10) // YYYY-MM-DD
    return stats.value.daily.find(d => d.date.slice(0, 10) === today) ?? { date: today, total: 0, unique: 0 }
  })

  return { stats, pending, error, fetchStats, todayStat }
}