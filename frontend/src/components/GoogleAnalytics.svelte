<script lang="ts">
  import { onMount } from 'svelte'

  onMount(() => {
    const script = document.createElement('script')
    script.src = 'https://www.googletagmanager.com/gtag/js?id=G-2GKBWFT9P3'
    script.async = true
    script.onload = () => {
      window.dataLayer = window.dataLayer || []
      function gtag() { dataLayer.push(arguments); }
      gtag('js', new Date())
      gtag('config', 'G-2GKBWFT9P3')
    };
    document.head.appendChild(script);
  });
</script>
