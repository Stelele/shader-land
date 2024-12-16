<template>
  <div id="editor-monaco" ref="editorMonaco"></div>
</template>

<script lang="ts" setup>
import loader from '@monaco-editor/loader';
import { editor } from "monaco-editor"
import { ref } from 'vue';

interface Props {
  theme?: "vs" | "vs-dark"
}

const props = withDefaults(defineProps<Props>(), {
  theme: "vs-dark"
})
const emit = defineEmits(["onValueChange"])

const editorMonaco = ref<HTMLDivElement | null>(null)
let editorInstance: editor.IStandaloneCodeEditor

loader.init().then(monacoInstance => {
  const sampleCode = `@fragment
fn fs() -> @location(0) vec4f {
  return vec4f(0.5, 0.3, 0.2, 1);
}`
  const modelTemp = monacoInstance.editor.createModel(sampleCode, "wgsl")
  editorInstance = monacoInstance.editor.create(editorMonaco.value as HTMLDivElement, {
    language: 'wgsl',
    minimap: { enabled: false },
    theme: props.theme,

  })
  editorInstance.setModel(modelTemp)

  editorInstance.getModel()?.onDidChangeContent(() => {
    emit("onValueChange", editorInstance.getModel()?.getValue())
  })
})


</script>

<style scoped>
#editor-monaco {
  height: 100%;
  width: 100%;
  min-height: 50vh;
  font-family: sans-serif;
}
</style>