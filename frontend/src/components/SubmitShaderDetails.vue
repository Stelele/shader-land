<template>
    <form @submit.prevent="submit">
        <div class="card card-compact bg-base-100 shadow-xl">
            <div class="card-body">
                <div class="flex flex-col gap-4">
                    <label class="input input-bordered flex items-center">
                        <input type="text" class="grow" v-model="name" placeholder="Name of your shader" />
                    </label>
                    <label class="flex items-center">
                        <textarea class="grow textarea textarea-bordered" v-model="description"
                            placeholder="Describe your shader"></textarea>
                    </label>
                    <button class="btn btn-primary" type="submit">Submit</button>
                </div>
            </div>
        </div>
    </form>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const name = ref('')
const description = ref('')

const emit = defineEmits(["onSubmit"])

function submit() {
    emit('onSubmit', name.value, description.value)
}


</script>