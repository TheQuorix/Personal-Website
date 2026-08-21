<script setup lang="ts">
  const { data: info, pending, error} = await useInfo()

  const currentTab = ref<'projects' | 'tools' | 'comments'>('comments')
</script>

<template>
    <BasePanel v-if="pending" :x="2" :y="2" style="--reveal-i: 0">
        <Text>Loading...</Text>
    </BasePanel>

    <BasePanel v-else-if="error" :x="2" :y="2" style="--reveal-i: 0">
        <Text>Error! {{ error }}</Text>
    </BasePanel>

    <div v-else-if="info" overflow-y-auto cursor-default fixed inset="0">
        <div flex flex-col w-1660px gap-20px mx-auto py-30>
            <div flex gap-20px>
                <AvatarPanel/>
                <BioPanel/>
                <WeatherPanel :weather="info.weather"/>
            </div>

            <div flex gap-20px>
                <MusicPanel :track="info.music"/>
                <TechnologiesPanel/> 
            </div>

            <div flex gap-20px>
                <GithubPanel :github="info.github"/>
                <LinksPanel/>
            </div>

            <div flex gap-20px>
                <SystemInfoPanel/>
                <GameActivityPanel :steam="info.steam"/>
            </div>

            <div flex gap-20px>
                <PageInfoPanel/>
                <NamedPanel :x="3" :y="1" :title="'Placeholder'"></NamedPanel>
            </div>

            <div flex gap-20px>
                <ControllerPanel v-model="currentTab" />
            </div>
        </div>
    </div>
</template>

<style>
html, body, #__nuxt {
  height: 100%;
  overflow: hidden;
  margin: 0;
  padding: 0;
}

body {
  background-color: #0A0A0A;
  background-image: radial-gradient(circle, #171717 2px, transparent 1px);
  background-size: 30px 30px;
}
</style>