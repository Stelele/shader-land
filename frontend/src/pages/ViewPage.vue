<template>
    <div v-if="clerk?.loaded">
        <ShaderPlayground :start-code="shaderDetails?.code ?? ''">
            <ViewShaderDetails :is-editable="isEditable" :name="shaderDetails?.name"
                :description="shaderDetails?.description" />
        </ShaderPlayground>
    </div>
    <ShaderSkeletonLoader v-else />
</template>

<script lang="ts" setup>
import { computed, onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useClerk, useUser } from '@clerk/vue'
import ShaderPlayground from '../components/ShaderPlayground.vue'
import ViewShaderDetails from '../components/ViewShaderDetails.vue'
import ShaderSkeletonLoader from '../views/ShaderSkeletonLoader.vue'
import { ShaderService } from '../services/ShaderService'
import { Shader } from '../services/types/ShaderServiceTypes'

const route = useRoute()
const { user } = useUser()
const clerk = useClerk()

const id = ref(route.params.id as string)
const shaderDetails = ref<Shader | undefined>(undefined)

const isEditable = computed(() => !!user.value?.id && !!shaderDetails.value?.userId && user.value?.id === shaderDetails.value?.userId)

onMounted(() => {
    getShaderDetails()
})

watch(
    () => route.params.id,
    (newId) => {
        id.value = newId as string
        getShaderDetails()
    }
)

async function getShaderDetails() {
    const details = await ShaderService.getShader(id.value)
    console.log(details)
    shaderDetails.value = details
}

</script>