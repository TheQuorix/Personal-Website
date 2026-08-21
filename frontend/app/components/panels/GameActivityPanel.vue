<script lang="ts" setup>
const props = defineProps<{
  steam: SteamData
}>()
</script>

<template>
  <NamedPanel :x="3.5" :y="3" :title="'Game Activity'" flex flex-col>
    <h1 line-height-snug font="[Sora] 500" text="[32px] white" relative inline-block>
      Recently played
    </h1>

    <div flex gap="10px">
      <GameCard v-for="game in steam.recent" :img="`${useRuntimeConfig().public.apiBase}${game.icon_url}`" :title=game.name :twoWeekPlaytime=game.playtime_2weeks :totalPlaytime=game.playtime_forever />
    </div>

    <h1 line-height-snug font="[Sora] 500" text="[32px] white" relative inline-block>
      Top games by playtime
    </h1>

    <div flex flex-wrap gap="x-10px">
      <GameCard v-for="game in steam.top" :img="`${useRuntimeConfig().public.apiBase}${game.icon_url}`" :title=game.name :twoWeekPlaytime=game.playtime_2weeks :totalPlaytime=game.playtime_forever />
    </div>
  </NamedPanel>
</template>