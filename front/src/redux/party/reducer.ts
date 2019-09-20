import { Party } from "../../entity/party";
import { SET_PARTY } from "./constant";

const init: Party = {
  isFinished: false
};

export const partyReducer = (state: Party = init, action: any): Party => {
  const { type, payload } = action;

  switch (type) {
    case SET_PARTY:
      return payload;
    default:
      return state;
  }
};
