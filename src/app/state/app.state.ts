import { ItemsEffects } from "./items/items.effects"
import { itemsReducer } from "./items/items.reducers"
import { ItemsState } from "./items/items.state"
import { UserEffects } from "./user/user.effects"
import { userReducer } from "./user/user.reducers"
import { UserState } from "./user/user.state"

export interface AppState {
    items: ItemsState,
    user: UserState
}

export const effects = [
    ItemsEffects,
    UserEffects
]

export const reducers = {
    items: itemsReducer,
    user: userReducer
}