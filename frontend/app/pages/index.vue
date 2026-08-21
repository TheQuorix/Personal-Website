<script setup lang="ts">
  const { data: info, pending, error} = await useInfo()

  const currentTab = ref<'projects' | 'tools' | 'comments'>('comments')
</script>

<template>
    <BasePanel v-if="pending" :x="1" :y="1" style="--reveal-i: 0" absolute top="1/2" left="1/2" -translate-x="1/2" -translate-y="1/2" flex justify-center items-center>
        <Text>Loading...</Text>
    </BasePanel>

    <BasePanel v-else-if="error" :x="1" :y="1" style="--reveal-i: 0" absolute top="1/2" left="1/2" -translate-x="1/2" -translate-y="1/2" flex justify-center items-center>
        <Text>Error!</Text>
    </BasePanel>

    <div v-else-if="info" overflow-y-auto cursor-default fixed inset="0">
        <div flex flex-col w-1660px gap-20px mx-auto py-30>
            <div flex gap-20px>
                <AvatarPanel style="--reveal-i: 0"/>
                <BioPanel style="--reveal-i: 1"/>
                <WeatherPanel :weather="info.weather" style="--reveal-i: 2"/>
            </div>

            <div flex gap-20px>
                <MusicPanel :track="info.music" style="--reveal-i: 0"/>
                <TechnologiesPanel style="--reveal-i: 1"/> 
            </div>

            <div flex gap-20px>
                <GithubPanel :github="info.github" style="--reveal-i: 0"/>
                <LinksPanel style="--reveal-i: 1"/>
            </div>

            <div flex gap-20px>
                <SystemInfoPanel style="--reveal-i: 0"/>
                <GameActivityPanel :steam="info.steam" style="--reveal-i: 1"/>
            </div>

            <div flex gap-20px>
                <PageInfoPanel style="--reveal-i: 0"/>
                <NamedPanel :x="3" :y="1" :title="'Placeholder'" style="--reveal-i: 1"></NamedPanel>
            </div>

            <div flex gap-20px>
                <ControllerPanel v-model="currentTab" style="--reveal-i: 0"/>
            </div>

            <div v-if="currentTab == 'comments'" flex gap-20px>
                <CommentsPanel/>
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