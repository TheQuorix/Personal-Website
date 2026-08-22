import { useFetch } from 'nuxt/app'

export function useComments() {
  const config = useRuntimeConfig()
  const baseURL = import.meta.server ? config.apiBase : config.public.apiBase
  return useFetch<Comment[]>('/api/v1/comments', {
    baseURL: baseURL as string,
    key: 'comments'
  })
}