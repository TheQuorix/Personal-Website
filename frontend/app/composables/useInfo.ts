import { useFetch } from 'nuxt/app'
import type { Info } from '../../shared/types/info'

export function useInfo() {
  const config = useRuntimeConfig()
  return useFetch<Info>('/api/v1/info', {
    baseURL: config.public.apiBase as string,
    key: 'info'
  })
}