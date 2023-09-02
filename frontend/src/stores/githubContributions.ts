import { writable } from 'svelte/store'

function createStore(): any {
  const { subscribe, set } = writable([])

  return {
    subscribe,

    fetchData: async (): any => {
      let data = []

      const response = await fetch("/api/contributions")
      const json = await response.json()
      if (json !== null) {
        json.user?.contributionsCollection?.contributionCalendar?.weeks?.forEach((item) => {
          item.contributionDays?.forEach((day) => {
            data.push({
              date: day.date,
              value: day.contributionCount,
            })
          })
        })
      }

      set(data)
    },
  }
}

export const githubContributions = createStore()
