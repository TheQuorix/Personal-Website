import { useFetch } from 'nuxt/app'

export function useComments() {
  const config = useRuntimeConfig()
  
  return useFetch<Comment[]>('/api/v1/comments', {
    baseURL: config.public.apiBase as string,
    key: 'comments'
  })
}