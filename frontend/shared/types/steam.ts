export interface SteamGame {
  app_id: number
  name: string
  icon_url: string
  playtime_2weeks: number
  playtime_forever: number
}

export interface SteamData {
  recent: SteamGame[]
  top: SteamGame[]
}