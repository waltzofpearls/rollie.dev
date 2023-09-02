import { writable } from 'svelte/store'

function createStore(): any {
  const { subscribe, set } = writable([])

  return {
    subscribe,

    fetchData: async (): any => {
      let data = []

      const response = await fetch("/api/projects")
      const json = await response.json()
      if (json !== null) {
        json.user?.repositories?.nodes?.forEach((repo) => {
          data.push({
            name: repo.name,
            description: repo.description,
            url: repo.url,
            stars: repo.stargazerCount,
            forks: repo.forkCount,
            languages: repo.languages.nodes,
          })
        })
      }

      set(data)
    },
  }
}

export const projects = createStore()
