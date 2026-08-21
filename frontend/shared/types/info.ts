import type { GitHubData } from './github'
import type { LastFmTrack } from './lastfm'
import type { SteamData } from './steam'
import type { WeatherData } from './weather'

export interface Info {
  weather: WeatherData
  music: LastFmTrack
  steam: SteamData
  github: GitHubData
}