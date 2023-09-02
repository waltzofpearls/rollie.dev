<script lang="ts">
  import { useNavigate, useLocation } from 'svelte-navigator'
  import { onMount } from 'svelte'
  import LoadingIndicator from '../components/LoadingIndicator.svelte'
  import { projects } from '../stores/projects'

  export let navigate = useNavigate()
	export let location = useLocation()

  let loaded = false

  onMount(async () => {
    await projects.fetchData()
    loaded = true
  })
</script>

<div class="projects-list">
  <h4>
    Public Repos from
    <a href="https://github.com/waltzofpearls" target="_blank">GitHub</a>
    <span class="badge bg-primary">{$projects.length}</span>
  </h4>

  {#if !loaded}
    <LoadingIndicator {loaded} />
  {:else}
    <div class="card-container row">
      {#each $projects as project}
        <div class="col-4 p-3">
          <div class="card">
            <div class="content">
              <h4>
                <a href="{project.url}" target="_blank">{project.name}</a>
              </h4>
              <p>
                <span class="badge text-bg-secondary">Stars: {project.stars}</span>
                <span class="badge text-bg-secondary">Forks: {project.forks}</span>
                <br>
                {#each project.languages as lang}
                  <span class="badge" style="background-color: {lang.color}; margin-right: 5px;">{lang.name}</span>
                {/each}
              </p>
              <p title="{project.description || ''}">
                {project.description || ''}
              </p>
            </div>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<style lang="scss">
  .projects-list {
    > h4 {
      margin-left: 15px;
    }

    .card-container {
      display: flex;
      flex-wrap: wrap;
      flex-direction: row;
      justify-content: start;

      .card {
        height: 100%;
        border: 1px solid #ccc;
        -moz-box-shadow: 0 0 8px rgba(0, 0, 0, 0.5);
        -webkit-box-shadow: 0 0 8px rgba(0, 0, 0, 0.5);
        box-shadow: 0 0 8px rgba(0, 0, 0, 0.5);

        .content {
          padding: 10px 15px;
        }
      }
    }
  }
</style>
