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
      hey: `Hey, I'm Rollie. Thanks for stopping by 🙏`,
      intro: [
        `I'm a Go developer, and I like fried chicken 🍗, burritos 🌯, LEGO bricks 🧱, indoor rowing 🚣, and retro games 🕹`,
        'I use Go, Rust, Python and JavaScript in my side projects, and I have a huge interest in OSS and computer vision.'
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
