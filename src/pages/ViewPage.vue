<template>
    <div class=" grid grid-cols-5 gap-4">
        <div class="card bg-base-100 col-span-2 w-full shadow-xl">
            <div class="card-body pl-1 pr-1">
                <Renderer />
            </div>
        </div>
        <div class="card bg-base-100 col-span-3 w-full shadow-xl">
            <div class="card-body rounded-2xl min-h-[70vh]">
                <Editor />
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import Editor from '../components/Editor.vue';
import Renderer from '../components/Renderer/Renderer.vue';

const route = useRoute()
const id = ref(route.params.id ?? "")

</script>