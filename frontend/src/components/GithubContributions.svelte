<script lang="ts">
  import { onMount } from 'svelte'
  import LoadingIndicator from '../components/LoadingIndicator.svelte'
  import { githubContributions } from '../stores/githubContributions'
  import CalHeatmap from 'cal-heatmap'
  import Tooltip from 'cal-heatmap/plugins/Tooltip'
  import LegendLite from 'cal-heatmap/plugins/LegendLite'
  import CalendarLabel from 'cal-heatmap/plugins/CalendarLabel'
  import 'cal-heatmap/cal-heatmap.css'

  let loaded = false
  const cal = new CalHeatmap()

  onMount(async () => {
    await githubContributions.fetchData()
    loaded = true

    const options = {
      itemSelector: '.heatmap',
      range: 13,
      domain: {
        type: 'month',
        gutter: 4,
        label: {text: 'MMM', textAlign: 'start', position: 'top'},
      },
      subDomain: {type: 'ghDay', radius: 2, width: 14, height: 14, gutter: 4},
      data: {source: $githubContributions, x: 'date', y: 'value'},
      date: {start: new Date($githubContributions[0].date)},
      scale: {
        color: {
          type: 'linear',
          range: ['#ebedf0', '#9be9a8', '#40c463', '#30a14e', '#216e39'],
          domain: [0, 2, 4, 6, 8],
        },
      },
    }
    const plugins = [
      [Tooltip, {
          text: (date, value, dayjsDate) => {
            return `${value ? value : 'No'} contributions on ${dayjsDate.format('dddd, MMMM D, YYYY')}`
          },
      }],
      [LegendLite, {
        itemSelector: '.legend', radius: 2, width: 14, height: 14, gutter: 4,
      }],
      [CalendarLabel, {
          width: 30,
          textAlign: 'start',
          text: () => dayjs.weekdaysShort().map((d, i) => (i % 2 == 0 ? '' : d)),
          padding: [25, 0, 0, 0],
      }],
    ]

    cal.paint(options, plugins)
  })
</script>

<div class="gh-contrib">
  <h4>GitHub Contributions</h4>
  {#if !loaded}
    <LoadingIndicator {loaded} />
  {:else}
    <div class="outer">
      <div class="inner">
        <div class="heatmap-row">
          <div class="heatmap"></div>
        </div>
        <div class="legend-row">
          <div class="legend-container">
            <span style="color: #768390;">Less</span>
            <div class="legend" style="display: inline-block; margin: 0 4px;"></div>
            <span style="color: #768390; font-size: 12px;">More</span>
          </div>
        </div>
      </div>
    </div>
  {/if}
</div>

<style lang="scss">
  .gh-contrib {
    margin-top: 25px;
    text-align: center;

    .outer {
      border: 1px solid #d0d7de;
      border-top-left-radius: 6px;
      border-top-right-radius: 6px;
      display: flex;
      justify-content: center;

      .inner {
        width: 100%;
        margin: 10px;
        padding: 10px;
        display: flex;
        flex-direction: column;

        .heatmap-row {
          display: flex;
          justify-content: center;
          padding: 10px;
          overflow-x: auto;
          overflow-y: hidden;
        }

        .legend-row {
          display: flex;
          justify-content: center;

          .legend-container {
            text-align: right;
            font-size: 12px;
            margin: 10px 20px;
            max-width: 1000px;
            flex-grow: 1;
          }
        }
      }
    }
  }
</style>
