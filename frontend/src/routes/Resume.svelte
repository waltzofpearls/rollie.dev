<script lang="ts">
  import { useNavigate, useLocation } from 'svelte-navigator'
  import { onMount } from 'svelte'
  import { marked } from 'marked'
  import LoadingIndicator from '../components/LoadingIndicator.svelte'
  import { resume } from '../stores/resume'

  export let navigate = useNavigate()
	export let location = useLocation()

  let loaded = false

  onMount(async () => {
    await resume.fetchData()
    loaded = true
  })
</script>

{#if !loaded}
  <LoadingIndicator {loaded} />
{:else}
  <section class="resume-section technology">
    <h3>Technology</h3>
    <ul>
      {#each $resume.technology as tech}
        <li>
          <span class="badge text-bg-{tech.label}">{tech.text}</span>
        </li>
      {/each}
    </ul>
  </section>

  <section class="resume-section strengths">
    <h3>Strengths</h3>
    <ol>
      {#each $resume.strengths as strn}
        <li>
          {strn}
        </li>
      {/each}
    </ol>
  </section>

  <section class="resume-section experience">
    <h3>Experience</h3>
    <ul>
      {#each $resume.experience as expr}
        <li>
          <h4>
            {expr.title}
          </h4>
          <h5>
            <span class="badge text-bg-primary">{expr.period[0]} - {expr.period[1]}</span>
            <span class="badge text-bg-info">{expr.location}</span>
          </h5>
          <div class="bg-secondary bg-gradient bg-opacity-25 p-3 my-3">
            <strong>{expr.company.name}</strong>,
            <a href="{expr.company.website}" rel="external">{expr.company.website}</a>
            <br>
            {expr.company.description}
          </div>
          <div class="job-description">
            {@html marked.parse(expr.description)}
          </div>
        </li>
      {/each}
    </ul>
  </section>

  <section class="resume-section education">
    <h3>Education</h3>
    <ul>
      {#each $resume.education as educ}
        <li>
          <h4>
            {educ.school}
          </h4>
          <h5>
            <span class="badge text-bg-primary">{educ.period[0]} - {educ.period[1]}</span>
            <span class="badge text-bg-info">{educ.degree}</span>
          </h5>
          <p>
            {educ.description}
          </p>
          <p>
            Courses included:
            {#each educ.courses as cour}
              <span class="badge text-bg-secondary me-1">{cour}</span>
            {/each}
          </p>
        </li>
      {/each}
    </ul>
  </section>

  <section class="resume-section awards">
    <h3>Awards</h3>
    <ul>
      {#each $resume.awards as awrd}
        <li>
          <span class="badge text-bg-success">{awrd.year}</span>
          {awrd.award}
        </li>
      {/each}
    </ul>
  </section>
{/if}

<style lang="scss">
  .resume-section {
    &.technology {
      ul {
        list-style: none;
        padding-left: 20px;

        li {
          display: inline-block;
          *display: inline;
          font-size: 16px;
          padding: 2px 1px;
        }
      }
    }

    &.strengths {
    }

    &.experience {
      ul {
        list-style: none;
        padding-left: 20px;

        >li {
          &:not(:last-child) {
            margin-bottom: 20px;
          }

          >h4 {
            margin-bottom: 0;
          }

          >h5 {
            margin-top: 5px;
            margin-bottom: 0;
          }

          .well {
            margin-top: 15px;
            margin-bottom: 10px;
          }

          .job-description {
            ol {
              margin-bottom: 10px;
            }
          }
        }
      }
    }

    &.education {
      ul {
        list-style: none;
        padding-left: 20px;

        li {
          >h4 {
            margin-bottom: 0;
          }

          >h5 {
            margin-top: 5px;
            margin-bottom: 15px;
          }
        }
      }
    }

    &.awards {
      ul {
        list-style: none;
        padding-left: 20px;
      }
    }
  }
</style>
