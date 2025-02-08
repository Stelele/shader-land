<template>
  <div>
    <TopAppBar />
    <RouterView />
  </div>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import TopAppBar from './components/TopAppBar.vue'
import { useClerk } from '@clerk/vue'
</script>
