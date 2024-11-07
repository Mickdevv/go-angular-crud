import { ItemsEffects } from "./items/items.effects"
import { itemsReducer } from "./items/items.reducers"
import { ItemsState } from "./items/items.state"

export interface AppState {
    items: ItemsState
}

export const effects = [
    ItemsEffects
]

export const reducers = {
    items: itemsReducer
}