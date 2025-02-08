import { StartShaderFs, StartShaderVs, Uniforms } from "./Start.shader"

export class Render {
    private static device: GPUDevice
    private canvas: HTMLCanvasElement
    private context!: GPUCanvasContext
    private presentationFormat!: GPUTextureFormat

    // pipeline
    private pipeline!: GPURenderPipeline
    private renderPassDescriptor!: GPURenderPassDescriptor

    // bind group
    private bindGroup!: GPUBindGroup
    private bindGroupLayout!: GPUBindGroupLayout

    // buffers
    private iTimeUniform!: GPUBuffer
    private iFrameUniform!: GPUBuffer
    private iFrameRateUniform!: GPUBuffer
    private iTimeDeltaUniform!: GPUBuffer
    private iResolutionUniform!: GPUBuffer

    // update functions
    private updateFunc!: (info: BufferInfo) => void
    public isPaused = false
    public reset = false

    // recording
    private videoStream?: MediaStream
    private mediaRecorder?: MediaRecorder
    private videoChunks?: Blob[]

    public constructor(canvas: HTMLCanvasElement) {
        this.canvas = canvas
    }

    public async init(updateFunc: (time: BufferInfo) => void, onReady: () => void) {
        await this.initDevice()
        this.loadFragmentShader(StartShaderFs)
        this.setupRenderPassDescriptor()
        this.setupBuffers()
        this.updateFunc = updateFunc
        this.startAnimation()
        onReady()
    }

    private async initDevice() {
        if (!Render.device) {
            const adapter = await navigator.gpu.requestAdapter()
            if (!adapter) {
                throw new Error("Failed to get adapter")
            }

            Render.device = await adapter.requestDevice()
        }

        this.presentationFormat = navigator.gpu.getPreferredCanvasFormat()
        const context = this.canvas.getContext("webgpu")
        if (!context) {
            throw new Error("Failed to get WebGPU context")
        }
        this.context = context
        this.context.configure({
            device: Render.device,
            format: this.presentationFormat,
        })
    }

    public loadFragmentShader(code: string) {
        const fragmentShader = `
        ${Uniforms}
        ${code}`
        this.setupPipeline(fragmentShader)
    }

    private setupPipeline(fragmentShaderCode: string) {
        const vertexShader = Render.device.createShaderModule({
            label: "Vertex Shader",
            code: StartShaderVs,
        })
        const fragmentShader = Render.device.createShaderModule({
            label: "Fragment Shader",
            code: fragmentShaderCode,
        })

        this.bindGroupLayout = Render.device.createBindGroupLayout({
            label: "Bind group layout",
            entries: [
                {
                    binding: 0,
                    visibility: GPUShaderStage.FRAGMENT,
                    buffer: { type: "uniform" }
                },
                {
                    binding: 1,
                    visibility: GPUShaderStage.FRAGMENT,
                    buffer: { type: "uniform" }
                },
                {
                    binding: 2,
                    visibility: GPUShaderStage.FRAGMENT,
                    buffer: { type: "uniform" }
                },
                {
                    binding: 3,
                    visibility: GPUShaderStage.FRAGMENT,
                    buffer: { type: "uniform" }
                },
                {
                    binding: 4,
                    visibility: GPUShaderStage.FRAGMENT,
                    buffer: { type: "uniform" }
                },
            ]
        })
        const pipelineLayout = Render.device.createPipelineLayout({
            label: "Pipeline Layout",
            bindGroupLayouts: [this.bindGroupLayout]
        })

        this.pipeline = Render.device.createRenderPipeline({
            label: "Render Pipeline",
            layout: pipelineLayout,
            vertex: {
                module: vertexShader
            },
            fragment: {
                module: fragmentShader,
                targets: [{ format: this.presentationFormat }]
            }
        })
    }

    private setupRenderPassDescriptor() {
        this.renderPassDescriptor = {
            label: "Render Pass Descriptor",
            // @ts-ignore
            colorAttachments: [
                {
                    storeOp: "store",
                    loadOp: "clear",
                    clearValue: [0, 0, 0, 0]
                }
            ]
        }
    }

