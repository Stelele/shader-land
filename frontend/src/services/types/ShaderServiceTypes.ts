export interface Shader {
    id: string
    email: string
    userName: string
    name: string
    tags: string
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