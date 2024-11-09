import { inject, Injectable } from "@angular/core";
import { Actions } from "@ngrx/effects";
import { ApiService } from "../../services/api.service";

@Injectable()
export class UserEffects {
    private readonly actions = inject(Actions);
    private readonly itemsService = inject(ApiService);


}