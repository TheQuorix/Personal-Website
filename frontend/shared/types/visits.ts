export interface DailyStat {
  date: string
  total: number
  unique: number
}

export interface VisitStats {
  total_visits: number
  unique_visits: number
  daily: DailyStat[]
}