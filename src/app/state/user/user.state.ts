import { Type } from "@angular/core"
import { UserEffects } from "./user.effects"

export interface UserState {
    username: string,
    access: string,
    refresh: string,
    loading: boolean,
    error: any
    success: boolean
}

export const userInitialState: UserState = {
    username: "",
    access: "",
    refresh: "",
    loading: false,
    error: null,
    success: false
}

export const userEffects: Type<unknown>[] = [
    UserEffects
]