export interface Shader {
    id: string
    url: string
    userId: string
    userName: string
    name: string
    description: string
    code: string
    creationDate: number
}

export interface ShaderRequest {
    name: string
    description: string
    code: string
    creationDate: number
}