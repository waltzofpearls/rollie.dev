import { writable, type Readable } from 'svelte/store'

export type Metadata = Readable<{
  title: string
  description: string[]
  url: string
  image: string
}>

function createStore(): any {
  const { subscribe, set } = writable<Metadata>({
    title: 'Rollie Ma - Gopher, Rustacean, Pythonist, Amateur Archaeologist and Corgi Wrangler from Vancouver, BC',
    description: {
      hey: `Hey, I'm Rollie:`,
      intro: [
        'LEGO bricks and Linux enthusiast.',
        'Self-motivated and fascinated by robotics, computer vision and machine learning.',
        'Polyglot developer captivated by Go, Python and JavaScript.',
      ],
    },
    url: 'https://rollie.dev',
    image: 'https://rollie.dev/images/logos/logo-120x120.png',
  })

  return {
    subscribe,

    description: (): string => {
      let description = ''
      subscribe((metadata) => {
        description = `${metadata.description.hey} ${metadata.description.intro.join(' ')}`
      })
      return description
    },
  }
}

export const metadata = createStore()
