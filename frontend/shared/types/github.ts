export interface GitHubDay {
    date: string
    count: number
    level: number
}

export interface GitHubWeek {
    days: GitHubDay[]
}

export interface GitHubData {
    followers: number
    repos: number
    contributions: number
    calendar: GitHubWeek[]
}