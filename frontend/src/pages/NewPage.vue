<template>
    <div class="grid grid-rows-2" v-if="clerk?.loaded">
        <div>
            <ShaderPlayground ref="shaderPlayground" :start-code="StartShaderFs">
                <SubmitShaderDetails v-if="isSignedIn" @on-submit="onSubmit" />
            </ShaderPlayground>
        </div>
        <div class="grid grid-cols-3">

        </div>
    </div>
    <ShaderSkeletonLoader v-else />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth, useClerk, useUser } from '@clerk/vue'
import ShaderPlayground from '../components/ShaderPlayground.vue'
import SubmitShaderDetails from '../components/SubmitShaderDetails.vue'
import ShaderSkeletonLoader from '../views/ShaderSkeletonLoader.vue'
import { StartShaderFs } from '../components/Renderer/Start.shader'
import { ShaderService } from '../services/ShaderService'
import { ShaderRequest } from '../services/types/ShaderServiceTypes'

const shaderPlayground = ref<InstanceType<typeof ShaderPlayground> | null>(null)
const clerk = useClerk()

const { isSignedIn } = useUser()
const { getToken } = useAuth()
const router = useRouter()

async function onSubmit(name: string, description: string) {
    const token = await getToken.value() ?? ""
    const request: ShaderRequest = {
        name,
        description,
        creationDate: new Date().getTime(),
        code: shaderPlayground.value?.getShaderCode() ?? "",
    }

    const response = await ShaderService.postShader(request, token)

    router.push(`/view/${response.url}`)
}

</script>