import { useFetch } from 'nuxt/app'
import type { Info } from '../../shared/types/info'

export function useInfo() {
  const config = useRuntimeConfig()
  const baseURL = import.meta.server ? config.apiBase : config.public.apiBase
  return useFetch<Info>('/api/v1/info', {
    baseURL: baseURL as string,
    key: 'info'
  })
}