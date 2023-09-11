import { writable } from 'svelte/store'

function createStore(): any {
  const { subscribe, set } = writable({})

  return {
    subscribe,

    fetchData: async (): any => {
      const response = await fetch("/api/resume")
      const json = await response.json()
      if (json !== null) {
        set(json)
      }
    },
  }
}

export const resume = createStore()
