<script lang="ts">
  import { metadata } from '../stores/metadata'
  import GithubContributions from '../components/GithubContributions.svelte'

  function getRandomInt(min, max) {
    return Math.floor(Math.random() * (max+1 - min)) + min;
  }
</script>

<div class="tetris-about row">
  <div class="avatar-img col-xs-12 col-sm-4 col-md-3 col-lg-2">
    <div class="avatar avatar-{getRandomInt(1, 5)}">&nbsp;</div>
  </div>
  <div class="about-txt col-xs-12 col-sm-8 col-md-9 col-lg-10">
    <div class="pre-ish">
      <p>{$metadata.description.hey}</p>
      <ul>
        {#each $metadata.description.intro as line}
          <li>{line}</li>
        {/each}
      </ul>
    </div>
  </div>
</div>

<GithubContributions />

<style lang="scss">
  @use '../src/assets/stylesheets/variables' as *;

  .tetris-about {
    margin-top: 15px;

    .avatar-img {
      .avatar {
        // smartphone
        @media (max-width: $screen-xs-max) {
          margin: 0 auto 25px auto;
        }
      }
    }

    .about-txt {
      .pre-ish {
        font-family: Consolas, monospace, 'Helvetica Neue', Helvetica, Arial, sans-serif;
        font-size: 1.2em;
        border: 1px solid #c0c0c0;
        background: #f5f5f5;
        padding: 10px;
        -moz-box-shadow: 0 0 8px rgba(0, 0, 0, 0.5);
        -webkit-box-shadow: 0 0 8px rgba(0, 0, 0, 0.5);
        box-shadow: 0 0 8px rgba(0, 0, 0, 0.5);
      }

      ul {
        padding-left: 30px;
      }
    }
  }
</style>
