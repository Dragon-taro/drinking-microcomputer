import { Party } from "../../entity/party";
import { SET_PARTY } from "./constant";
import { Action } from "../utils/action";

export const setParty = (payload: Party): Action<Party> => ({
  type: SET_PARTY,
  payload: payload
});
