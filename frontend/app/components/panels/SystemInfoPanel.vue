<script lang="ts" setup>
const activeTab = ref<'main_os' | 'second_os'>('main_os')

const asciiLogo = `///////////////  ///////////////
///////////////  ///////////////
///////////////  ///////////////
///////////////  ///////////////
///////////////  ///////////////
///////////////  ///////////////
                                
///////////////  ///////////////
///////////////  ///////////////
///////////////  ///////////////
///////////////  ///////////////
///////////////  ///////////////
///////////////  ///////////////`

const systemInfo = `quorix@DESKTOP-UK6S5BS
---------------------
OS: Windows 11 Pro (25H2) x86_64
Host: MS-7D48 (2.0)
Kernel: WIN32_NT 10.0.26200.8875
Uptime: 9 hours, 42 mins
Shell: CMD 10.0.26100.8737
Display (DP): 1920x1080, 165 Hz
Display (V22FAB-RA): 1920x1080, 75 Hz
WM: Desktop Window Manager 10.0.26100.8737
Theme: Fluent
Font: Segoe UI (12pt)
Cursor: Posy's Default Improved
Terminal: Windows Console 10.0.26100.8737
Terminal Font: Consolas (16pt)
CPU: 12th Gen Intel(R) Core(TM) i3-12100F
GPU: NVIDIA GeForce RTX 2060 SUPER
Memory: 13.77 GiB / 15.86 GiB (87%)
Disk (C:\\): 222.59 GiB / 487.51 GiB (46%)
Disk (D:\\): 385.20 GiB / 953.87 GiB (40%)`
</script>

<template>
  <NamedPanel :x="3.5" :y="3" :title="'System Info'">
    <div flex gap-30px rounded-10px mt-5px justify-center items-center w-760px h-480px bg="neutral-950/50">
      <pre font="Fira_Code" text="14px violet-400">{{ asciiLogo }}</pre>
      <pre font="Fira_Code" text="14px white">{{ systemInfo }}</pre>
    </div>

    <div flex justify-evenly mt-25px>
      <Button :active="activeTab === 'main_os'" @click="activeTab = 'main_os'">
        Main OS
      </Button>
      <Button :active="activeTab === 'second_os'" @click="activeTab = 'second_os'">
        Second OS
      </Button>
    </div>
  </NamedPanel>
</template>