import { Party } from "../../entity/party";
import { SET_PARTY } from "./constant";

const init = null;

export const partyReducer = (
  state: Party | null = init,
  action: any
): Party | null => {
  const { type, payload } = action;

  switch (type) {
    case SET_PARTY:
      return payload;
    default:
      return state;
  }
};