    private setupBuffers() {
        this.iResolutionUniform = Render.device.createBuffer({
            label: "iResolution Uniform",
            size: 3 * 4,
            usage: GPUBufferUsage.UNIFORM | GPUBufferUsage.COPY_DST,
        })
        this.iTimeUniform = Render.device.createBuffer({
            label: "iTime Uniform",
            size: 1 * 4,
            usage: GPUBufferUsage.UNIFORM | GPUBufferUsage.COPY_DST,
        })
        this.iTimeDeltaUniform = Render.device.createBuffer({
            label: "iTimeDelta Uniform",
            size: 1 * 4,
            usage: GPUBufferUsage.UNIFORM | GPUBufferUsage.COPY_DST,
        })
        this.iFrameRateUniform = Render.device.createBuffer({
            label: "iFrameRate Uniform",
            size: 1 * 4,
            usage: GPUBufferUsage.UNIFORM | GPUBufferUsage.COPY_DST,
        })
        this.iFrameUniform = Render.device.createBuffer({
            label: "iFrame Uniform",
            size: 1 * 4,
            usage: GPUBufferUsage.UNIFORM | GPUBufferUsage.COPY_DST,
        })

        this.bindGroup = Render.device.createBindGroup({
            label: "Bind Group",
            layout: this.bindGroupLayout,
            entries: [
                { binding: 0, resource: { buffer: this.iResolutionUniform } },
                { binding: 1, resource: { buffer: this.iTimeUniform } },
                { binding: 2, resource: { buffer: this.iTimeDeltaUniform } },
                { binding: 3, resource: { buffer: this.iFrameRateUniform } },
                { binding: 4, resource: { buffer: this.iFrameUniform } },
            ]
        })

    }

    private startAnimation() {
        let prev = new Date()
        let render = this

        const bufferInfo: BufferInfo = {
            resolution: [0, 0, 0],
            time: 0,
            timeDelta: 0,
            frameRate: 0,
            frame: 0
        }

        requestAnimationFrame(animate)
        function animate() {
            if (render.isPaused) {
                prev = new Date()
                requestAnimationFrame(animate)
                return
            }
            if (render.reset) {
                bufferInfo.time = 0
                bufferInfo.timeDelta = 0
                bufferInfo.frameRate = 0
                bufferInfo.frame = 0
                render.reset = false
            }

            const cur = new Date()
            const diff = (cur.getTime() - prev.getTime()) / 1000

            prev = cur

            bufferInfo.time = bufferInfo.time + diff
            bufferInfo.timeDelta = diff
            bufferInfo.frameRate = 1 / diff
            bufferInfo.frame += 1

            render.render(bufferInfo)

            requestAnimationFrame(animate)
        }

    }

    private render(info: BufferInfo) {
        this.context.canvas.width = this.canvas.clientWidth
        this.context.canvas.height = this.canvas.clientHeight
        const texture = this.context.getCurrentTexture()
        info.resolution = [this.context.canvas.width, this.context.canvas.height, 0]
        const view = texture.createView()
        for (const colorAttachment of this.renderPassDescriptor.colorAttachments) {
            if (colorAttachment) {
                colorAttachment.view = view
            }
        }

        const encoder = Render.device.createCommandEncoder()
        const pass = encoder.beginRenderPass(this.renderPassDescriptor)

        pass.setPipeline(this.pipeline)

        Render.device.queue.writeBuffer(this.iResolutionUniform, 0, new Float32Array(info.resolution))
        Render.device.queue.writeBuffer(this.iTimeUniform, 0, new Float32Array([info.time]))
        Render.device.queue.writeBuffer(this.iTimeDeltaUniform, 0, new Float32Array([info.timeDelta]))
        Render.device.queue.writeBuffer(this.iFrameRateUniform, 0, new Float32Array([info.frameRate]))
        Render.device.queue.writeBuffer(this.iFrameUniform, 0, new Uint32Array([info.frame]))

        pass.setBindGroup(0, this.bindGroup)
        pass.draw(3)

        pass.end()
        Render.device.queue.submit([encoder.finish()])

        // call update function
        this.updateFunc(info)
    }

    public recordVideo(isRecording: boolean) {
        if (!isRecording) {
            this.saveVideo()
            return
        }
        this.startRecordingVideo()
    }

    private saveVideo() {
        this.mediaRecorder?.stop()
    }

    private startRecordingVideo() {
        this.videoStream = this.canvas.captureStream()
        this.mediaRecorder = new MediaRecorder(this.videoStream)
        this.videoChunks = []
        const renderer = this

        this.mediaRecorder.start()

        this.mediaRecorder.ondataavailable = (e) => {
            renderer.videoChunks?.push(e.data)
        }

        this.mediaRecorder.onstop = () => {
            const blob = new Blob(renderer.videoChunks, { type: "video/mp4" })
            const videoUrl = URL.createObjectURL(blob)

            const link = document.createElement('a')
            link.download = "recordingVideo"
            link.href = videoUrl
            link.click()
        }
    }
}

export interface BufferInfo {
    resolution: number[],
    time: number,
    timeDelta: number,
    frameRate: number,
    frame: number
}