import { Party } from "../../entity/party";
import { SET_PARTY, LOADING_PARTY } from "./constant";
import { Action } from "../utils/action";

export const setParty = (payload: Party | null): Action<Party | null> => ({
  type: SET_PARTY,
  payload: payload
});

export const loadingParty = (): Action<{}> => ({
  type: LOADING_PARTY
});
