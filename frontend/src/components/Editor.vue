<template>
  <div id="editor-monaco" ref="editorMonaco" class="w-full h-full min-h-3"></div>
</template>

<script lang="ts" setup>
import loader from '@monaco-editor/loader';
import { editor } from "monaco-editor"
import { onMounted, ref } from 'vue';

interface Props {
  theme?: "vs" | "vs-dark"
  startingCode: string
}

const props = withDefaults(defineProps<Props>(), {
  theme: "vs-dark"
})
const emit = defineEmits(["onValueChange"])

onMounted(() => {
  loadEditor()
})

const editorMonaco = ref<HTMLDivElement | null>(null)
let editorInstance: editor.IStandaloneCodeEditor

async function loadEditor() {
  const monacoInstance = await loader.init()
  const modelTemp = monacoInstance.editor.createModel(props.startingCode, "wgsl")
  editorInstance = monacoInstance.editor.create(editorMonaco.value as HTMLDivElement, {
    language: 'wgsl',
    minimap: { enabled: false },
    theme: props.theme,
    automaticLayout: true,
  })
  editorInstance.setModel(modelTemp)

  editorInstance.getModel()?.onDidChangeContent(() => {
    emit("onValueChange", editorInstance.getModel()?.getValue())
  })
}


</script>