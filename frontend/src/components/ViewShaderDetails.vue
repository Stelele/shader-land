<template>
    <form @submit.prevent="update">
        <div class="card card-compact bg-base-100 shadow-xl">
            <div class="card-body">
                <div class="flex flex-col gap-4">
                    <label v-if="props.isEditable" class="input input-bordered flex items-center">
                        <input type="text" class="grow" v-model="name" placeholder="Name of your shader" />
                    </label>
                    <label v-if="!props.isEditable" class="">
                        {{ props.name }}
                    </label>
                    <label v-if="props.isEditable" class="flex items-center">
                        <textarea class="grow textarea textarea-bordered" v-model="description"
                            placeholder="Describe your shader"></textarea>
                    </label>
                    <label v-if="!props.isEditable" class="">
                        {{ props.description }}
                    </label>
                    <button v-if="isEditable" class="btn btn-primary" type="submit">Update</button>
                </div>
            </div>
        </div>
    </form>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

export interface Props {
    isEditable?: boolean
    name?: string
    description?: string
}

const props = withDefaults(defineProps<Props>(), {
    isEditable: false,
    name: '',
    description: '',
})

const name = ref(props.name)
const description = ref(props.description)


const emit = defineEmits<{
    onUpdate: [string, string]
}>()

function update() {
    emit('onUpdate', name.value, description.value)
}


</script>