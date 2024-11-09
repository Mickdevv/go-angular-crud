import { Type } from "@angular/core"
import { UserEffects } from "./user.effects"

export interface UserState {
    username: string,
    token: string,
    loading: boolean,
    error: any
}

export const userInitialState: UserState = {
    username: "",
    token: "",
    loading: false,
    error: null
}

export const userEffects: Type<unknown>[] = [
    UserEffects
]