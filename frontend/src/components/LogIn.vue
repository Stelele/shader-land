<template>
    <div class="relative flex flex-col justify-center h-screen overflow-hidden">
        <div class="w-full p-6 m-auto rounded-md shadow-md lg:max-w-lg">
            <h1 class="text-3xl font-semibold text-center text-gray-700">DaisyUI</h1>
            <div class="space-y-4">
                <div>
                    <label class="label">
                        <span class="text-base label-text">Email</span>
                    </label>
                    <input type="text" placeholder="Email Address" class="w-full input input-bordered" />
                </div>
                <div>
                    <label class="label">
                        <span class="text-base label-text">Password</span>
                    </label>
                    <input type="password" placeholder="Enter Password" class="w-full input input-bordered" />
                </div>
                <a href="#" class="text-xl text-gray-600 hover:underline hover:text-blue-600">Forget Password?</a>
                <div>
                    <button class="btn-neutral btn btn-block">Login</button>
                </div>
                <p class="text-xl text-gray-600">Don't have an account? <span
                        class="text-blue-600 hover:cursor-pointer hover:underline">Sign
                        Up</span></p>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
const emit = defineEmits(["onForgotPassword", "onLogin"])
</script>