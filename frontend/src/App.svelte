<script lang="ts">
  import { Router, Route, Link, createHistory } from 'svelte-navigator'
  import createHashSource from './lib/history'
  import { metadata } from './stores/metadata'
  import favicon from './assets/images/favicon.ico'
  import Home from './routes/Home.svelte'
  import Projects from './routes/Projects.svelte'
  import Resume from './routes/Resume.svelte'
  import GoogleAnalytics from './components/GoogleAnalytics.svelte'
  import 'bootstrap/dist/js/bootstrap.bundle.min.js'

  const history = createHistory(createHashSource())

  let cssLoaded = false
  let smallLogo: HTMLElement
  let smallLogoContainer: HTMLElement
  let bigLogoContainer: HTMLElement
  let navbar: HTMLElement
  let navbarHeader: HTMLElement

  function watchScroll() {
    const smallLogoHeight = smallLogo.clientHeight
    const bigLogoHeight = bigLogoContainer.clientHeight
    const navbarHeight = navbarHeader.clientHeight

    const smallLogoEndPos = 0;
    const smallSpeed = (smallLogoHeight / bigLogoHeight);
    const ySmall = (window.scrollY * smallSpeed);

    let smallPadding = navbarHeight - ySmall;
    if (smallPadding > navbarHeight) {
      smallPadding = navbarHeight;
    }
    if (smallPadding < smallLogoEndPos) {
      smallPadding = smallLogoEndPos;
    }
    if (smallPadding < 0) {
      smallPadding = 0;
    }

    smallLogoContainer.style.width = (smallLogoHeight > smallPadding) ? '50px' : 0
    smallLogoContainer.style.paddingTop = `${smallPadding}px`

    let navOpacity = ySmall / smallLogoHeight;
    if (navOpacity > 1) {
      navOpacity = 1;
    }
    if (navOpacity < 0) {
      navOpacity = 0;
    }

    let shadowOpacity = navOpacity * 0.3;
    if (ySmall > 1) {
      navbar.style.boxShadow = `0 2px 3px rgba(0,0,0,${shadowOpacity})`
    } else {
      navbar.style.boxShadow = 'none'
    }
  }
</script>

<svelte:head>
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
  <meta name="description" content="{metadata.description()}">
  <meta property="og:title" content="{$metadata.title}">
  <meta property="og:description" content="{metadata.description()}">
  <meta property="og:image" content="{$metadata.image}">
  <meta property="og:type" content="article">
  <meta property="og:url" content="{$metadata.url}">
  <meta name="twitter:card" content="{$metadata.url}">
  <meta name="twitter:description" content="{metadata.description()}">
  <meta name="twitter:image" content="{$metadata.image}">
  <meta name="twitter:title" content="{$metadata.title}">
  <link rel="icon" type="image/x-icon" href="{favicon}" />
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/svelte-material-ui@7.0.0-beta.15/bare.min.css" on:load={() => cssLoaded = true} />
</svelte:head>

<svelte:window on:scroll={watchScroll} />

{#if !cssLoaded}
  <div class="preloading">Loading...</div>
{:else}
  <Router {history}>
    <header>
      <nav class="navbar navbar-expand-lg navbar-light bg-light fixed-top tetris-nav" bind:this={navbar}>
        <div class="container" bind:this={navbarHeader}>
          <div class="navbar-brand">RM</div>
          <div class="small-logo-container" bind:this={smallLogoContainer}>
            <a class="small-logo icon-48 icon-logo-48" href="/" bind:this={smallLogo}>&nbsp;</a>
          </div>
          <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target=".tetris-navbar-collapse" aria-controls="tetris-navbar-collapse" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
          </button>

          <div class="collapse navbar-collapse tetris-navbar-collapse">
            <ul class="navbar-nav">
              <li class="nav-item"><Link class="nav-link" to="/">0: "Home"</Link></li>
              <li class="nav-item"><Link class="nav-link" to="projects">1: "Projects"</Link></li>
              <li class="nav-item"><Link class="nav-link" to="resume">2: "Resume"</Link></li>
              <li class="nav-item"><a class="nav-link" href="https://blog.rollie.dev/">3: "Blog"</a></li>
            </ul>
            <div class="tetris-social-network ms-auto">
              <a class="icon-36 icon-github-36" href="https://github.com/waltzofpearls" target="_blank">&nbsp;</a>
              <a class="icon-36 icon-facebook-36" href="https://facebook.com/rolli3ma" target="_blank">&nbsp;</a>
              <a class="icon-36 icon-linkedin-36" href="https://linkedin.com/in/rolli3ma" target="_blank">&nbsp;</a>
              <a class="icon-36 icon-twitter-36" href="https://twitter.com/rolli3ma" target="_blank">&nbsp;</a>
            </div>
          </div>
        </div>
      </nav>

      <div class="container">
        <div class="row big-logo-row" bind:this={bigLogoContainer}>
          <!-- Logo container -->
          <div class="col-2 ps-0">
            <a href="/">
              <span class="logo logo-150x150"></span>
              <span class="logo logo-120x120"></span>
              <span class="logo logo-80x80"></span>
              <span class="logo logo-50x50"></span>
            </a>
          </div>
          <!-- Text container -->
          <div class="col-10 pe-0">
            <h1 class="long">let rollie_ma = "dev"</h1>
            <h1 class="medium">let rollie = "dev"</h1>
            <h1 class="short">let rm = "dev"</h1>
          </div>
        </div>
      </div>
    </header>

    <div class="container">
      <Route path="projects" component="{Projects}" primary={false} />
      <Route path="resume" component="{Resume}" primary={false} />
      <Route component="{Home}" primary={false}></Route>
    </div>

    <footer>
      <div class="container">
        <div class="row">
          <div class="col-5">
            &copy; {new Date().getFullYear()} Rollie Ma
            <br>
            Built with <a href="https://svelte.dev/" target="_blank">Svelte</a> + <a href="https://www.rust-lang.org/" target="_blank">Rust</a>
            <br>
            Source code on <a href="https://github.com/waltzofpearls/rollie.dev" target="_blank">GitHub</a>
          </div>
          <div class="col-7 text-end">
            rollie (at) rollie (dot) dev
            <br>
            rollie (at) topbass (dot) studio
            <br>
            rollie (at) topbasslabs (dot) com
          </div>
        </div>
      </div>
    </footer>
  </Router>

  <GoogleAnalytics />
{/if}

<style lang="scss">
  .preloading {
    width: 100vw;
    height: 100vh;
    line-height: 100vh;
    text-align: center;
    font-size: 1.5em;
  }
</style>
