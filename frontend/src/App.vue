<template>
  <div id="app" class="container">
    <img src="@/assets/logo.png" alt="Logo" class="logo">
    <webhook-table :webhooks="webhooks" @toggleActive="toggleActive" />
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import WebhookTable from './components/WebhookTable.vue'
import axios from 'axios'

export default {
  name: 'App',
  components: {
    WebhookTable
  },
  setup() {
    const webhooks = ref({})

    const fetchWebhooks = async () => {
      const response = await axios.get('/webhooks')
      webhooks.value = response.data
    }

    onMounted(fetchWebhooks)

    const toggleActive = async (repo_name) => {
      await axios.get(`/webhooks/${repo_name}/toggle_active`)
      await axios.get(`/webhooks/refresh`)
      await fetchWebhooks()
    }

    return {
      webhooks,
      toggleActive
    }
  }
}
</script>

<style>
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 15px;
  box-sizing: border-box;
  font-family: Arial, sans-serif;
}

.logo {
  display: block;
  width: 200px; 
  height: auto; 
  margin: 0 auto 20px;
}
</style>
