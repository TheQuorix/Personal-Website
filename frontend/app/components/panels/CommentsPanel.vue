<script setup lang="ts">
const { data: comments, pending, error, refresh } = await useComments()

const config = useRuntimeConfig()

const author = ref('')
const message = ref('')
const publish = ref(true)
const isSubmitting = ref(false)

async function postComment() {
  if (!message.value.trim() || isSubmitting.value) return

  const authorName = author.value.trim() || 'Anonymous'
  isSubmitting.value = true

  try {
    await $fetch(`${config.public.apiBase}/api/v1/comments`, {
      method: 'POST',
      body: {
        author: authorName,
        message: message.value.trim(),
        publish: publish.value
      }
    })

    author.value = ''
    message.value = ''
    publish.value = true

    await refresh()

  } catch (err) {
    console.error('Ошибка при отправке комментария:', err)
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <NamedPanel :x="7" :y="3" :title="'Comments'">
    <div flex justify-between>
      <div>
        <p font="[Fira_Code]" text="32px neutral-500" ml-20px mb-3px>Nickname</p>
        <input v-model="author" placeholder="Anonymous" type="text" w-150 font="[Fira_Code]" text="32px white" p-20px rounded-20px bg="neutral-800/25" backdrop-blur="[3.5px]" border="~ neutral-500">
        <p font="[Fira_Code]" text="32px neutral-500" ml-20px mt-7px mb-3px>Message</p>
        <textarea v-model="message" resize-none placeholder="Say something" h-65 w-150 font="[Fira_Code]" text="32px white" p-20px rounded-20px bg="neutral-800/25" backdrop-blur="[3.5px]" border="~ neutral-500"></textarea>
        
        <div flex w-150 justify-between mt-5>
            <Button 
            :active="message.trim() !== '' && !isSubmitting" 
            @click="postComment"
          >
            {{ isSubmitting ? 'Sending...' : 'Send' }}
          </Button>

            <label flex items-center duration-300 hover:scale-99>
                <input v-model="publish" appearance-none border="2px neutral-500" bg="neutral-800/25" backdrop-blur="[3.5px]" rounded="10px" h-10 w-10 bg-cover class="checked:bg-[url(/vector/check.svg)]" type="checkbox">
                <p font="[Fira_Code]" text="32px neutral-500" ml-20px>Publish</p>
            </label>
        </div>

        <p font="[Fira_Code]" text="20px neutral-500" mt-10px text-center>All comments are moderated first</p>
    </div>
      <div flex flex-col gap-20px rounded-20px h-140 overflow-scroll>
        <div v-if="pending" text="neutral-400 24px" p-20px>
          Loading comments...
        </div>

        <div v-else-if="error" text="red-400 24px" p-20px>
          Failed to load comments
        </div>

        <div v-else-if="!comments || comments.length === 0" text="neutral-500 24px" p-20px>
          No comments yet. Be the first!
        </div>

        <Comment 
          v-for="(comment, index) in comments?.toReversed()" 
          :comment="comment"
        />
      </div>
    </div>
  </NamedPanel>
</template>