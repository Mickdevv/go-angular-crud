import { ItemsEffects } from "./items/items.effects"
import { itemsReducer } from "./items/items.reducers"
import { ItemsState } from "./items/items.state"

export interface AppState {
    toDoList: ItemsState
}

export const effects = [
    ItemsEffects
]

export const reducers = {
    items: itemsReducer
}