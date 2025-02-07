<template>
  <div v-if="clerk?.loaded">
    <TopAppBar />
    <RouterView />
  </div>
  <SkeletonLoader v-else />
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import TopAppBar from './components/TopAppBar.vue'
import { useClerk } from '@clerk/vue'
import SkeletonLoader from './views/SkeletonLoader.vue';

const clerk = useClerk()
</script>
