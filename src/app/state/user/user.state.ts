import { Type } from "@angular/core"
import { UserEffects } from "./user.effects"

export interface UserState {
    username: string,
    access: string,
    refresh: string,
    loading: boolean,
    error: any
}

export const userInitialState: UserState = {
    username: "",
    access: "",
    refresh: "",
    loading: false,
    error: null
}

export const userEffects: Type<unknown>[] = [
    UserEffects
]