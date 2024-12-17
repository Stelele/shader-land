import { StartShaderFs, StartShaderVs } from "./Start.shader"

export class Render {
    private canvas: HTMLCanvasElement
    private device!: GPUDevice
    private presentationFormat!: GPUTextureFormat
    private context!: GPUCanvasContext

    // pipeline
    private pipeline!: GPURenderPipeline
    private renderPassDescriptor!: GPURenderPassDescriptor

    public constructor(canvas: HTMLCanvasElement) {
        this.canvas = canvas
    }

    public async init() {
        await this.initDevice()
        this.setupPipeline()
        this.setupRenderPassDescriptor()
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

    private setupPipeline() {
        const vertexShader = this.device.createShaderModule({
            label: "Vertex Shader",
            code: StartShaderVs,
        })
        const fragmentShader = this.device.createShaderModule({
            label: "Fragment Shader",
            code: StartShaderFs,
        })

        this.pipeline = this.device.createRenderPipeline({
            label: "Render Pipeline",
            layout: "auto",
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

    private startAnimation(targetFps: number) {
        let prev = new Date()
        let render = this
        const targetMs = 1000 / targetFps

        requestAnimationFrame(animate)
        function animate() {
            const cur = new Date()
            const diff = cur.getTime() - prev.getTime()
            if (diff > targetMs) {
                prev = cur
                render.render()
            }

            requestAnimationFrame(animate)
        }

    }

    private render() {
        const view = this.context.getCurrentTexture().createView()
        for (const colorAttachment of this.renderPassDescriptor.colorAttachments) {
            if (colorAttachment) {
                colorAttachment.view = view
            }
        }

        const encoder = this.device.createCommandEncoder()
        const pass = encoder.beginRenderPass(this.renderPassDescriptor)

        pass.setPipeline(this.pipeline)

        pass.draw(3)

        pass.end()
        this.device.queue.submit([encoder.finish()])
    }
}