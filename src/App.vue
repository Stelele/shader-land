<template>
  <Editor @on-value-change="onEditorChange" />
</template>

<script setup lang="ts">
import Editor from './components/Editor.vue';

function onEditorChange(code: string) {
  console.log(code)
}
</script>
