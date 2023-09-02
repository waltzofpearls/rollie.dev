<script lang="ts">
  import CircularProgress from '@smui/circular-progress'

  export let loaded = false
</script>

<div class="loading-indicator">
  <CircularProgress class="loader" indeterminate closed={loaded} />
</div>

<style lang="scss">
  @use '@material/circular-progress/mixins' as circular-progress;
  @use '@material/theme/color-palette';

  :global {
    .loading-indicator {
      display: flex;
      justify-content: center;
      margin: 50px auto;

      .loader {
        @include circular-progress.color(color-palette.$blue-600);
        @include circular-progress.track-color(color-palette.$blue-900);
        width: 50px;
        height: 50px;
      }
    }
  }
</style>
