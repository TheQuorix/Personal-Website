<script lang="ts" setup>
const activeTab = defineModel<'projects' | 'tools' | 'comments'>({ 
  default: 'comments' 
})
</script>

<template>
  <NamedPanel :x="7" :y="1" :title="'Controller'">
    <div flex justify-center>
      <Button :active="activeTab === 'projects'" @click="activeTab = 'projects'">
        Projects
      </Button>
      <Button :active="activeTab === 'tools'" mx="150px" @click="activeTab = 'tools'">
        Tools
      </Button>
      <Button :active="activeTab === 'comments'" @click="activeTab = 'comments'">
        Comments
      </Button>
    </div>
  </NamedPanel>
</template>