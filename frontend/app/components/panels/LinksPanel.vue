<template>
  <NamedPanel :x="3" :y="2" :title="'Links'">
    <div flex justify-between w-640px>
      <a href="https://t.me/thequorix">
        <Badge :path="'/vector/links/Telegram.webp'"/>
      </a>
      <a href="https://discord.com/users/408268839206256652">
        <Badge :path="'/vector/links/Discord.webp'"/>
      </a>
      <a href="https://steamcommunity.com/profiles/76561199108831532/">
        <Badge :path="'/vector/links/Steam.webp'"/>
      </a>
      <a href="https://github.com/TheQuorix">
        <Badge :path="'/vector/links/Github.webp'"/>
      </a>
      <a href="https://hackatime.hackclub.com/@quorix">
        <Badge :path="'/vector/links/Hackatime.webp'"/>
      </a>
    </div>
    <a href="https://discord.gg/CKkdMZTmcA">
      <NamedBadge :path="'/vector/links/Discord.webp'">
        <Text text="32px white" pl-15px w-full>My discord server</Text>
      </NamedBadge>
    </a>
    <a href="https://t.me/quorix_dev">
      <NamedBadge :path="'/vector/links/Telegram.webp'">
        <Text text="32px white" pl-15px w-full>My dev channel</Text>
      </NamedBadge>
    </a>
  </NamedPanel>
</template>