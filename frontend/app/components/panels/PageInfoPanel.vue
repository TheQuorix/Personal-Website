<template>
  <NamedPanel :x="4" :y="1" :title="'Page Info'">
    <div flex items-end mt="20px">
      <Text text="48px white" font-bold line-height-tight>?</Text>
      <Text font-bold pb-1 ml="3px">total</Text>

      <Text text="48px white" ml="25px" font-bold line-height-tight>?</Text>
      <Text font-bold pb-1 ml="3px">today</Text>

      <Text text="48px white" ml="25px" font-bold line-height-tight>?</Text>
      <Text font-bold pb-1 ml="3px">unique</Text>
    </div>
  </NamedPanel>
</template>