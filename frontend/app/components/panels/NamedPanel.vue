<script lang="ts" setup>
const props = defineProps<{
  x: number
  y: number
  title: string
  subTitle?: string
  miniSubTitle?: string
}>()
</script>

<template>
    <BasePanel :x=x :y=y>
        <div flex justify-between items-center>
            <h1
                line-height-snug font="[Sora] 900" text="[36px] violet-400" relative inline-block
                class="after:absolute after:left-0 after:w-0 after:duration-300 group-hover:after:w-full after:bottom-1 after:h-0.5 after:-skew-y-3 after:bg-gradient-to-r after:from-violet-400/75 after:via-violet-400/50 after:to-violet-400/25 after:content-['']"
            >
            {{ title }}
            </h1>
            <h1 v-if="props.subTitle" font="Fira_Code 900" text="36px white">{{ props.subTitle }}</h1>
            <h1 v-if="props.miniSubTitle" font="Fira_Code" text="24px neutral-500">{{ props.miniSubTitle }}</h1>
        </div>
        <slot/>
    </BasePanel>
</template>