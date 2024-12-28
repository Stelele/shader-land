import { StartShaderFs, StartShaderVs, Uniforms } from "./Start.shader"

export class Render {
    private device!: GPUDevice
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

    public constructor(canvas: HTMLCanvasElement) {
        this.canvas = canvas
    }

    public async init() {
        await this.initDevice()
        this.loadFragmentShader(StartShaderFs)
        this.setupRenderPassDescriptor()
        this.setupBuffers()
        this.startAnimation(120)
    }

    private async initDevice() {
        const adapter = await navigator.gpu.requestAdapter()
        if (!adapter) {
            throw new Error("Failed to get adapter")
        }

        this.device = await adapter.requestDevice()

        this.presentationFormat = navigator.gpu.getPreferredCanvasFormat()
        const context = this.canvas.getContext("webgpu")
        if (!context) {
            throw new Error("Failed to get WebGPU context")
        }
        this.context = context
        this.context.configure({
            device: this.device,
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
        const vertexShader = this.device.createShaderModule({
            label: "Vertex Shader",
            code: StartShaderVs,
        })
        const fragmentShader = this.device.createShaderModule({
            label: "Fragment Shader",
            code: fragmentShaderCode,
        })

        this.bindGroupLayout = this.device.createBindGroupLayout({
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
        const pipelineLayout = this.device.createPipelineLayout({
            label: "Pipeline Layout",
            bindGroupLayouts: [this.bindGroupLayout]
        })

        this.pipeline = this.device.createRenderPipeline({
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
        this.iResolutionUniform = this.device.createBuffer({
            label: "iResolution Uniform",
            size: 3 * 4,
            usage: GPUBufferUsage.UNIFORM | GPUBufferUsage.COPY_DST,
        })
        this.iTimeUniform = this.device.createBuffer({
            label: "iTime Uniform",
            size: 1 * 4,
            usage: GPUBufferUsage.UNIFORM | GPUBufferUsage.COPY_DST,
        })
        this.iTimeDeltaUniform = this.device.createBuffer({
            label: "iTimeDelta Uniform",
            size: 1 * 4,
            usage: GPUBufferUsage.UNIFORM | GPUBufferUsage.COPY_DST,
        })
        this.iFrameRateUniform = this.device.createBuffer({
            label: "iFrameRate Uniform",
            size: 1 * 4,
            usage: GPUBufferUsage.UNIFORM | GPUBufferUsage.COPY_DST,
        })
        this.iFrameUniform = this.device.createBuffer({
            label: "iFrame Uniform",
            size: 1 * 4,
            usage: GPUBufferUsage.UNIFORM | GPUBufferUsage.COPY_DST,
        })

        this.bindGroup = this.device.createBindGroup({
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

    private startAnimation(targetFps: number) {
        let prev = new Date()
        let render = this
        const targetS = 1 / targetFps

        const bufferInfo: BufferInfo = {
            resolution: [0, 0, 0],
            time: 0,
            timeDelta: 0,
            frameRate: 0,
            frame: 0
        }

        requestAnimationFrame(animate)
        function animate() {
            const cur = new Date()
            const diff = (cur.getTime() - prev.getTime()) / 1000

            if (diff > targetS) {
                prev = cur

                bufferInfo.time = bufferInfo.time + diff
                bufferInfo.timeDelta = diff
                bufferInfo.frameRate = 1 / diff
                bufferInfo.frame += 1

                render.render(bufferInfo)
            }

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

        const encoder = this.device.createCommandEncoder()
        const pass = encoder.beginRenderPass(this.renderPassDescriptor)

        pass.setPipeline(this.pipeline)

        this.device.queue.writeBuffer(this.iResolutionUniform, 0, new Float32Array(info.resolution))
        this.device.queue.writeBuffer(this.iTimeUniform, 0, new Float32Array([info.time]))
        this.device.queue.writeBuffer(this.iTimeDeltaUniform, 0, new Float32Array([info.timeDelta]))
        this.device.queue.writeBuffer(this.iFrameRateUniform, 0, new Float32Array([info.frameRate]))
        this.device.queue.writeBuffer(this.iFrameUniform, 0, new Uint32Array([info.frame]))

        pass.setBindGroup(0, this.bindGroup)
        pass.draw(3)

        pass.end()
        this.device.queue.submit([encoder.finish()])
    }
}

interface BufferInfo {
    resolution: number[],
    time: number,
    timeDelta: number,
    frameRate: number,
    frame: number
}