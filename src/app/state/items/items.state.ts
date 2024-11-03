import { Type } from '@angular/core'
import { Item } from '../../models/todo.model'
import { ItemsEffects } from './items.effects'

export interface ItemsState {
    items: Item[],
    success: boolean,
    error: any,
    loading: boolean
}

export const itemsInitialState: ItemsState = {
    items: [],
    success: false,
    error: undefined,
    loading: false
}

export const itemEffects: Type<unknown>[] = [
    ItemsEffects
]